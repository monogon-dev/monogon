// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// This package tests the network policy controller, using the Cyclonus network
// policy conformance test suite. It uses the real policy controller and
// nftables controller, but fake Kubernetes API and pods, which allows the tests
// to run much faster than would be possible with the real API server and pods.
//
// By default, the test runs in a ktest, which means we are testing with the
// same kernel version, and thus nftables implementation, which is used in
// Monogon OS. But you can also run tests manually in a new user and network
// namespace, which allows you to use all the debugging tools available on your
// machine. Useful commands:
//
//	bazel build //metropolis/node/kubernetes/networkpolicy:networkpolicy_test
//		Build the test.
//	unshare --map-user=0 --net
//		Create a user and network namespace.
//	tcpdump -i any &
//		Capture traffic in the background.
//	IN_KTEST=true bazel-bin/metropolis/node/kubernetes/networkpolicy/networkpolicy_test_/networkpolicy_test
//		Run the test.
package networkpolicy_test

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"git.dolansoft.org/dolansoft/k8s-nft-npc/nftctrl"
	"github.com/mattfenwick/cyclonus/pkg/connectivity"
	"github.com/mattfenwick/cyclonus/pkg/connectivity/probe"
	"github.com/mattfenwick/cyclonus/pkg/generator"
	"github.com/mattfenwick/cyclonus/pkg/kube"
	"github.com/mattfenwick/cyclonus/pkg/matcher"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

// Change this to 1 to test IPv6 instead of IPv4. Tests pass with IPv6, but run
// much slower, apparently because some ICMPv6 packets are dropped and need to
// be retransmitted.
//
// TODO: Fix slow IPv6 tests and run them automatically, e.g. as a separate
// ktest with a flag to select IPv6.
const testIPIndex = 0

const podIfaceGroup uint32 = 8

// Loopback traffic is only affected by network policies when connecting through
// a ClusterIP service. For now, we don't simulate service IPs, so this is
// disabled.
const ignoreLoopback = true

var serverPorts = []int{80, 81}
var serverProtocols = []corev1.Protocol{corev1.ProtocolTCP, corev1.ProtocolUDP}

var gatewayIPv4 = net.ParseIP("10.8.0.1").To4()
var gatewayIPv6 = net.ParseIP("fd08::1")

func netListen(port int, protocol corev1.Protocol, payload string, stop <-chan struct{}) error {
	switch protocol {
	case corev1.ProtocolTCP:
		listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: port})
		if err != nil {
			return err
		}
		go func() {
			<-stop
			listener.Close()
		}()
		for {
			conn, err := listener.Accept()
			if err != nil {
				return err
			}
			conn.Write([]byte(payload))
			conn.Close()
		}
	case corev1.ProtocolUDP:
		conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: port})
		if err != nil {
			return err
		}
		go func() {
			<-stop
			conn.Close()
		}()
		var buf [16]byte
		for {
			_, clientAddr, err := conn.ReadFrom(buf[:])
			if err != nil {
				return err
			}
			conn.WriteTo([]byte(payload), clientAddr)
		}
	default:
		return fmt.Errorf("unsupported protocol: %q", protocol)
	}
}

func netConnect(ip net.IP, port int, protocol corev1.Protocol, expectPayload string) (error, error) {
	addr := (&net.TCPAddr{IP: ip, Port: port}).String()
	payload := make([]byte, len(expectPayload))
	switch protocol {
	case corev1.ProtocolTCP:
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if errors.Is(err, unix.EHOSTUNREACH) || errors.Is(err, unix.EACCES) {
			return err, nil
		} else if err != nil {
			return nil, err
		}
		defer conn.Close()
		_, err = conn.Read(payload)
		if err != nil {
			return nil, err
		}
	case corev1.ProtocolUDP:
		conn, err := net.Dial("udp", addr)
		if err != nil {
			return nil, err
		}
		defer conn.Close()
		err = conn.SetDeadline(time.Now().Add(2 * time.Second))
		if err != nil {
			return nil, err
		}
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			return nil, err
		}
		_, err = conn.Read(payload)
		if errors.Is(err, unix.EHOSTUNREACH) || errors.Is(err, unix.EACCES) {
			return err, nil
		} else if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported protocol: %q", protocol)
	}
	if string(payload) != expectPayload {
		return nil, fmt.Errorf("wrong payload, expected %q, got %q", expectPayload, payload)
	}
	return nil, nil
}

