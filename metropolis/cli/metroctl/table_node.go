package main

import (
	"fmt"
	"sort"
	"strings"

	"source.monogon.dev/metropolis/node/core/identity"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
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

	tpm := "unk"
	switch n.TpmUsage {
	case cpb.NodeTPMUsage_NODE_TPM_PRESENT_AND_USED:
		tpm = "yes"
	case cpb.NodeTPMUsage_NODE_TPM_PRESENT_BUT_UNUSED:
		tpm = "unused"
	case cpb.NodeTPMUsage_NODE_TPM_NOT_PRESENT:
		tpm = "no"
	}
	res.add("tpm", tpm)

	tshs := n.TimeSinceHeartbeat.GetSeconds()
	res.add("heartbeat", fmt.Sprintf("%ds", tshs))

	return res
}
