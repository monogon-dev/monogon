package main

import (
	"fmt"

	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	report, errs := gatherHWReport()
	// Just print the report for now
	fmt.Println(prototext.Format(report))
	fmt.Println("Encountered errors:", errs)
}
