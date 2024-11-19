// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubernetes

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	storageinformers "k8s.io/client-go/informers/storage/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	ref "k8s.io/client-go/tools/reference"
	"k8s.io/client-go/util/workqueue"

	"source.monogon.dev/go/logging"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/osbase/fsquota"
	"source.monogon.dev/osbase/supervisor"
)

// inodeCapacityRatio describes the ratio between the byte capacity of a volume
// and its inode capacity. One inode on XFS is 512 bytes and by default 25%
// (1/4) of capacity can be used for metadata.
const inodeCapacityRatio = 4 * 512

// ONCHANGE(//metropolis/node/kubernetes/reconciler:resources_csi.go): needs to
// match csiProvisionerServerName declared.
const csiProvisionerServerName = "dev.monogon.metropolis.vfs"

// csiProvisionerServer is responsible for the provisioning and deprovisioning
// of CSI-based container volumes. It runs on all nodes and watches PVCs for
// ones assigned to the node it's running on and fulfills the provisioning
// request by creating a directory, applying a quota and creating the
// corresponding PV. When the PV is released and its retention policy is
// Delete, the directory and the PV resource are deleted.
type csiProvisionerServer struct {
	NodeName         string
	Kubernetes       kubernetes.Interface
	InformerFactory  informers.SharedInformerFactory
	VolumesDirectory *localstorage.DataVolumesDirectory

	claimQueue           workqueue.RateLimitingInterface
	pvQueue              workqueue.RateLimitingInterface
	recorder             record.EventRecorder
	pvcInformer          coreinformers.PersistentVolumeClaimInformer
	pvInformer           coreinformers.PersistentVolumeInformer
	storageClassInformer storageinformers.StorageClassInformer
	logger               logging.Leveled
}

// runCSIProvisioner runs the main provisioning machinery. It consists of a
// bunch of informers which keep track of the events happening on the
// Kubernetes control plane and informs us when something happens. If anything
// happens to PVCs or PVs, we enqueue the identifier of that resource in a work
// queue. Queues are being worked on by only one worker to limit load and avoid
// complicated locking infrastructure. Failed items are requeued.
func (p *csiProvisionerServer) Run(ctx context.Context) error {
	// The recorder is used to log Kubernetes events for successful or failed
	// volume provisions. These events then show up in `kubectl describe pvc`
	// and can be used by admins to debug issues with this provisioner.
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: p.Kubernetes.CoreV1().Events("")})
	p.recorder = eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: csiProvisionerServerName, Host: p.NodeName})

	p.pvInformer = p.InformerFactory.Core().V1().PersistentVolumes()
	p.pvcInformer = p.InformerFactory.Core().V1().PersistentVolumeClaims()
	p.storageClassInformer = p.InformerFactory.Storage().V1().StorageClasses()

	p.claimQueue = workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	p.pvQueue = workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	p.logger = supervisor.Logger(ctx)

	p.pvcInformer.Informer().SetWatchErrorHandler(func(_ *cache.Reflector, err error) {
		p.logger.Errorf("pvcInformer watch error: %v", err)
	})
	p.pvcInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: p.enqueueClaim,
		UpdateFunc: func(old, new interface{}) {
			p.enqueueClaim(new)
		},
	})
	p.pvInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: p.enqueuePV,
		UpdateFunc: func(old, new interface{}) {
			p.enqueuePV(new)
		},
	})
	p.pvInformer.Informer().SetWatchErrorHandler(func(_ *cache.Reflector, err error) {
		p.logger.Errorf("pvInformer watch error: %v", err)
	})

	p.storageClassInformer.Informer().SetWatchErrorHandler(func(_ *cache.Reflector, err error) {
		p.logger.Errorf("storageClassInformer watch error: %v", err)
	})

	go p.pvcInformer.Informer().Run(ctx.Done())
	go p.pvInformer.Informer().Run(ctx.Done())
	go p.storageClassInformer.Informer().Run(ctx.Done())

	// These will self-terminate once the queues are shut down
	go p.processQueueItems(p.claimQueue, func(key string) error {
		return p.processPVC(key)
	})
	go p.processQueueItems(p.pvQueue, func(key string) error {
		return p.processPV(key)
	})

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	p.claimQueue.ShutDown()
	p.pvQueue.ShutDown()
	return nil
}

// isOurPVC checks if the given PVC is is to be provisioned by this provisioner
// and has been scheduled onto this node
func (p *csiProvisionerServer) isOurPVC(pvc *v1.PersistentVolumeClaim) bool {
	if pvc.ObjectMeta.Annotations["volume.beta.kubernetes.io/storage-provisioner"] != csiProvisionerServerName {
		return false
	}
	if pvc.ObjectMeta.Annotations["volume.kubernetes.io/selected-node"] != p.NodeName {
		return false
	}
	return true
}

