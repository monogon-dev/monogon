// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
