package main

import (
	"context"
	"fmt"

	"source.monogon.dev/osbase/bringup"
)

func main() {
	bringup.Runnable(func(ctx context.Context) error {
		return fmt.Errorf("this is an error")
	}).Run()
}