// isOurPV checks if the given PV has been provisioned by this provisioner and
// has been scheduled onto this node
func (p *csiProvisionerServer) isOurPV(pv *v1.PersistentVolume) bool {
	if pv.ObjectMeta.Annotations["pv.kubernetes.io/provisioned-by"] != csiProvisionerServerName {
		return false
	}
	if pv.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions[0].Values[0] != p.NodeName {
		return false
	}
	return true
}

// enqueueClaim adds an added/changed PVC to the work queue
func (p *csiProvisionerServer) enqueueClaim(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		p.logger.Errorf("Not queuing PVC because key could not be derived: %v", err)
		return
	}
	p.claimQueue.Add(key)
}

// enqueuePV adds an added/changed PV to the work queue
func (p *csiProvisionerServer) enqueuePV(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		p.logger.Errorf("Not queuing PV because key could not be derived: %v", err)
		return
	}
	p.pvQueue.Add(key)
}

// processQueueItems gets items from the given work queue and calls the process
// function for each of them. It self- terminates once the queue is shut down.
func (p *csiProvisionerServer) processQueueItems(queue workqueue.RateLimitingInterface, process func(key string) error) {
	for {
		obj, shutdown := queue.Get()
		if shutdown {
			return
		}

		func(obj interface{}) {
			defer queue.Done(obj)
			key, ok := obj.(string)
			if !ok {
				queue.Forget(obj)
				p.logger.Errorf("Expected string in workqueue, got %+v", obj)
				return
			}

			if err := process(key); err != nil {
				p.logger.Warningf("Failed processing item %q, requeueing (numrequeues: %d): %v", key, queue.NumRequeues(obj), err)
				queue.AddRateLimited(obj)
			}

			queue.Forget(obj)
		}(obj)
	}
}

// volumePath gets the path where the volume is stored.
func (p *csiProvisionerServer) volumePath(volumeID string) string {
	return filepath.Join(p.VolumesDirectory.FullPath(), volumeID)
}

// processPVC looks at a single PVC item from the queue, determines if it needs
// to be provisioned and logs the provisioning result to the recorder
func (p *csiProvisionerServer) processPVC(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return fmt.Errorf("invalid resource key: %s", key)
	}
	pvc, err := p.pvcInformer.Lister().PersistentVolumeClaims(namespace).Get(name)
	if apierrs.IsNotFound(err) {
		return nil // nothing to do, no error
	} else if err != nil {
		return fmt.Errorf("failed to get PVC for processing: %w", err)
	}

	if !p.isOurPVC(pvc) {
		return nil
	}

	if pvc.Status.Phase != "Pending" {
		// If the PVC is not pending, we don't need to provision anything
		return nil
	}

	storageClass, err := p.storageClassInformer.Lister().Get(*pvc.Spec.StorageClassName)
	if err != nil {
		return fmt.Errorf("could not get storage class: %w", err)
	}

	if storageClass.Provisioner != csiProvisionerServerName {
		// We're not responsible for this PVC. Can only happen if
		// controller-manager makes a mistake setting the annotations, but
		// we're bailing here anyways for safety.
		return nil
	}

	err = p.provisionPVC(pvc, storageClass)

	if err != nil {
		p.recorder.Eventf(pvc, v1.EventTypeWarning, "ProvisioningFailed", "Failed to provision PV: %v", err)
		return err
	}
	p.recorder.Eventf(pvc, v1.EventTypeNormal, "Provisioned", "Successfully provisioned PV")

	return nil
}

