package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"source.monogon.dev/metropolis/node/core/identity"
	apb "source.monogon.dev/metropolis/proto/api"
)

type encoder struct {
	out io.WriteCloser
}

func (e *encoder) writeNodeID(n *apb.Node) error {
	id := identity.NodeID(n.Pubkey)
	_, err := fmt.Fprintf(e.out, "%s\n", id)
	return err
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
