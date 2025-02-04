// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package logtree

import (
	"fmt"

	"source.monogon.dev/go/logging"
	lpb "source.monogon.dev/osbase/logtree/proto"
)

func SeverityFromProto(s lpb.LeveledLogSeverity) (logging.Severity, error) {
	switch s {
	case lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_INFO:
		return logging.INFO, nil
	case lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_WARNING:
		return logging.WARNING, nil
	case lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_ERROR:
		return logging.ERROR, nil
	case lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_FATAL:
		return logging.FATAL, nil
	default:
		return "", fmt.Errorf("unknown severity value %d", s)
	}
}

func SeverityToProto(s logging.Severity) lpb.LeveledLogSeverity {
	switch s {
	case logging.INFO:
		return lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_INFO
	case logging.WARNING:
		return lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_WARNING
	case logging.ERROR:
		return lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_ERROR
	case logging.FATAL:
		return lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_FATAL
	default:
		return lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_INVALID
	}
}
