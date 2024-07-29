package kubernetes

import (
	"context"
	"fmt"
	"net/netip"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/miekg/dns"
	api "k8s.io/api/core/v1"
	discovery "k8s.io/api/discovery/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/utils/ptr"

	netDNS "source.monogon.dev/osbase/net/dns"
)

const testdataClusterDomain = "cluster.local"

var testdataIPRanges = []string{
	// service IP
	"10.0.0.1/16",
	"1234:abcd::/64",
	// pod IP
	"172.32.0.0/11",
	"170::/14",
}

var testdataNamespaces = []string{"testns"}

var testdataServices = []*api.Service{
	{
		ObjectMeta: meta.ObjectMeta{
			Name:      "svc-clusterip",
			Namespace: "testns",
		},
		Spec: api.ServiceSpec{
			Type:       api.ServiceTypeClusterIP,
			ClusterIPs: []string{"10.0.0.10"},
			Ports: []api.ServicePort{
				{Name: "http", Protocol: api.ProtocolTCP, Port: 80, TargetPort: intstr.FromInt32(82)},
			},
		},
	},
	{
		ObjectMeta: meta.ObjectMeta{
			Name:      "svc-dualstack",
			Namespace: "testns",
		},
		Spec: api.ServiceSpec{
			Type:       api.ServiceTypeClusterIP,
			ClusterIPs: []string{"10.0.0.11", "1234:abcd::11"},
			Ports: []api.ServicePort{
				{Name: "http", Protocol: api.ProtocolTCP, Port: 80},
				{Name: "dns", Protocol: api.ProtocolUDP, Port: 53},
			},
		},
	},
	{
		ObjectMeta: meta.ObjectMeta{
			Name:      "svc-headless",
			Namespace: "testns",
		},
		Spec: api.ServiceSpec{
			Type:       api.ServiceTypeClusterIP,
			ClusterIP:  api.ClusterIPNone,
			ClusterIPs: []string{api.ClusterIPNone},
			Ports: []api.ServicePort{
				{Name: "http", Protocol: api.ProtocolTCP, Port: 80, TargetPort: intstr.FromString("http")},
			},
		},
	},
	{
		ObjectMeta: meta.ObjectMeta{
			Name:      "svc-headless-notready",
			Namespace: "testns",
		},
		Spec: api.ServiceSpec{
			Type:       api.ServiceTypeClusterIP,
			ClusterIP:  api.ClusterIPNone,
			ClusterIPs: []string{api.ClusterIPNone},
			Ports: []api.ServicePort{
				{Name: "http", Protocol: api.ProtocolTCP, Port: 80, TargetPort: intstr.FromString("http")},
			},
		},
	},
	{
		ObjectMeta: meta.ObjectMeta{
			Name:      "svc-external",
			Namespace: "testns",
		},
		Spec: api.ServiceSpec{
			Type:         api.ServiceTypeExternalName,
			ExternalName: "external.example.com",
		},
	},
}