// provisionPVC creates the directory where the volume lives, sets a quota for
// the requested amount of storage and creates the PV object representing this
// new volume
func (p *csiProvisionerServer) provisionPVC(pvc *v1.PersistentVolumeClaim, storageClass *storagev1.StorageClass) error {
	claimRef, err := ref.GetReference(scheme.Scheme, pvc)
	if err != nil {
		return fmt.Errorf("failed to get reference to PVC: %w", err)
	}

	storageReq := pvc.Spec.Resources.Requests[v1.ResourceStorage]
	if storageReq.IsZero() {
		return fmt.Errorf("PVC is not requesting any storage, this is not supported")
	}
	capacity, ok := storageReq.AsInt64()
	if !ok {
		return fmt.Errorf("PVC requesting more than 2^63 bytes of storage, this is not supported")
	}

	volumeID := "pvc-" + string(pvc.ObjectMeta.UID)
	volumePath := p.volumePath(volumeID)
	var mountOptions []string

	p.logger.Infof("Creating local PV %s", volumeID)

	switch *pvc.Spec.VolumeMode {
	case "", v1.PersistentVolumeFilesystem:
		if err := os.Mkdir(volumePath, 0644); err != nil && !os.IsExist(err) {
			return fmt.Errorf("failed to create volume directory: %w", err)
		}
		files, err := os.ReadDir(volumePath)
		if err != nil {
			return fmt.Errorf("failed to list files in newly-created volume: %w", err)
		}
		if len(files) > 0 {
			return errors.New("newly-created volume already contains data, bailing")
		}
		if err := fsquota.SetQuota(volumePath, uint64(capacity), uint64(capacity)/inodeCapacityRatio); err != nil {
			return fmt.Errorf("failed to update quota: %w", err)
		}
		mountOptions = storageClass.MountOptions
	case v1.PersistentVolumeBlock:
		imageFile, err := os.OpenFile(volumePath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return fmt.Errorf("failed to create volume image: %w", err)
		}
		defer imageFile.Close()
		if err := unix.Fallocate(int(imageFile.Fd()), 0, 0, capacity); err != nil {
			return fmt.Errorf("failed to fallocate() volume image: %w", err)
		}
	default:
		return fmt.Errorf("VolumeMode \"%s\" is unsupported", *pvc.Spec.VolumeMode)
	}

	vol := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: volumeID,
			Annotations: map[string]string{
				"pv.kubernetes.io/provisioned-by": csiProvisionerServerName},
		},
		Spec: v1.PersistentVolumeSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			Capacity: v1.ResourceList{
				v1.ResourceStorage: storageReq, // We're always giving the exact amount
			},
			PersistentVolumeSource: v1.PersistentVolumeSource{
				CSI: &v1.CSIPersistentVolumeSource{
					Driver:       csiProvisionerServerName,
					VolumeHandle: volumeID,
				},
			},
			ClaimRef:   claimRef,
			VolumeMode: pvc.Spec.VolumeMode,
			NodeAffinity: &v1.VolumeNodeAffinity{
				Required: &v1.NodeSelector{
					NodeSelectorTerms: []v1.NodeSelectorTerm{
						{
							MatchExpressions: []v1.NodeSelectorRequirement{
								{
									Key:      "kubernetes.io/hostname",
									Operator: v1.NodeSelectorOpIn,
									Values:   []string{p.NodeName},
								},
							},
						},
					},
				},
			},
			StorageClassName:              *pvc.Spec.StorageClassName,
			MountOptions:                  mountOptions,
			PersistentVolumeReclaimPolicy: *storageClass.ReclaimPolicy,
		},
	}

	_, err = p.Kubernetes.CoreV1().PersistentVolumes().Create(context.Background(), vol, metav1.CreateOptions{})
	if err != nil && !apierrs.IsAlreadyExists(err) {
		return fmt.Errorf("failed to create PV object: %w", err)
	}
	return nil
}

// processPV looks at a single PV item from the queue and checks if it has been
// released and needs to be deleted. If yes it deletes the associated quota,
// directory and the PV object and logs the result to the recorder.
func (p *csiProvisionerServer) processPV(key string) error {
	_, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return fmt.Errorf("invalid resource key: %s", key)
	}
	pv, err := p.pvInformer.Lister().Get(name)
	if apierrs.IsNotFound(err) {
		return nil // nothing to do, no error
	} else if err != nil {
		return fmt.Errorf("failed to get PV for processing: %w", err)
	}

	if !p.isOurPV(pv) {
		return nil
	}
	if pv.Spec.PersistentVolumeReclaimPolicy != v1.PersistentVolumeReclaimDelete || pv.Status.Phase != "Released" {
		return nil
	}
	volumePath := p.volumePath(pv.Spec.CSI.VolumeHandle)

	// Log deletes for auditing purposes
	p.logger.Infof("Deleting persistent volume %s", pv.Spec.CSI.VolumeHandle)
	switch *pv.Spec.VolumeMode {
	case "", v1.PersistentVolumeFilesystem:
		if err := fsquota.SetQuota(volumePath, 0, 0); err != nil {
			// We record these here manually since a successful deletion
			// removes the PV we'd be attaching them to.
			p.recorder.Eventf(pv, v1.EventTypeWarning, "DeprovisioningFailed", "Failed to remove quota: %v", err)
			return fmt.Errorf("failed to remove quota: %w", err)
		}
		if err := os.RemoveAll(volumePath); err != nil && !os.IsNotExist(err) {
			p.recorder.Eventf(pv, v1.EventTypeWarning, "DeprovisioningFailed", "Failed to delete volume: %v", err)
			return fmt.Errorf("failed to delete volume: %w", err)
		}
	case v1.PersistentVolumeBlock:
		if err := os.Remove(volumePath); err != nil && !os.IsNotExist(err) {
			p.recorder.Eventf(pv, v1.EventTypeWarning, "DeprovisioningFailed", "Failed to delete volume: %v", err)
			return fmt.Errorf("failed to delete volume: %w", err)
		}
	default:
		p.recorder.Eventf(pv, v1.EventTypeWarning, "DeprovisioningFailed", "Invalid volume mode \"%v\"", *pv.Spec.VolumeMode)
		return fmt.Errorf("invalid volume mode \"%v\"", *pv.Spec.VolumeMode)
	}

	err = p.Kubernetes.CoreV1().PersistentVolumes().Delete(context.Background(), pv.Name, metav1.DeleteOptions{})
	if err != nil && !apierrs.IsNotFound(err) {
		p.recorder.Eventf(pv, v1.EventTypeWarning, "DeprovisioningFailed", "Failed to delete PV object from K8s API: %v", err)
		return fmt.Errorf("failed to delete PV object: %w", err)
	}
	return nil
}
