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

// Package clusternet implements a WireGuard-based overlay network for Kubernetes. It relies on controller-manager's
// IPAM to assign IP ranges to nodes and on Kubernetes' Node objects to distribute the Node IPs and public keys.
//
// It sets up a single WireGuard network interface and routes the entire ClusterCIDR into that network interface,
// relying on WireGuard's AllowedIPs mechanism to look up the correct peer node to send the traffic to. This means
// that the routing table doesn't change and doesn't have to be separately managed. When clusternet is started
// it annotates its WireGuard public key onto its node object.
// For each node object that's created or updated on the K8s apiserver it checks if a public key annotation is set and
// if yes a peer with that public key, its InternalIP as endpoint and the CIDR for that node as AllowedIPs is created.
package clusternet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/vishvananda/netlink"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	common "git.monogon.dev/source/nexantic.git/metropolis/node"
	"git.monogon.dev/source/nexantic.git/metropolis/node/common/jsonpatch"
	"git.monogon.dev/source/nexantic.git/metropolis/node/common/supervisor"
	"git.monogon.dev/source/nexantic.git/metropolis/node/core/localstorage"
	"git.monogon.dev/source/nexantic.git/metropolis/node/core/logtree"
)

const (
	clusterNetDeviceName = "clusternet"
	publicKeyAnnotation  = "node.smalltown.nexantic.com/wg-pubkey"
)

type Service struct {
	NodeName        string
	Kubernetes      kubernetes.Interface
	ClusterNet      net.IPNet
	InformerFactory informers.SharedInformerFactory
	DataDirectory   *localstorage.DataKubernetesClusterNetworkingDirectory

	wgClient *wgctrl.Client
	privKey  wgtypes.Key
	logger   logtree.LeveledLogger
}

// ensureNode creates/updates the corresponding WireGuard peer entry for the given node objet
func (s *Service) ensureNode(newNode *corev1.Node) error {
	if newNode.Name == s.NodeName {
		// Node doesn't need to connect to itself
		return nil
	}
	pubKeyRaw := newNode.Annotations[publicKeyAnnotation]
	if pubKeyRaw == "" {
		return nil
	}
	pubKey, err := wgtypes.ParseKey(pubKeyRaw)
	if err != nil {
		return fmt.Errorf("failed to parse public-key annotation: %w", err)
	}
	var internalIP net.IP
	for _, addr := range newNode.Status.Addresses {
		if addr.Type == corev1.NodeInternalIP {
			if internalIP != nil {
				s.logger.Warningf("More than one NodeInternalIP specified, using the first one")
				break
			}
			internalIP = net.ParseIP(addr.Address)
			if internalIP == nil {
				s.logger.Warningf("Failed to parse Internal IP %s", addr.Address)
			}
		}
	}
	if internalIP == nil {
		return errors.New("node has no Internal IP")
	}
	var allowedIPs []net.IPNet
	for _, podNetStr := range newNode.Spec.PodCIDRs {
		_, podNet, err := net.ParseCIDR(podNetStr)
		if err != nil {
			s.logger.Warningf("Node %s PodCIDR failed to parse, ignored: %v", newNode.Name, err)
			continue
		}
		allowedIPs = append(allowedIPs, *podNet)
	}
	allowedIPs = append(allowedIPs, net.IPNet{IP: internalIP, Mask: net.CIDRMask(32, 32)})
	s.logger.V(1).Infof("Adding/Updating WireGuard peer node %s, endpoint %s, allowedIPs %+v", newNode.Name, internalIP.String(), allowedIPs)
	// WireGuard's kernel side has create/update semantics on peers by default. So we can just add the peer multiple
	// times to update it.
	err = s.wgClient.ConfigureDevice(clusterNetDeviceName, wgtypes.Config{
		Peers: []wgtypes.PeerConfig{{
			PublicKey:         pubKey,
			Endpoint:          &net.UDPAddr{Port: common.WireGuardPort, IP: internalIP},
			ReplaceAllowedIPs: true,
			AllowedIPs:        allowedIPs,
		}},
	})
	if err != nil {
		return fmt.Errorf("failed to add WireGuard peer node: %w", err)
	}
	return nil
}

// removeNode removes the corresponding WireGuard peer entry for the given node object
func (s *Service) removeNode(oldNode *corev1.Node) error {
	if oldNode.Name == s.NodeName {
		// Node doesn't need to connect to itself
		return nil
	}
	pubKeyRaw := oldNode.Annotations[publicKeyAnnotation]
	if pubKeyRaw == "" {
		return nil
	}
	pubKey, err := wgtypes.ParseKey(pubKeyRaw)
	if err != nil {
		return fmt.Errorf("node public-key annotation not decodable: %w", err)
	}
	err = s.wgClient.ConfigureDevice(clusterNetDeviceName, wgtypes.Config{
		Peers: []wgtypes.PeerConfig{{
			PublicKey: pubKey,
			Remove:    true,
		}},
	})
	if err != nil {
		return fmt.Errorf("failed to remove WireGuard peer node: %w", err)
	}
	return nil
}

