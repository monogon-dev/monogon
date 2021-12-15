package clicontext

import (
	"context"
	"os"
	"os/signal"
)

// WithInterrupt returns a context for use in a command-line utility. It gets
// cancelled if the user interrupts the command, for example by pressing
// Ctrl+C.
func WithInterrupt(parent context.Context) context.Context {
	ctx, cancel := context.WithCancel(parent)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()
	return ctx
}
