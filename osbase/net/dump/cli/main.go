package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/encoding/prototext"

	netdump "source.monogon.dev/osbase/net/dump"
)

func main() {
	netconf, _, err := netdump.Dump()
	if err != nil {
		log.Fatalf("failed to dump network configuration: %v", err)
	}
	fmt.Println(prototext.Format(netconf))
}