// ensureOnDiskKey loads the private key from disk or (if none exists) generates one and persists it.
func (s *Service) ensureOnDiskKey() error {
	keyRaw, err := s.DataDirectory.Key.Read()
	if os.IsNotExist(err) {
		key, err := wgtypes.GeneratePrivateKey()
		if err != nil {
			return fmt.Errorf("failed to generate private key: %w", err)
		}
		if err := s.DataDirectory.Key.Write([]byte(key.String()), 0600); err != nil {
			return fmt.Errorf("failed to store newly generated key: %w", err)
		}

		s.privKey = key
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to load on-disk key: %w", err)
	}

	key, err := wgtypes.ParseKey(string(keyRaw))
	if err != nil {
		return fmt.Errorf("invalid private key in file: %w", err)
	}
	s.privKey = key
	return nil
}

// annotateThisNode annotates the node (as defined by NodeName) with the wireguard public key of this node.
func (s *Service) annotateThisNode(ctx context.Context) error {
	patch := []jsonpatch.JsonPatchOp{{
		Operation: "add",
		Path:      "/metadata/annotations/" + jsonpatch.EncodeJSONRefToken(publicKeyAnnotation),
		Value:     s.privKey.PublicKey().String(),
	}}

	patchRaw, err := json.Marshal(patch)
	if err != nil {
		return fmt.Errorf("failed to encode JSONPatch: %w", err)
	}

	if _, err := s.Kubernetes.CoreV1().Nodes().Patch(ctx, s.NodeName, types.JSONPatchType, patchRaw, metav1.PatchOptions{}); err != nil {
		return fmt.Errorf("failed to patch resource: %w", err)
	}

	return nil
}

// Run runs the ClusterNet service. See package description for what it does.
func (s *Service) Run(ctx context.Context) error {
	logger := supervisor.Logger(ctx)
	s.logger = logger

	wgClient, err := wgctrl.New()
	if err != nil {
		return fmt.Errorf("failed to connect to netlink's WireGuard config endpoint: %w", err)
	}
	s.wgClient = wgClient

	if err := s.ensureOnDiskKey(); err != nil {
		return fmt.Errorf("failed to ensure on-disk key: %w", err)
	}

	wgInterface := &Wireguard{LinkAttrs: netlink.LinkAttrs{Name: clusterNetDeviceName, Flags: net.FlagUp}}
	if err := netlink.LinkAdd(wgInterface); err != nil {
		return fmt.Errorf("failed to add WireGuard network interfacee: %w", err)
	}
	defer netlink.LinkDel(wgInterface)

	listenPort := common.WireGuardPort
	if err := wgClient.ConfigureDevice(clusterNetDeviceName, wgtypes.Config{
		PrivateKey: &s.privKey,
		ListenPort: &listenPort,
	}); err != nil {
		return fmt.Errorf("failed to set up WireGuard interface: %w", err)
	}

	if err := netlink.RouteAdd(&netlink.Route{
		Dst:       &s.ClusterNet,
		LinkIndex: wgInterface.Index,
	}); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to add cluster net route to Wireguard interface: %w", err)
	}

	if err := s.annotateThisNode(ctx); err != nil {
		return fmt.Errorf("when annotating this node with public key: %w", err)
	}

	nodeInformer := s.InformerFactory.Core().V1().Nodes()
	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(new interface{}) {
			newNode, ok := new.(*corev1.Node)
			if !ok {
				logger.Errorf("Received non-node item %+v in node event handler", new)
				return
			}
			if err := s.ensureNode(newNode); err != nil {
				logger.Warningf("Failed to sync node: %v", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			newNode, ok := new.(*corev1.Node)
			if !ok {
				logger.Errorf("Received non-node item %+v in node event handler", new)
				return
			}
			if err := s.ensureNode(newNode); err != nil {
				logger.Warningf("Failed to sync node: %v", err)
			}
		},
		DeleteFunc: func(old interface{}) {
			oldNode, ok := old.(*corev1.Node)
			if !ok {
				logger.Errorf("Received non-node item %+v in node event handler", oldNode)
				return
			}
			if err := s.removeNode(oldNode); err != nil {
				logger.Warningf("Failed to sync node: %v", err)
			}
		},
	})

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	nodeInformer.Informer().Run(ctx.Done())
	return ctx.Err()
}
