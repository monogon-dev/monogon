package main

import (
	"context"

	"source.monogon.dev/metropolis/pkg/pstore"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// dumpAndCleanPstore dumps all files accumulated in the pstore into the log
// and clears them from the pstore. This allows looking at these logs and also
// keeps the pstore from overflowing the generally limited storage it has.
func dumpAndCleanPstore(ctx context.Context) error {
	logger := supervisor.Logger(ctx)
	dumps, err := pstore.GetKmsgDumps()
	if err != nil {
		logger.Errorf("Failed to recover logs from pstore: %v", err)
		return nil
	}
	for _, dump := range dumps {
		logger.Errorf("Recovered log from %v at %v. Reconstructed log follows.", dump.Reason, dump.OccurredAt)
		for _, line := range dump.Lines {
			logger.Warning(line)
		}
	}
	cleanErr := pstore.ClearAll()
	if cleanErr != nil {
		logger.Errorf("Failed to clear pstore: %v", err)
	}
	// Retrying this is extremely unlikely to result in any change and is most
	// likely just going to generate large amounts of useless logs obscuring
	// errors.
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