type fakePod struct {
	name     string
	ips      []podIP
	netns    *os.File
	rawNetns syscall.RawConn
	stopCh   chan struct{}
}

type podIP struct {
	ip        net.IP
	gatewayIP net.IP
	zeroIP    net.IP
	fullMask  net.IPMask
	zeroMask  net.IPMask
}

func createFakePod(namespace, name string, ips []net.IP) (*fakePod, error) {
	p := &fakePod{
		name:   fmt.Sprintf("%s/%s", namespace, name),
		ips:    make([]podIP, len(ips)),
		stopCh: make(chan struct{}),
	}
	for i, ip := range ips {
		p := &p.ips[i]
		switch {
		case ip.To4() != nil:
			p.ip = ip.To4()
			p.gatewayIP = gatewayIPv4
			p.zeroIP = net.ParseIP("0.0.0.0").To4()
			p.fullMask = net.CIDRMask(32, 32)
			p.zeroMask = net.CIDRMask(0, 32)
		case ip.To16() != nil:
			p.ip = ip.To16()
			p.gatewayIP = gatewayIPv6
			p.zeroIP = net.ParseIP("::")
			p.fullMask = net.CIDRMask(128, 128)
			p.zeroMask = net.CIDRMask(0, 128)
		default:
			return nil, fmt.Errorf("invalid IP: %v", ip)
		}
	}

	// Create new network namespace.
	runLockedThread(func() {
		err := unix.Unshare(unix.CLONE_NEWNET)
		if err != nil {
			panic(err)
		}
		// Obtain an FD of the new net namespace.
		p.netns, err = os.Open(fmt.Sprintf("/proc/%d/task/%d/ns/net", os.Getpid(), unix.Gettid()))
		if err != nil {
			panic(err)
		}
	})

	var err error
	p.rawNetns, err = p.netns.SyscallConn()
	if err != nil {
		panic(err)
	}

	// Create veth pair.
	linkAttrs := netlink.NewLinkAttrs()
	linkAttrs.Name = fmt.Sprintf("veth_%s_%s", namespace, name)
	if len(linkAttrs.Name) > 15 {
		hash := sha256.Sum256([]byte(linkAttrs.Name))
		linkAttrs.Name = fmt.Sprintf("veth_%x", hash[:5])
	}
	linkAttrs.Group = podIfaceGroup
	linkAttrs.Flags = net.FlagUp
	p.rawNetns.Control(func(fd uintptr) {
		err = netlink.LinkAdd(&netlink.Veth{LinkAttrs: linkAttrs, PeerName: "veth", PeerNamespace: netlink.NsFd(fd)})
	})
	if err != nil {
		return nil, err
	}

	veth, err := netlink.LinkByName(linkAttrs.Name)
	if err != nil {
		return nil, err
	}

	for _, ip := range p.ips {
		// Add gateway address to link.
		addr := &netlink.Addr{IPNet: &net.IPNet{IP: ip.gatewayIP, Mask: ip.fullMask}, Flags: unix.IFA_F_NODAD}
		err = netlink.AddrAdd(veth, addr)
		if err != nil {
			return nil, err
		}

		// Add route.
		err = netlink.RouteAdd(&netlink.Route{
			LinkIndex: veth.Attrs().Index,
			Scope:     netlink.SCOPE_HOST,
			Dst:       &net.IPNet{IP: ip.ip, Mask: ip.fullMask},
		})
		if err != nil {
			return nil, err
		}
	}

	p.runInNetns(func() {
		// Enable loopback traffic in the pod.
		loopback, err := netlink.LinkByName("lo")
		if err != nil {
			panic(err)
		}
		err = netlink.LinkSetUp(loopback)
		if err != nil {
			panic(err)
		}

		// Set up the veth in the pod namespace.
		veth, err := netlink.LinkByName("veth")
		if err != nil {
			panic(err)
		}
		err = netlink.LinkSetUp(veth)
		if err != nil {
			panic(err)
		}

		for _, ip := range p.ips {
			// IFA_F_NODAD disables duplicate address detection for IPv6.
			addr := &netlink.Addr{IPNet: &net.IPNet{IP: ip.ip, Mask: ip.fullMask}, Flags: unix.IFA_F_NODAD}
			err = netlink.AddrAdd(veth, addr)
			if err != nil {
				panic(err)
			}

			err = netlink.RouteAdd(&netlink.Route{
				LinkIndex: veth.Attrs().Index,
				Scope:     netlink.SCOPE_LINK,
				Dst:       &net.IPNet{IP: ip.gatewayIP, Mask: ip.fullMask},
				Src:       ip.ip,
			})
			if err != nil {
				panic(err)
			}
			err = netlink.RouteAdd(&netlink.Route{
				LinkIndex: veth.Attrs().Index,
				Scope:     netlink.SCOPE_UNIVERSE,
				Dst:       &net.IPNet{IP: ip.zeroIP, Mask: ip.zeroMask},
				Gw:        ip.gatewayIP,
			})
			if err != nil {
				panic(err)
			}
		}
	})

	for _, protocol := range serverProtocols {
		for _, port := range serverPorts {
			go p.runInNetns(func() {
				err := netListen(port, protocol, p.name, p.stopCh)
				select {
				case <-p.stopCh:
				default:
					panic(err)
				}
			})
		}
	}

	return p, nil
}