var testdataEndpointSlices = []*discovery.EndpointSlice{
	{
		ObjectMeta: meta.ObjectMeta{
			Name:      "svc-clusterip-slice",
			Namespace: "testns",
			Labels:    map[string]string{discovery.LabelServiceName: "svc-clusterip"},
		},
		Endpoints: []discovery.Endpoint{
			{Addresses: []string{"172.45.0.1"}},
		},
	},
	{
		ObjectMeta: meta.ObjectMeta{
			Name:      "svc-headless-slice1",
			Namespace: "testns",
			Labels: map[string]string{
				discovery.LabelServiceName: "svc-headless",
				api.IsHeadlessService:      "",
			},
		},
		Endpoints: []discovery.Endpoint{
			{
				Addresses: []string{"172.45.0.2"},
			},
			{
				Addresses: []string{"172.45.0.2"},
			},
			{
				Hostname:   ptr.To("pod3"),
				Addresses:  []string{"172.45.0.3"},
				Conditions: discovery.EndpointConditions{Ready: ptr.To(true)},
			},
			{
				Addresses:  []string{"172.45.0.4"},
				Conditions: discovery.EndpointConditions{Ready: ptr.To(false)},
			},
			{
				Hostname:  ptr.To("pod5"),
				Addresses: []string{"172.45.0.5", "172.45.0.2"},
			},
		},
		Ports: []discovery.EndpointPort{
			{Name: ptr.To("http"), Port: ptr.To(int32(8000)), Protocol: ptr.To(api.ProtocolTCP)},
		},
	},
	{
		ObjectMeta: meta.ObjectMeta{
			Name:      "svc-headless-slice2",
			Namespace: "testns",
			Labels: map[string]string{
				discovery.LabelServiceName: "svc-headless",
				api.IsHeadlessService:      "",
			},
		},
		Endpoints: []discovery.Endpoint{
			{
				Hostname:   ptr.To("pod3"),
				Addresses:  []string{"172::3"},
				Conditions: discovery.EndpointConditions{Ready: ptr.To(false)},
			},
			{
				Hostname:  ptr.To("pod5"),
				Addresses: []string{"172::5"},
			},
			{
				Addresses: []string{"172::7"},
			},
		},
		Ports: []discovery.EndpointPort{
			{Name: ptr.To("http"), Port: ptr.To(int32(8001)), Protocol: ptr.To(api.ProtocolTCP)},
		},
	},
	{
		ObjectMeta: meta.ObjectMeta{
			Name:      "svc-headless-notready-slice1",
			Namespace: "testns",
			Labels: map[string]string{
				discovery.LabelServiceName: "svc-headless-notready",
				api.IsHeadlessService:      "",
			},
		},
		Endpoints: []discovery.Endpoint{
			{
				Addresses:  []string{"172.45.0.20"},
				Conditions: discovery.EndpointConditions{Ready: ptr.To(false)},
			},
			{
				Hostname:   ptr.To("pod21"),
				Addresses:  []string{"172.45.0.21"},
				Conditions: discovery.EndpointConditions{Ready: ptr.To(false)},
			},
		},
		Ports: []discovery.EndpointPort{
			{Name: ptr.To("http"), Port: ptr.To(int32(8000)), Protocol: ptr.To(api.ProtocolTCP)},
		},
	},
}

// handlerTestcase contains a query name, and the expected records
// under that name given the above test data.
type handlerTestcase struct {
	// Query name
	qname string

	// Expected reply

	rcode         int
	answer, extra []string
	notHandled    bool
	// zone is the zone that is expected in the NS SOA if the answer is empty.
	// If zone is empty, defaults to "cluster.local."
	zone string
}

// nameErrorIfSynced means name error if synced, else server failure.
const nameErrorIfSynced = -1

