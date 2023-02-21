package main

import (
	"fmt"

	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	fmt.Println("Monogon BMaaS Agent started")
	report, errs := gatherHWReport()
	// Just print the report for now
	fmt.Println(prototext.Format(report))
	fmt.Println("Encountered errors:", errs)
}