func (p *fakePod) stop() {
	close(p.stopCh)
	p.netns.Close()
}

func (p *fakePod) runInNetns(f func()) {
	runLockedThread(func() {
		p.rawNetns.Control(func(fd uintptr) {
			err := unix.Setns(int(fd), unix.CLONE_NEWNET)
			if err != nil {
				panic(err)
			}
		})
		f()
	})
}

func (p *fakePod) connect(target *fakePod, port int, protocol corev1.Protocol) (error, error) {
	var errConnectivity, err error
	p.runInNetns(func() {
		errConnectivity, err = netConnect(target.ips[testIPIndex].ip, port, protocol, target.name)
	})
	return errConnectivity, err
}

// runLockedThread runs f locked on an OS thread, which allows f to switch to a
// different network namespace. After f returns, the previous network namespace
// is restored, and if that succeeds, the thread is unlocked, allowing the
// thread to be reused for other goroutines.
func runLockedThread(f func()) {
	runtime.LockOSThread()
	oldNs, err := os.Open(fmt.Sprintf("/proc/%d/task/%d/ns/net", os.Getpid(), unix.Gettid()))
	if err != nil {
		panic(err)
	}
	defer oldNs.Close()
	f()
	err = unix.Setns(int(oldNs.Fd()), unix.CLONE_NEWNET)
	if err == nil {
		runtime.UnlockOSThread()
	}
}

type fakeKubernetes struct {
	*kube.MockKubernetes
	nft       *nftctrl.Controller
	fakePods  map[string]*fakePod
	nextPodIP uint64
}

func (k *fakeKubernetes) CreateNamespace(kubeNamespace *corev1.Namespace) (*corev1.Namespace, error) {
	kubeNamespace, err := k.MockKubernetes.CreateNamespace(kubeNamespace)
	if err == nil && k.nft != nil {
		k.nft.SetNamespace(kubeNamespace.Name, kubeNamespace)
		if err := k.nft.Flush(); err != nil {
			return nil, err
		}
	}
	return kubeNamespace, err
}

func (k *fakeKubernetes) SetNamespaceLabels(namespace string, labels map[string]string) (*corev1.Namespace, error) {
	kubeNamespace, err := k.MockKubernetes.SetNamespaceLabels(namespace, labels)
	if err == nil && k.nft != nil {
		k.nft.SetNamespace(namespace, kubeNamespace)
		if err := k.nft.Flush(); err != nil {
			return nil, err
		}
	}
	return kubeNamespace, err
}

func (k *fakeKubernetes) DeleteNamespace(namespace string) error {
	err := k.MockKubernetes.DeleteNamespace(namespace)
	if err == nil && k.nft != nil {
		k.nft.SetNamespace(namespace, nil)
		if err := k.nft.Flush(); err != nil {
			return err
		}
	}
	return err
}

