package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"source.monogon.dev/metropolis/node/core/identity"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

type encoder struct {
	out io.WriteCloser
}

func (e *encoder) writeNodeID(n *apb.Node) error {
	id := identity.NodeID(n.Pubkey)
	_, err := fmt.Fprintf(e.out, "%s\n", id)
	return err
}

func (e *encoder) writeNode(n *apb.Node) error {
	id := identity.NodeID(n.Pubkey)
	if _, err := fmt.Fprintf(e.out, "%s", id); err != nil {
		return err
	}

	state := cpb.NodeState_name[int32(n.State)]
	if _, err := fmt.Fprintf(e.out, "\t%s", state); err != nil {
		return err
	}

	addr := n.Status.ExternalAddress
	if _, err := fmt.Fprintf(e.out, "\t%s", addr); err != nil {
		return err
	}

	health := apb.Node_Health_name[int32(n.Health)]
	if _, err := fmt.Fprintf(e.out, "\t%s", health); err != nil {
		return err
	}

	var roles string
	if n.Roles.KubernetesWorker != nil {
		roles += "KubernetesWorker"
	}
	if n.Roles.ConsensusMember != nil {
		roles += ",ConsensusMember"
	}
	if _, err := fmt.Fprintf(e.out, "\t%s", roles); err != nil {
		return err
	}

	tshs := n.TimeSinceHeartbeat.GetSeconds()
	if _, err := fmt.Fprintf(e.out, "\t%ds\n", tshs); err != nil {
		return err
	}
	return nil
}

func (e *encoder) close() error {
	if e.out != os.Stdout {
		return e.out.Close()
	}
	return nil
}

func newOutputEncoder() *encoder {
	var o io.WriteCloser
	o = os.Stdout

	// Redirect output to the file at flags.output, if the flag was provided.
	if flags.output != "" {
		of, err := os.Create(flags.output)
		if err != nil {
			log.Fatalf("Couldn't create the output file at %s: %v", flags.output, err)
		}
		o = of
	}

	if flags.format != "plaintext" {
		log.Fatalf("Currently only the plaintext output format is supported.")
	}
	return &encoder{
		out: o,
	}
}
