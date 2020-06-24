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
	"io/ioutil"
	"net"
	"os"

	"github.com/vishvananda/netlink"
	"go.uber.org/zap"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/pkg/jsonpatch"
)

const (
	clusterNetDeviceName = "clusternet"
	publicKeyAnnotation  = "node.smalltown.nexantic.com/wg-pubkey"

	privateKeyPath = "/data/kubernetes/clusternet.key"
)

type clusternet struct {
	nodeName     string
	wgClient     *wgctrl.Client
	nodeInformer coreinformers.NodeInformer
	logger       *zap.Logger
}

// ensureNode creates/updates the corresponding WireGuard peer entry for the given node objet
func (c *clusternet) ensureNode(newNode *corev1.Node) error {
	if newNode.Name == c.nodeName {
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
				c.logger.Warn("More than one NodeInternalIP specified, using the first one")
				break
			}
			internalIP = net.ParseIP(addr.Address)
			if internalIP == nil {
				c.logger.Warn("failed to parse Internal IP")
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
			c.logger.Warn("Node PodCIDR failed to parse, ignored", zap.Error(err), zap.String("node", newNode.Name))
			continue
		}
		allowedIPs = append(allowedIPs, *podNet)
	}
	c.logger.Debug("Adding/Updating WireGuard peer node", zap.String("node", newNode.Name),
		zap.String("endpointIP", internalIP.String()), zap.Any("allowedIPs", allowedIPs))
	// WireGuard's kernel side has create/update semantics on peers by default. So we can just add the peer multiple
	// times to update it.
	err = c.wgClient.ConfigureDevice(clusterNetDeviceName, wgtypes.Config{
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
func (c *clusternet) removeNode(oldNode *corev1.Node) error {
	if oldNode.Name == c.nodeName {
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
	err = c.wgClient.ConfigureDevice(clusterNetDeviceName, wgtypes.Config{
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

// EnsureOnDiskKey loads the private key from disk or (if none exists) generates one and persists it.
func EnsureOnDiskKey() (*wgtypes.Key, error) {
	privKeyRaw, err := ioutil.ReadFile(privateKeyPath)
	if os.IsNotExist(err) {
		privKey, err := wgtypes.GeneratePrivateKey()
		if err != nil {
			return nil, fmt.Errorf("failed to generate private key: %w", err)
		}
		if err := ioutil.WriteFile(privateKeyPath, []byte(privKey.String()), 0600); err != nil {
			return nil, fmt.Errorf("failed to store newly generated key: %w", err)
		}
		return &privKey, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to load on-disk key: %w", err)
	}
	privKey, err := wgtypes.ParseKey(string(privKeyRaw))
	if err != nil {
		return nil, fmt.Errorf("invalid private key in file: %w", err)
	}
	return &privKey, nil
}

// Run runs the ClusterNet service. See package description for what it does.
func Run(informerFactory informers.SharedInformerFactory, clusterNet net.IPNet, clientSet kubernetes.Interface, key *wgtypes.Key) supervisor.Runnable {
	return func(ctx context.Context) error {
		logger := supervisor.Logger(ctx)
		nodeName, err := os.Hostname()
		if err != nil {
			return fmt.Errorf("failed to determine hostname: %w", err)
		}
		wgClient, err := wgctrl.New()
		if err != nil {
			return fmt.Errorf("failed to connect to netlink's WireGuard config endpoint: %w", err)
		}

		nodeAnnotationPatch := []jsonpatch.JsonPatchOp{{
			Operation: "add",
			Path:      "/metadata/annotations/" + jsonpatch.EncodeJSONRefToken(publicKeyAnnotation),
			Value:     key.PublicKey().String(),
		}}

		nodeAnnotationPatchRaw, err := json.Marshal(nodeAnnotationPatch)
		if err != nil {
			return fmt.Errorf("failed to encode JSONPatch: %w", err)
		}

		if _, err := clientSet.CoreV1().Nodes().Patch(ctx, nodeName, types.JSONPatchType, nodeAnnotationPatchRaw, metav1.PatchOptions{}); err != nil {
			return fmt.Errorf("failed to set ClusterNet public key for node: %w", err)
		}

		nodeInformer := informerFactory.Core().V1().Nodes()
		wgInterface := &Wireguard{LinkAttrs: netlink.LinkAttrs{Name: clusterNetDeviceName, Flags: net.FlagUp}}
		if err := netlink.LinkAdd(wgInterface); err != nil {
			return fmt.Errorf("failed to add WireGuard network interfacee: %w", err)
		}
		defer netlink.LinkDel(wgInterface)

		listenPort := common.WireGuardPort
		if err := wgClient.ConfigureDevice(clusterNetDeviceName, wgtypes.Config{
			PrivateKey: key,
			ListenPort: &listenPort,
		}); err != nil {
			return fmt.Errorf("failed to set up WireGuard interface: %w", err)
		}

		if err := netlink.RouteAdd(&netlink.Route{
			Dst:       &clusterNet,
			LinkIndex: wgInterface.Index,
		}); err != nil && !os.IsExist(err) {
			return fmt.Errorf("failed to add cluster net route to Wireguard interface: %w", err)
		}

		cnet := clusternet{
			wgClient:     wgClient,
			nodeInformer: nodeInformer,
			logger:       logger,
			nodeName:     nodeName,
		}

		nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc: func(new interface{}) {
				newNode, ok := new.(*corev1.Node)
				if !ok {
					logger.Error("Received non-node item in node event handler", zap.Reflect("item", new))
					return
				}
				if err := cnet.ensureNode(newNode); err != nil {
					logger.Warn("Failed to sync node", zap.Error(err))
				}
			},
			UpdateFunc: func(old, new interface{}) {
				newNode, ok := new.(*corev1.Node)
				if !ok {
					logger.Error("Received non-node item in node event handler", zap.Reflect("item", new))
					return
				}
				if err := cnet.ensureNode(newNode); err != nil {
					logger.Warn("Failed to sync node", zap.Error(err))
				}
			},
			DeleteFunc: func(old interface{}) {
				oldNode, ok := old.(*corev1.Node)
				if !ok {
					logger.Error("Received non-node item in node event handler", zap.Reflect("item", oldNode))
					return
				}
				if err := cnet.removeNode(oldNode); err != nil {
					logger.Warn("Failed to sync node", zap.Error(err))
				}
			},
		})
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		nodeInformer.Informer().Run(ctx.Done())
		return ctx.Err()
	}
}
