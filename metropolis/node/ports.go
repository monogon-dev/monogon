// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package node

import (
	"strconv"
)

// Port is a TCP and/or UDP port number reserved for and used by Metropolis
// node code.
type Port uint16

const (
	// CuratorServicePort is the TCP port on which the Curator listens for gRPC
	// calls and services Management/AAA/Curator RPCs.
	CuratorServicePort Port = 7835
	// ConsensusPort is the TCP port on which etcd listens for peer traffic.
	ConsensusPort Port = 7834
	// DebugServicePort is the TCP port on which the debug service serves gRPC
	// traffic. This is only available in debug builds.
	DebugServicePort Port = 7837
	// WireGuardPort is the UDP port on which the Wireguard Kubernetes network
	// overlay listens for incoming peer traffic.
	WireGuardPort Port = 7838
	// NodeManagementPort is the TCP port on which the node-local management service
	// serves gRPC traffic for NodeManagement.
	NodeManagementPort Port = 7839
	// MetricsPort is the TCP port on which the Metrics Service exports
	// Prometheus-compatible metrics for this node, secured using TLS and the
	// Cluster/Node certificates.
	MetricsPort Port = 7840
	// MetricsNodeListenerPort is the TCP port on which the Prometheus node_exporter
	// runs, bound to 127.0.0.1. The Metrics Service proxies traffic to it from the
	// public MetricsPort.
	MetricsNodeListenerPort Port = 7841
	// MetricsEtcdListenerPort is the TCP port on which the etcd exporter
	// runs, bound to 127.0.0.1. The metrics service proxies traffic to it from the
	// public MetricsPort.
	MetricsEtcdListenerPort Port = 7842
	// MetricsKubeSchedulerListenerPort is the TCP port on which the proxy for
	// the kube-scheduler runs, bound to 127.0.0.1. The metrics service proxies
	// traffic to it from the public MetricsPort.
	MetricsKubeSchedulerListenerPort Port = 7843
	// MetricsKubeControllerManagerListenerPort is the TCP port on which the
	// proxy for the controller-manager runs, bound to 127.0.0.1. The metrics
	// service proxies traffic to it from the public MetricsPort.
	MetricsKubeControllerManagerListenerPort Port = 7844
	// MetricsKubeAPIServerListenerPort is the TCP port on which the
	// proxy for the api-server runs, bound to 127.0.0.1. The metrics
	// service proxies traffic to it from the public MetricsPort.
	MetricsKubeAPIServerListenerPort Port = 7845
	// MetricsContainerdListenerPort is the TCP port on which the
	// containerd metrics endpoint, bound to 127.0.0.1, is exposed.
	MetricsContainerdListenerPort Port = 7846
	// KubernetesAPIPort is the TCP port on which the Kubernetes API is
	// exposed.
	KubernetesAPIPort Port = 6443
	// KubernetesAPIWrappedPort is the TCP port on which the Metropolis
	// authenticating proxy for the Kubernetes API is exposed.
	KubernetesAPIWrappedPort Port = 6444
	// KubernetesWorkerLocalAPIPort is the TCP port on which Kubernetes worker nodes
	// run a loadbalancer to access the cluster's API servers before cluster
	// networking is available. This port is only bound to 127.0.0.1.
	KubernetesWorkerLocalAPIPort Port = 6445
	// DebuggerPort is the port on which the delve debugger runs (on debug
	// builds only). Not to be confused with DebugServicePort.
	DebuggerPort Port = 2345
)

var SystemPorts = []Port{
	CuratorServicePort,
	ConsensusPort,
	DebugServicePort,
	WireGuardPort,
	NodeManagementPort,
	MetricsPort,
	MetricsNodeListenerPort,
	MetricsEtcdListenerPort,
	MetricsKubeSchedulerListenerPort,
	MetricsKubeControllerManagerListenerPort,
	MetricsKubeAPIServerListenerPort,
	MetricsContainerdListenerPort,
	KubernetesAPIPort,
	KubernetesAPIWrappedPort,
	KubernetesWorkerLocalAPIPort,
	DebuggerPort,
}

func (p Port) String() string {
	switch p {
	case CuratorServicePort:
		return "curator"
	case ConsensusPort:
		return "consensus"
	case DebugServicePort:
		return "debug"
	case WireGuardPort:
		return "wireguard"
	case NodeManagementPort:
		return "node-mgmt"
	case MetricsPort:
		return "metrics"
	case MetricsNodeListenerPort:
		return "metrics-node-exporter"
	case MetricsEtcdListenerPort:
		return "metrics-etcd"
	case MetricsKubeSchedulerListenerPort:
		return "metrics-kubernetes-scheduler"
	case MetricsKubeControllerManagerListenerPort:
		return "metrics-kubernetes-controller-manager"
	case MetricsKubeAPIServerListenerPort:
		return "metrics-kubernetes-api-server"
	case MetricsContainerdListenerPort:
		return "metrics-containerd"
	case KubernetesAPIPort:
		return "kubernetes-api"
	case KubernetesAPIWrappedPort:
		return "kubernetes-api-wrapped"
	case KubernetesWorkerLocalAPIPort:
		return "kubernetes-worker-local-api"
	case DebuggerPort:
		return "delve"
	}
	return "unknown"
}

func (p Port) PortString() string {
	return strconv.Itoa(int(p))
}
