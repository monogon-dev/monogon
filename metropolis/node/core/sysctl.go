package main

import (
	"context"
	"strconv"

	"source.monogon.dev/osbase/supervisor"
	"source.monogon.dev/osbase/sysctl"
)

func nodeSysctls(ctx context.Context) error {
	const vmMaxMapCount = 2<<30 - 1
	options := sysctl.Options{
		// We increase the max mmap count to nearly the maximum, as it gets
		// accounted by the cgroup memory limit.
		"vm.max_map_count": strconv.Itoa(vmMaxMapCount),
	}

	if err := options.Apply(); err != nil {
		return err
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