var handlerTestcases = []handlerTestcase{
	// cluster domain root
	{
		qname: "cluster.local.",
		answer: []string{
			"cluster.local.	5	IN	SOA	ns.dns.cluster.local. nobody.invalid. 12345 7200 1800 86400 5",
			"cluster.local.	5	IN	NS	ns.dns.cluster.local.",
		},
	},
	{
		qname: "example.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	// dns-version
	{
		qname: "dns-version.cluster.local.",
		answer: []string{
			`dns-version.cluster.local.	5	IN	TXT	"1.1.0"`,
		},
	},
	{
		qname: "example.dns-version.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	// ns.dns
	{
		qname: "dns.cluster.local.",
	},
	{
		qname: "ns.dns.cluster.local.",
	},
	{
		qname: "example.dns.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	// svc
	{
		qname: "svc.cluster.local.",
	},
	// namespace
	{
		qname: "testns.svc.cluster.local.",
	},
	{
		qname: "inexistent-ns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	// cluster IP service
	{
		qname: "svc-clusterip.testns.svc.cluster.local.",
		answer: []string{
			"svc-clusterip.testns.svc.cluster.local.	5	IN	A	10.0.0.10",
		},
	},
	{
		qname: "_http._tcp.svc-clusterip.testns.svc.cluster.local.",
		answer: []string{
			"_http._tcp.svc-clusterip.testns.svc.cluster.local.	5	IN	SRV	0 0 80 svc-clusterip.testns.svc.cluster.local.",
		},
		extra: []string{
			"svc-clusterip.testns.svc.cluster.local.	5	IN	A	10.0.0.10",
		},
	},
	{
		qname: "_udp.svc-clusterip.testns.svc.cluster.local.",
	},
	{
		qname: "_http._udp.svc-clusterip.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "http._tcp.svc-clusterip.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "example._http._tcp.svc-clusterip.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "10.0.0.10.in-addr.arpa.",
		answer: []string{
			"10.0.0.10.in-addr.arpa.	5	IN	PTR	svc-clusterip.testns.svc.cluster.local.",
		},
		zone: "0.10.in-addr.arpa.",
	},
	{
		qname: "172-45-0-1.svc-clusterip.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "1.0.45.172.in-addr.arpa.",
		rcode: nameErrorIfSynced,
		zone:  "45.172.in-addr.arpa.",
	},
	// dual stack cluster IP service
	{
		qname: "svc-dualstack.testns.svc.cluster.local.",
		answer: []string{
			"svc-dualstack.testns.svc.cluster.local.	5	IN	A	10.0.0.11",
			"svc-dualstack.testns.svc.cluster.local.	5	IN	AAAA	1234:abcd::11",
		},
	},
	{
		qname: "_http._tcp.svc-dualstack.testns.svc.cluster.local.",
		answer: []string{
			"_http._tcp.svc-dualstack.testns.svc.cluster.local.	5	IN	SRV	0 0 80 svc-dualstack.testns.svc.cluster.local.",
		},
		extra: []string{
			"svc-dualstack.testns.svc.cluster.local.	5	IN	A	10.0.0.11",
			"svc-dualstack.testns.svc.cluster.local.	5	IN	AAAA	1234:abcd::11",
		},
	},
	{
		qname: "11.0.0.10.in-addr.arpa.",
		answer: []string{
			"11.0.0.10.in-addr.arpa.	5	IN	PTR	svc-dualstack.testns.svc.cluster.local.",
		},
		zone: "0.10.in-addr.arpa.",
	},
	{
		qname: "1.1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.d.c.b.a.4.3.2.1.ip6.arpa.",
		answer: []string{
			"1.1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.d.c.b.a.4.3.2.1.ip6.arpa.	5	IN	PTR	svc-dualstack.testns.svc.cluster.local.",
		},
		zone: "0.0.0.0.0.0.0.0.d.c.b.a.4.3.2.1.ip6.arpa.",
	},
	// headless service
	{
		qname: "svc-headless.testns.svc.cluster.local.",
		answer: []string{
			"svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.2",
			"svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.3",
			"svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.5",
			"svc-headless.testns.svc.cluster.local.	5	IN	AAAA	172::5",
			"svc-headless.testns.svc.cluster.local.	5	IN	AAAA	172::7",
		},
	},
	{
		qname: "_http._tcp.svc-headless.testns.svc.cluster.local.",
		answer: []string{
			"_http._tcp.svc-headless.testns.svc.cluster.local.	5	IN	SRV	0 0 8000 172-45-0-2.svc-headless.testns.svc.cluster.local.",
			"_http._tcp.svc-headless.testns.svc.cluster.local.	5	IN	SRV	0 0 8000 pod3.svc-headless.testns.svc.cluster.local.",
			"_http._tcp.svc-headless.testns.svc.cluster.local.	5	IN	SRV	0 0 8000 pod5.svc-headless.testns.svc.cluster.local.",
			"_http._tcp.svc-headless.testns.svc.cluster.local.	5	IN	SRV	0 0 8001 pod5.svc-headless.testns.svc.cluster.local.",
			"_http._tcp.svc-headless.testns.svc.cluster.local.	5	IN	SRV	0 0 8001 172--7.svc-headless.testns.svc.cluster.local.",
		},
		extra: []string{
			"172-45-0-2.svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.2",
			"pod3.svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.3",
			"pod5.svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.5",
			"pod5.svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.2",
			"pod5.svc-headless.testns.svc.cluster.local.	5	IN	AAAA	172::5",
			"172--7.svc-headless.testns.svc.cluster.local.	5	IN	AAAA	172::7",
		},
	},
	{
		qname: "_udp.svc-headless.testns.svc.cluster.local.",
	},
	{
		qname: "_http._udp.svc-headless.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "http._tcp.svc-headless.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "_._udp.svc-headless.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "example._http._tcp.svc-headless.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "172-45-0-2.svc-headless.testns.svc.cluster.local.",
		answer: []string{
			"172-45-0-2.svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.2",
		},
	},
	{
		qname: "pod5.svc-headless.testns.svc.cluster.local.",
		answer: []string{
			"pod5.svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.5",
			"pod5.svc-headless.testns.svc.cluster.local.	5	IN	A	172.45.0.2",
			"pod5.svc-headless.testns.svc.cluster.local.	5	IN	AAAA	172::5",
		},
	},
	{
		qname: "example.pod5.svc-headless.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "172-45-0-5.svc-headless.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "2.0.45.172.in-addr.arpa.",
		answer: []string{
			"2.0.45.172.in-addr.arpa.	5	IN	PTR	172-45-0-2.svc-headless.testns.svc.cluster.local.",
			"2.0.45.172.in-addr.arpa.	5	IN	PTR	pod5.svc-headless.testns.svc.cluster.local.",
		},
		zone: "45.172.in-addr.arpa.",
	},
	{
		qname: "5.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.2.7.1.0.ip6.arpa.",
		answer: []string{
			"5.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.2.7.1.0.ip6.arpa.	5	IN	PTR	pod5.svc-headless.testns.svc.cluster.local.",
		},
		zone: "2.7.1.0.ip6.arpa.",
	},
	// not ready headless service
	{
		qname: "svc-headless-notready.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "_tcp.svc-headless-notready.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "_http._tcp.svc-headless-notready.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "pod21.svc-headless-notready.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "21.0.45.172.in-addr.arpa.",
		rcode: nameErrorIfSynced,
		zone:  "45.172.in-addr.arpa.",
	},
	// external service
	{
		qname: "svc-external.testns.svc.cluster.local.",
		answer: []string{
			"svc-external.testns.svc.cluster.local.	5	IN	CNAME	external.example.com.",
		},
	},
	{
		qname: "_tcp.svc-external.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "_http._tcp.svc-external.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "pod.svc-external.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	// service does not exist
	{
		qname: "inexistent-svc.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "_tcp.inexistent-svc.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "_http._tcp.inexistent-svc.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "example._tcp.inexistent-svc.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "example._http._tcp.inexistent-svc.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "pod.inexistent-svc.testns.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "example.pod.inexistent-svc.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	// names which do not exist but will get queried because of ndots=5
	{
		qname: "www.example.com.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	{
		qname: "www.example.com.svc.cluster.local.",
		rcode: nameErrorIfSynced,
	},
	{
		qname: "www.example.com.testns.svc.cluster.local.",
		rcode: dns.RcodeNameError,
	},
	// names which are not handled
	{
		qname:      "www.example.com.",
		notHandled: true,
	},
	{
		qname:      "12.0.31.172.in-addr.arpa.",
		notHandled: true,
	},
	{
		qname:      "5.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.4.7.1.0.ip6.arpa.",
		notHandled: true,
	},
	{
		qname:      "10.in-addr.arpa.",
		notHandled: true,
	},
	{
		qname:      "7.1.0.ip6.arpa.",
		notHandled: true,
	},
	// reverse lookup zone
	{
		qname: "45.172.in-addr.arpa.",
		answer: []string{
			"45.172.in-addr.arpa.	5	IN	SOA	ns.dns.cluster.local. nobody.invalid. 12345 7200 1800 86400 5",
			"45.172.in-addr.arpa.	5	IN	NS	ns.dns.cluster.local.",
		},
		zone: "45.172.in-addr.arpa.",
	},
	{
		qname: "255.45.172.in-addr.arpa.",
		zone:  "45.172.in-addr.arpa.",
	},
	{
		qname: "02.0.45.172.in-addr.arpa.",
		rcode: dns.RcodeNameError,
		zone:  "45.172.in-addr.arpa.",
	},
	{
		qname: "1.2.0.45.172.in-addr.arpa.",
		rcode: dns.RcodeNameError,
		zone:  "45.172.in-addr.arpa.",
	},
	{
		qname: "2.7.1.0.ip6.arpa.",
		answer: []string{
			"2.7.1.0.ip6.arpa.	5	IN	SOA	ns.dns.cluster.local. nobody.invalid. 12345 7200 1800 86400 5",
			"2.7.1.0.ip6.arpa.	5	IN	NS	ns.dns.cluster.local.",
		},
		zone: "2.7.1.0.ip6.arpa.",
	},
	{
		qname: "a.2.7.1.0.ip6.arpa.",
		zone:  "2.7.1.0.ip6.arpa.",
	},
	{
		qname: "x.a.2.7.1.0.ip6.arpa.",
		rcode: dns.RcodeNameError,
		zone:  "2.7.1.0.ip6.arpa.",
	},
	// mixed case
	{
		qname: "SvC-cLUSteRIp.TesTNS.sVC.ClUSTer.locAL.",
		answer: []string{
			"SvC-cLUSteRIp.TesTNS.sVC.ClUSTer.locAL.	5	IN	A	10.0.0.10",
		},
		zone: "ClUSTer.locAL.",
	},
	{
		qname: "_hTTp._tCp.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.",
		answer: []string{
			"_hTTp._tCp.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	SRV	0 0 8000 172-45-0-2.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.",
			"_hTTp._tCp.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	SRV	0 0 8000 pod3.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.",
			"_hTTp._tCp.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	SRV	0 0 8000 pod5.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.",
			"_hTTp._tCp.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	SRV	0 0 8001 pod5.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.",
			"_hTTp._tCp.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	SRV	0 0 8001 172--7.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.",
		},
		extra: []string{
			"172-45-0-2.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	A	172.45.0.2",
			"pod3.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	A	172.45.0.3",
			"pod5.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	A	172.45.0.5",
			"pod5.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	A	172.45.0.2",
			"pod5.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	AAAA	172::5",
			"172--7.SVc-hEADlEsS.teSTNs.SVC.ClUSTer.locAL.	5	IN	AAAA	172::7",
		},
		zone: "ClUSTer.locAL.",
	},
	{
		qname: "1.1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.d.C.b.a.4.3.2.1.iP6.ARpa.",
		answer: []string{
			"1.1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.d.C.b.a.4.3.2.1.iP6.ARpa.	5	IN	PTR	svc-dualstack.testns.svc.cluster.local.",
		},
		zone: "0.0.0.0.0.0.0.0.d.C.b.a.4.3.2.1.iP6.ARpa.",
	},
}

// TestHandler constructs a fake Kubernetes clientset containing the above
// testdata, and then evaluates each test case in handlerTestcases.
func TestHandler(t *testing.T) {
	ctx := context.Background()
	client := fake.NewSimpleClientset()

	// Add resources
	for _, name := range testdataNamespaces {
		namespace := &api.Namespace{
			ObjectMeta: meta.ObjectMeta{Name: name},
		}
		_, err := client.CoreV1().Namespaces().Create(ctx, namespace, meta.CreateOptions{})
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, service := range testdataServices {
		_, err := client.CoreV1().Services(service.Namespace).Create(ctx, service, meta.CreateOptions{})
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, slice := range testdataEndpointSlices {
		_, err := client.DiscoveryV1().EndpointSlices(slice.Namespace).Create(ctx, slice, meta.CreateOptions{})
		if err != nil {
			t.Fatal(err)
		}
	}

	// Create handler
	var ipRanges []netip.Prefix
	for _, ipRange := range testdataIPRanges {
		ipRanges = append(ipRanges, netip.MustParsePrefix(ipRange))
	}
	handler := New(testdataClusterDomain, ipRanges)
	handler.ClientSet = client

	wrapper := &dnsControllerWrapper{dnsController: newdnsController(ctx, handler.ClientSet)}
	handler.apiConn = wrapper

	stopCh := make(chan struct{})
	defer close(stopCh)
	handler.apiConn.Start(stopCh)
	for !wrapper.dnsController.HasSynced() {
		time.Sleep(time.Millisecond)
	}

	for _, hasSynced := range []bool{true, false} {
		wrapper.hasSynced = hasSynced
		for _, testcase := range handlerTestcases {
			if testcase.zone == "" {
				testcase.zone = "cluster.local."
			}
			if testcase.rcode == nameErrorIfSynced {
				if hasSynced {
					testcase.rcode = dns.RcodeNameError
				} else {
					testcase.rcode = dns.RcodeServerFailure
				}
			}

			qtypes := []uint16{
				dns.TypeANY, dns.TypeA, dns.TypeAAAA, dns.TypeSRV, dns.TypeTXT,
				dns.TypeNS, dns.TypeSOA, dns.TypePTR, dns.TypeMX, dns.TypeCNAME,
			}
			for _, qtype := range qtypes {
				doHandlerTestcase(t, handler, testcase, qtype)
			}
		}
	}

	wrapper.hasSynced = false
	testNotSyncedOpt(t, handler)
}

func doHandlerTestcase(t *testing.T, handler *Kubernetes, testcase handlerTestcase, qtype uint16) {
	// Create request
	req := netDNS.CreateTestRequest(testcase.qname, qtype, "udp")
	req.Reply.RecursionDesired = false
	req.Qopt = nil
	req.Ropt = nil

	handler.HandleDNS(req)

	caseName := fmt.Sprintf("Query %s %s", testcase.qname, dns.TypeToString[qtype])
	if !handler.apiConn.HasSynced() {
		caseName += " not_synced"
	}

	if req.Handled != !testcase.notHandled {
		t.Errorf("%s: Expected handled %v, got %v", caseName,
			!testcase.notHandled, req.Handled,
		)
		return
	}
	if !req.Handled {
		return
	}

	if req.Reply.Rcode != testcase.rcode {
		t.Errorf("%s: Expected rcode %s, got %s", caseName,
			dns.RcodeToString[testcase.rcode], dns.RcodeToString[req.Reply.Rcode],
		)
		return
	}

	// Create expected answer
	var answer []string
	for _, rr := range testcase.answer {
		rrParsed, err := dns.NewRR(rr)
		if err != nil {
			t.Fatalf("Failed to parse DNS RR %q: %v", rr, err)
		}
		if qtype == dns.TypeANY || qtype == rrParsed.Header().Rrtype || rrParsed.Header().Rrtype == dns.TypeCNAME {
			answer = append(answer, rr)
		}
	}
	var extra []string
	var ns []string
	if len(answer) != 0 {
		extra = testcase.extra
	} else {
		ns = []string{
			testcase.zone + "	5	IN	SOA	ns.dns.cluster.local. nobody.invalid. 12345 7200 1800 86400 5",
		}
	}

	checkReplySection(t, caseName, "answer", answer, req.Reply.Answer)
	checkReplySection(t, caseName, "ns", ns, req.Reply.Ns)
	checkReplySection(t, caseName, "extra", extra, req.Reply.Extra)
}

func checkReplySection(t *testing.T, caseName string, sectionName string, expected []string, got []dns.RR) {
	slices.Sort(expected)
	var gotStr []string
	for _, rr := range got {
		gotStr = append(gotStr, rr.String())
	}
	slices.Sort(gotStr)
	if !slices.Equal(expected, gotStr) {
		t.Errorf("%s: Expected %s:\n%s\nGot:\n%v", caseName, sectionName,
			strings.Join(expected, "\n"), strings.Join(gotStr, "\n"))
	}
}

// testNotSyncedOpt tests that we get the Not Ready extended error
// when not synced and an OPT is present and no result was found.
func testNotSyncedOpt(t *testing.T, handler *Kubernetes) {
	req := netDNS.CreateTestRequest("inexistent-ns.svc.cluster.local.", dns.TypeA, "udp")

	handler.HandleDNS(req)
	extra := []string{
		"\n" +
			";; OPT PSEUDOSECTION:\n" +
			"; EDNS: version 0; flags:; udp: 1232\n" +
			"; EDE: 14 (Not Ready): (Kubernetes objects not yet synced)",
	}
	checkReplySection(t, "testNotSyncedOpt", "extra", extra, req.Reply.Extra)
}

type dnsControllerWrapper struct {
	dnsController
	hasSynced bool
}

func (dns *dnsControllerWrapper) HasSynced() bool {
	return dns.hasSynced
}

func (dns *dnsControllerWrapper) Modified() int64 {
	return 12345
}
