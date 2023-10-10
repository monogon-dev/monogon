// Package mflags implements custom flags for use in monogon projects.
// It provides them to be used like normal flag.$Var and registers the
// required flag functions.
package mflags

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

// Limiter configures a *rate.Limiter as a flag.
func Limiter(l **rate.Limiter, name, defval, help string) {
	syntax := "'duration,count' eg. '2m,10' for a 10-sized bucket refilled at one token every 2 minutes"
	help = help + fmt.Sprintf(" (default: %q, syntax: %s)", defval, syntax)
	flag.Func(name, help, func(val string) error {
		if val == "" {
			val = defval
		}
		parts := strings.Split(val, ",")
		if len(parts) != 2 {
			return fmt.Errorf("invalid syntax, want: %s", syntax)
		}
		duration, err := time.ParseDuration(parts[0])
		if err != nil {
			return fmt.Errorf("invalid duration: %w", err)
		}
		refill, err := strconv.ParseUint(parts[1], 10, 31)
		if err != nil {
			return fmt.Errorf("invalid refill rate: %w", err)
		}
		*l = rate.NewLimiter(rate.Every(duration), int(refill))
		return nil
	})
	flag.Set(name, defval)
}
