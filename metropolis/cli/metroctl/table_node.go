package main

import (
	"fmt"
	"sort"
	"strings"

	"source.monogon.dev/metropolis/node/core/identity"
	apb "source.monogon.dev/metropolis/proto/api"
)

func nodeEntry(n *apb.Node) entry {
	res := entry{}

	res.add("node id", identity.NodeID(n.Pubkey))
	state := n.State.String()
	state = strings.ReplaceAll(state, "NODE_STATE_", "")
	res.add("state", state)
	res.add("address", n.Status.ExternalAddress)
	res.add("health", n.Health.String())

	var roles []string
	if n.Roles.ConsensusMember != nil {
		roles = append(roles, "ConsensusMember")
	}
	if n.Roles.KubernetesController != nil {
		roles = append(roles, "KubernetesController")
	}
	if n.Roles.KubernetesWorker != nil {
		roles = append(roles, "KubernetesWorker")
	}
	sort.Strings(roles)
	res.add("roles", strings.Join(roles, ","))

	tshs := n.TimeSinceHeartbeat.GetSeconds()
	res.add("heartbeat", fmt.Sprintf("%ds", tshs))

	return res
}
