package clusternet

import (
	"testing"

	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

func TestNodeUpdate(t *testing.T) {
	var n node

	for i, te := range []struct {
		p    *apb.Node
		want bool
	}{
		// Case 0: empty incoming node, no change.
		{
			p:    &apb.Node{Status: &cpb.NodeStatus{}, Clusternet: &cpb.NodeClusterNetworking{}},
			want: false,
		},
		// Case 1: wireguard key set.
		{
			p:    &apb.Node{Status: &cpb.NodeStatus{}, Clusternet: &cpb.NodeClusterNetworking{WireguardPubkey: "fake"}},
			want: true,
		},
		// Case 2: wireguard key updated.
		{
			p:    &apb.Node{Status: &cpb.NodeStatus{}, Clusternet: &cpb.NodeClusterNetworking{WireguardPubkey: "fake2"}},
			want: true,
		},
		// Case 3: external address added.
		{
			p:    &apb.Node{Status: &cpb.NodeStatus{ExternalAddress: "1.2.3.4"}, Clusternet: &cpb.NodeClusterNetworking{WireguardPubkey: "fake2"}},
			want: true,
		},
		// Case 4: external address changed.
		{
			p:    &apb.Node{Status: &cpb.NodeStatus{ExternalAddress: "1.2.3.5"}, Clusternet: &cpb.NodeClusterNetworking{WireguardPubkey: "fake2"}},
			want: true,
		},
		// Case 5: prefixes added
		{
			p: &apb.Node{
				Status: &cpb.NodeStatus{ExternalAddress: "1.2.3.5"},
				Clusternet: &cpb.NodeClusterNetworking{
					WireguardPubkey: "fake2",
					Prefixes: []*cpb.NodeClusterNetworking_Prefix{
						{Cidr: "10.0.2.0/24"},
						{Cidr: "10.0.1.0/24"},
					},
				},
			},
			want: true,
		},
		// Case 6: prefixes changed
		{
			p: &apb.Node{
				Status: &cpb.NodeStatus{ExternalAddress: "1.2.3.5"},
				Clusternet: &cpb.NodeClusterNetworking{
					WireguardPubkey: "fake2",
					Prefixes: []*cpb.NodeClusterNetworking_Prefix{
						{Cidr: "10.0.3.0/24"},
						{Cidr: "10.0.1.0/24"},
					},
				},
			},
			want: true,
		},
		// Case 6: prefixes reordered (no change expected)
		{
			p: &apb.Node{
				Status: &cpb.NodeStatus{ExternalAddress: "1.2.3.5"},
				Clusternet: &cpb.NodeClusterNetworking{
					WireguardPubkey: "fake2",
					Prefixes: []*cpb.NodeClusterNetworking_Prefix{
						{Cidr: "10.0.3.0/24"},
						{Cidr: "10.0.1.0/24"},
					},
				},
			},
			want: false,
		},
		// Case 6: prefixes removed
		{
			p: &apb.Node{
				Status: &cpb.NodeStatus{ExternalAddress: "1.2.3.5"},
				Clusternet: &cpb.NodeClusterNetworking{
					WireguardPubkey: "fake2",
					Prefixes:        []*cpb.NodeClusterNetworking_Prefix{},
				},
			},
			want: true,
		},
	} {
		got := n.update(te.p)
		if te.want && !got {
			t.Fatalf("Case %d: expected change, got no change", i)
		}
		if !te.want && got {
			t.Fatalf("Case %d: expected no change, got change", i)
		}
	}
}
