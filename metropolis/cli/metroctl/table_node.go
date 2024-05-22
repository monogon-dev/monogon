package main

import (
	"fmt"
	"sort"
	"strings"

	"source.monogon.dev/go/clitable"
	"source.monogon.dev/metropolis/node/core/identity"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
	"source.monogon.dev/version"
)

func nodeEntry(n *apb.Node) clitable.Entry {
	res := clitable.Entry{}

	res.Add("node id", identity.NodeID(n.Pubkey))
	state := n.State.String()
	state = strings.ReplaceAll(state, "NODE_STATE_", "")
	res.Add("state", state)
	address := "unknown"
	if n.Status != nil && n.Status.ExternalAddress != "" {
		address = n.Status.ExternalAddress
	}
	res.Add("address", address)
	res.Add("health", n.Health.String())

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
	res.Add("roles", strings.Join(roles, ","))

	tpm := "unk"
	switch n.TpmUsage {
	case cpb.NodeTPMUsage_NODE_TPM_PRESENT_AND_USED:
		tpm = "yes"
	case cpb.NodeTPMUsage_NODE_TPM_PRESENT_BUT_UNUSED:
		tpm = "unused"
	case cpb.NodeTPMUsage_NODE_TPM_NOT_PRESENT:
		tpm = "no"
	}
	res.Add("tpm", tpm)

	if n.Status != nil && n.Status.Version != nil {
		res.Add("version", version.Semver(n.Status.Version))
	}

	tshs := n.TimeSinceHeartbeat.GetSeconds()
	res.Add("heartbeat", fmt.Sprintf("%ds", tshs))

	return res
}