func (k *fakeKubernetes) CreateNetworkPolicy(kubePolicy *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
	kubePolicy, err := k.MockKubernetes.CreateNetworkPolicy(kubePolicy)
	if err == nil && k.nft != nil {
		k.nft.SetNetworkPolicy(cache.MetaObjectToName(kubePolicy), kubePolicy)
		if err := k.nft.Flush(); err != nil {
			return nil, err
		}
	}
	return kubePolicy, err
}

func (k *fakeKubernetes) UpdateNetworkPolicy(kubePolicy *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {
	kubePolicy, err := k.MockKubernetes.UpdateNetworkPolicy(kubePolicy)
	if err == nil && k.nft != nil {
		k.nft.SetNetworkPolicy(cache.MetaObjectToName(kubePolicy), kubePolicy)
		if err := k.nft.Flush(); err != nil {
			return nil, err
		}
	}
	return kubePolicy, err
}

func (k *fakeKubernetes) DeleteNetworkPolicy(namespace string, name string) error {
	err := k.MockKubernetes.DeleteNetworkPolicy(namespace, name)
	if err == nil && k.nft != nil {
		k.nft.SetNetworkPolicy(cache.NewObjectName(namespace, name), nil)
		if err := k.nft.Flush(); err != nil {
			return err
		}
	}
	return err
}

func (k *fakeKubernetes) DeleteAllNetworkPoliciesInNamespace(namespace string) error {
	policies, err := k.GetNetworkPoliciesInNamespace(namespace)
	if err == nil && k.nft != nil {
		for _, kubePolicy := range policies {
			k.nft.SetNetworkPolicy(cache.MetaObjectToName(&kubePolicy), nil)
			if err := k.nft.Flush(); err != nil {
				return err
			}
		}
	}
	return k.MockKubernetes.DeleteAllNetworkPoliciesInNamespace(namespace)
}

func (k *fakeKubernetes) CreatePod(kubePod *corev1.Pod) (*corev1.Pod, error) {
	kubePod, err := k.MockKubernetes.CreatePod(kubePod)
	objectName := cache.MetaObjectToName(kubePod)
	if err == nil {
		ipSuffix := []byte{uint8(k.nextPodIP >> 8), uint8(k.nextPodIP)}
		podIPv4 := slices.Concat(gatewayIPv4[:2], ipSuffix)
		podIPv6 := slices.Concat(gatewayIPv6[:14], ipSuffix)
		k.nextPodIP++
		kubePod.Status.PodIPs = []corev1.PodIP{{IP: podIPv4.String()}, {IP: podIPv6.String()}}
		kubePod.Status.PodIP = kubePod.Status.PodIPs[testIPIndex].IP
		ips := []net.IP{podIPv4, podIPv6}
		fakePod, err := createFakePod(kubePod.Namespace, kubePod.Name, ips)
		if err != nil {
			return nil, fmt.Errorf("failed to create fake pod: %w", err)
		}
		k.fakePods[objectName.String()] = fakePod
	}
	if err == nil && k.nft != nil {
		k.nft.SetPod(objectName, kubePod)
		if err := k.nft.Flush(); err != nil {
			return nil, err
		}
	}
	return kubePod, err
}

func (k *fakeKubernetes) DeletePod(namespace string, pod string) error {
	err := k.MockKubernetes.DeletePod(namespace, pod)
	objectName := cache.NewObjectName(namespace, pod)
	if err == nil {
		k.fakePods[objectName.String()].stop()
		delete(k.fakePods, objectName.String())
	}
	if err == nil && k.nft != nil {
		k.nft.SetPod(objectName, nil)
		if err := k.nft.Flush(); err != nil {
			return err
		}
	}
	return err
}

func (k *fakeKubernetes) SetPodLabels(namespace string, pod string, labels map[string]string) (*corev1.Pod, error) {
	kubePod, err := k.MockKubernetes.SetPodLabels(namespace, pod, labels)
	if err == nil && k.nft != nil {
		k.nft.SetPod(cache.MetaObjectToName(kubePod), kubePod)
		if err := k.nft.Flush(); err != nil {
			return nil, err
		}
	}
	return kubePod, err
}

var protocolMap = map[string]corev1.Protocol{
	"tcp":  corev1.ProtocolTCP,
	"udp":  corev1.ProtocolUDP,
	"sctp": corev1.ProtocolSCTP,
}

func (k *fakeKubernetes) ExecuteRemoteCommand(namespace string, pod string, container string, command []string) (string, string, error, error) {
	// command is expected to have this format:
	// /agnhost connect s-x-a.x.svc.cluster.local:80 --timeout=1s --protocol=tcp
	if len(command) != 5 {
		return "", "", nil, fmt.Errorf("unexpected command length: %v", command)
	}
	targetService, targetPortStr, ok := strings.Cut(command[2], ".svc.cluster.local:")
	if !ok {
		return "", "", nil, fmt.Errorf("failed to parse target: %q", command[2])
	}
	targetService, targetNamespace, ok := strings.Cut(targetService, ".")
	if !ok {
		return "", "", nil, fmt.Errorf("failed to parse target: %q", command[2])
	}
	targetPod := strings.TrimPrefix(targetService, fmt.Sprintf("s-%s-", targetNamespace))
	protocol := strings.TrimPrefix(command[4], "--protocol=")

	sourceFakePod := k.fakePods[cache.NewObjectName(namespace, pod).String()]
	targetFakePod := k.fakePods[cache.NewObjectName(targetNamespace, targetPod).String()]

	targetPort, err := strconv.Atoi(targetPortStr)
	if err != nil {
		return "", "", nil, err
	}
	kubeProtocol, ok := protocolMap[protocol]
	if !ok {
		return "", "", nil, fmt.Errorf("invalid protocol: %q", protocol)
	}

	connectErr, err := sourceFakePod.connect(targetFakePod, targetPort, kubeProtocol)
	return "", "", connectErr, err
}

func (k *fakeKubernetes) initializeNft() error {
	namespaces, err := k.GetAllNamespaces()
	if err != nil {
		return err
	}
	for _, kubeNamespace := range namespaces.Items {
		k.nft.SetNamespace(kubeNamespace.Name, &kubeNamespace)
		pods, err := k.GetPodsInNamespace(kubeNamespace.Name)
		if err != nil {
			return err
		}
		for _, kubePod := range pods {
			k.nft.SetPod(cache.MetaObjectToName(&kubePod), &kubePod)
		}
		policies, err := k.GetNetworkPoliciesInNamespace(kubeNamespace.Name)
		if err != nil {
			return err
		}
		if len(policies) != 0 {
			return fmt.Errorf("expected no initial policies, but found %s", cache.MetaObjectToName(&policies[0]).String())
		}
	}
	if err := k.nft.Flush(); err != nil {
		return fmt.Errorf("initial flush failed: %w", err)
	}
	return nil
}

type testRecorder struct {
	t *testing.T
}

func (t *testRecorder) Event(object kruntime.Object, eventtype, reason, message string) {
	if eventtype == corev1.EventTypeWarning {
		t.t.Errorf("Warning event: object %v, %s: %s", object, reason, message)
	} else {
		t.t.Logf("%s event: object %v, %s: %s", eventtype, object, reason, message)
	}
}

func (t *testRecorder) Eventf(object kruntime.Object, eventtype, reason, messageFmt string, args ...interface{}) {
	t.Event(object, eventtype, reason, fmt.Sprintf(messageFmt, args...))
}

func (t *testRecorder) AnnotatedEventf(object kruntime.Object, annotations map[string]string, eventtype, reason, messageFmt string, args ...interface{}) {
	t.Event(object, eventtype, reason, fmt.Sprintf(messageFmt, args...))
}

const maxNamespaceName = "ns-3456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012"
const maxPodName = "pod-456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012"
const maxPolicyName = "policy-789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012"

func setPolicyName(name string) generator.Setter {
	return func(policy *generator.Netpol) {
		policy.Name = name
	}
}

func extraTestCases() []*generator.TestCase {
	return []*generator.TestCase{
		{
			Description: "Update namespace so that additional peer selector matches",
			Tags:        generator.NewStringSet(generator.TagSetNamespaceLabels),
			Steps: []*generator.TestStep{
				generator.NewTestStep(generator.ProbeAllAvailable,
					generator.CreatePolicy(generator.BuildPolicy(generator.SetPeers(true, []networkingv1.NetworkPolicyPeer{
						{NamespaceSelector: &metav1.LabelSelector{MatchLabels: map[string]string{"ns": "y"}}},
						{NamespaceSelector: &metav1.LabelSelector{MatchLabels: map[string]string{"new-ns": "qrs"}}},
					})).NetworkPolicy())),
				generator.NewTestStep(generator.ProbeAllAvailable,
					generator.SetNamespaceLabels("y", map[string]string{"ns": "y", "new-ns": "qrs"})),
				generator.NewTestStep(generator.ProbeAllAvailable,
					generator.SetNamespaceLabels("y", map[string]string{"ns": "y"})),
			},
		},
		{
			Description: "Delete one of multiple policies",
			Tags:        generator.NewStringSet(generator.TagDeletePolicy),
			Steps: []*generator.TestStep{
				generator.NewTestStep(generator.ProbeAllAvailable,
					generator.CreatePolicy(generator.BuildPolicy(setPolicyName("policy1")).NetworkPolicy()),
					generator.CreatePolicy(generator.BuildPolicy(setPolicyName("policy2"), generator.SetPorts(true, nil)).NetworkPolicy()),
					generator.DeletePolicy("x", "policy1"),
				),
			},
		},
		{
			Description: "Create namespace, pod and policy with maximum name length",
			Tags:        generator.NewStringSet(generator.TagCreateNamespace),
			Steps: []*generator.TestStep{
				generator.NewTestStep(generator.ProbeAllAvailable,
					generator.CreateNamespace(maxNamespaceName, nil),
					generator.CreatePod(maxNamespaceName, maxPodName, nil),
					generator.CreatePolicy(generator.BuildPolicy(
						generator.SetNamespace(maxNamespaceName),
						setPolicyName(maxPolicyName),
						generator.SetPodSelector(metav1.LabelSelector{})).NetworkPolicy())),
				generator.NewTestStep(generator.ProbeAllAvailable,
					generator.DeleteNamespace(maxNamespaceName)),
			},
		},
	}
}

type initTestCase struct {
	description string
	init        func(*connectivity.TestCaseState, *nftctrl.Controller) error
}

func initializationTestCases() []initTestCase {
	return []initTestCase{
		{
			description: "Delete rule during initialization",
			// It is possible that we already receive updates/deletes during
			// initialization, if one of the resource types has not yet finished
			// loading. This means that we need to be able to handle such updates
			// before the first Flush.
			init: func(t *connectivity.TestCaseState, nft *nftctrl.Controller) error {
				namespaces, err := t.Kubernetes.GetAllNamespaces()
				if err != nil {
					return err
				}
				for _, kubeNamespace := range namespaces.Items {
					nft.SetNamespace(kubeNamespace.Name, &kubeNamespace)
					pods, err := t.Kubernetes.GetPodsInNamespace(kubeNamespace.Name)
					if err != nil {
						return err
					}
					for _, kubePod := range pods {
						nft.SetPod(cache.MetaObjectToName(&kubePod), &kubePod)
					}
				}

				policy1 := generator.BuildPolicy().NetworkPolicy()
				policy1.Name = "policy1"
				nft.SetNetworkPolicy(cache.MetaObjectToName(policy1), policy1)

				policy2 := generator.BuildPolicy(generator.SetPorts(true, nil)).NetworkPolicy()
				policy2.Name = "policy2"
				if err := t.CreatePolicy(policy2); err != nil {
					return fmt.Errorf("failed to create policy: %w", err)
				}
				nft.SetNetworkPolicy(cache.MetaObjectToName(policy2), policy2)

				// Delete policy 1.
				nft.SetNetworkPolicy(cache.MetaObjectToName(policy1), nil)

				return nil
			},
		},
		{
			description: "Initialize namespaces last",
			init: func(t *connectivity.TestCaseState, nft *nftctrl.Controller) error {
				policy := generator.BuildPolicy().NetworkPolicy()
				policy.Spec.Ingress[0].From[0].NamespaceSelector = nil
				if err := t.CreatePolicy(policy); err != nil {
					return fmt.Errorf("failed to create policy: %w", err)
				}

				namespaces, err := t.Kubernetes.GetAllNamespaces()
				if err != nil {
					return err
				}
				for _, kubeNamespace := range namespaces.Items {
					pods, err := t.Kubernetes.GetPodsInNamespace(kubeNamespace.Name)
					if err != nil {
						return err
					}
					for _, kubePod := range pods {
						nft.SetPod(cache.MetaObjectToName(&kubePod), &kubePod)
					}
					policies, err := t.Kubernetes.GetNetworkPoliciesInNamespace(kubeNamespace.Name)
					if err != nil {
						return err
					}
					for _, policy := range policies {
						nft.SetNetworkPolicy(cache.MetaObjectToName(&policy), &policy)
					}
				}
				for _, kubeNamespace := range namespaces.Items {
					nft.SetNamespace(kubeNamespace.Name, &kubeNamespace)
				}
				return nil
			},
		},
	}
}

func TestCyclonus(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}

	if err := os.WriteFile("/proc/sys/net/ipv4/ip_forward", []byte("1\n"), 0644); err != nil {
		t.Fatalf("Failed to enable IPv4 forwarding: %v", err)
	}
	if err := os.WriteFile("/proc/sys/net/ipv6/conf/all/forwarding", []byte("1\n"), 0644); err != nil {
		t.Fatalf("Failed to enable IPv6 forwarding: %v", err)
	}
	// By default, linux rate limits ICMP replies, which slows down our tests.
	if err := os.WriteFile("/proc/sys/net/ipv4/icmp_ratemask", []byte("0\n"), 0644); err != nil {
		t.Fatalf("Failed to disable ICMPv4 rate limiting: %v", err)
	}
	if err := os.WriteFile("/proc/sys/net/ipv6/icmp/ratemask", []byte("\n"), 0644); err != nil {
		t.Fatalf("Failed to disable ICMPv6 rate limiting: %v", err)
	}
	// Disable IPv6 duplicate address detection, which delays IPv6 connectivity
	// becoming available.
	if err := os.WriteFile("/proc/sys/net/ipv6/conf/default/accept_dad", []byte("0\n"), 0644); err != nil {
		t.Fatalf("Failed to disable IPv6 DAD: %v", err)
	}

	kubernetes := &fakeKubernetes{
		MockKubernetes: kube.NewMockKubernetes(1.0),
		fakePods:       make(map[string]*fakePod),
		nextPodIP:      2,
	}

	allowDNS := false // This creates policies which allow port 53, we don't need this.
	serverNamespaces := []string{"x", "y", "z"}
	serverPods := []string{"a", "b", "c"}
	externalIPs := []string{}
	podCreationTimeoutSeconds := 1
	batchJobs := false
	imageRegistry := "registry.k8s.io"

	resources, err := probe.NewDefaultResources(kubernetes, serverNamespaces, serverPods, serverPorts, serverProtocols, externalIPs, podCreationTimeoutSeconds, batchJobs, imageRegistry)
	if err != nil {
		t.Fatal(err)
	}

	interpreterConfig := &connectivity.InterpreterConfig{
		ResetClusterBeforeTestCase:       false,
		KubeProbeRetries:                 1,
		PerturbationWaitSeconds:          0,
		VerifyClusterStateBeforeTestCase: true,
		BatchJobs:                        batchJobs,
		IgnoreLoopback:                   ignoreLoopback,
		JobTimeoutSeconds:                1,
		FailFast:                         false,
	}
	interpreter := connectivity.NewInterpreter(kubernetes, resources, interpreterConfig)
	printer := &connectivity.Printer{
		Noisy:            false,
		IgnoreLoopback:   ignoreLoopback,
		JunitResultsFile: "",
	}

	zcPod, err := resources.GetPod("z", "c")
	if err != nil {
		t.Fatal(err)
	}
	testCaseGenerator := generator.NewTestCaseGenerator(allowDNS, zcPod.IP, serverNamespaces, nil, nil)
	testCases := slices.Concat(testCaseGenerator.GenerateTestCases(), extraTestCases())

	for _, testCase := range testCases {
		for _, step := range testCase.Steps {
			step.Probe.Mode = generator.ProbeModeServiceName
		}
		t.Run(testCase.Description, func(t *testing.T) {
			recorder := &testRecorder{t: t}
			nft, err := nftctrl.New(recorder, podIfaceGroup)
			if err != nil {
				t.Fatalf("Failed to create nftctrl: %v", err)
			}
			defer nft.Close()
			kubernetes.nft = nft

			if err := kubernetes.initializeNft(); err != nil {
				t.Fatalf("nftctrl initialization failed: %v", err)
			}

			result := interpreter.ExecuteTestCase(testCase)
			if result.Err != nil {
				t.Fatal(result.Err)
			}
			if !result.Passed(ignoreLoopback) {
				printer.PrintTestCaseResult(result)

				cmd := exec.Command("nft", "list", "ruleset")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()

				t.Error("connectivity test failed")
			}

			kubernetes.nft = nil
			testCaseState := &connectivity.TestCaseState{
				Kubernetes: kubernetes,
				Resources:  resources,
			}
			err = testCaseState.ResetClusterState()
			if err != nil {
				t.Errorf("failed to reset cluster: %v", err)
			}

			// Flush the conntrack table. Otherwise, UDP connectivity tests can
			// spuriously succeed when they should be blocked, because they match an
			// entry in the conntrack table from a previous test.
			err = netlink.ConntrackTableFlush(netlink.ConntrackTable)
			if err != nil {
				t.Errorf("failed to flush conntrack table: %v", err)
			}
		})
	}

	initTestCases := initializationTestCases()
	for _, testCase := range initTestCases {
		t.Run(testCase.description, func(t *testing.T) {
			recorder := &testRecorder{t: t}
			nft, err := nftctrl.New(recorder, podIfaceGroup)
			if err != nil {
				t.Fatalf("Failed to create nftctrl: %v", err)
			}
			defer nft.Close()

			testCaseState := &connectivity.TestCaseState{
				Kubernetes: kubernetes,
				Resources:  resources,
			}

			if err := testCase.init(testCaseState, nft); err != nil {
				t.Fatalf("initialization failed: %v", err)
			}

			if err := nft.Flush(); err != nil {
				t.Fatalf("flush failed: %v", err)
			}

			parsedPolicy := matcher.BuildNetworkPolicies(true, testCaseState.Policies)
			jobBuilder := &probe.JobBuilder{TimeoutSeconds: 1}
			simRunner := probe.NewSimulatedRunner(parsedPolicy, jobBuilder)
			probeConfig := generator.NewAllAvailable(generator.ProbeModeServiceName)
			stepResult := connectivity.NewStepResult(
				simRunner.RunProbeForConfig(probeConfig, testCaseState.Resources),
				parsedPolicy,
				slices.Clone(testCaseState.Policies))
			kubeRunner := probe.NewKubeRunner(kubernetes, 15, jobBuilder)
			for i := 0; i <= interpreterConfig.KubeProbeRetries; i++ {
				stepResult.AddKubeProbe(kubeRunner.RunProbeForConfig(probeConfig, testCaseState.Resources))
				if stepResult.Passed(ignoreLoopback) {
					break
				}
			}

			if !stepResult.Passed(ignoreLoopback) {
				printer.PrintStep(0, &generator.TestStep{Probe: probeConfig}, stepResult)

				cmd := exec.Command("nft", "list", "ruleset")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()

				t.Error("connectivity test failed")
			}

			err = testCaseState.ResetClusterState()
			if err != nil {
				t.Errorf("failed to reset cluster: %v", err)
			}

			// Flush the conntrack table. Otherwise, UDP connectivity tests can
			// spuriously succeed when they should be blocked, because they match an
			// entry in the conntrack table from a previous test.
			err = netlink.ConntrackTableFlush(netlink.ConntrackTable)
			if err != nil {
				t.Errorf("failed to flush conntrack table: %v", err)
			}
		})
	}
}
