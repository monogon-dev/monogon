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

	apb "git.monogon.dev/source/nexantic.git/metropolis/proto/api"
)

// LeveledLogger is a generic interface for glog-style logging. There are four hardcoded log severities, in increasing
// order: INFO, WARNING, ERROR, FATAL. Logging at a certain severity level logs not only to consumers expecting data
// at that severity level, but also all lower severity levels. For example, an ERROR log will also be passed to
// consumers looking at INFO or WARNING logs.
type LeveledLogger interface {
	// Info logs at the INFO severity. Arguments are handled in the manner of fmt.Print, a terminating newline is added
	// if missing.
	Info(args ...interface{})
	// Infof logs at the INFO severity. Arguments are handled in the manner of fmt.Printf, a terminating newline is
	// added if missing.
	Infof(format string, args ...interface{})

	// Warning logs at the WARNING severity. Arguments are handled in the manner of fmt.Print, a terminating newline is
	// added if missing.
	Warning(args ...interface{})
	// Warningf logs at the WARNING severity. Arguments are handled in the manner of fmt.Printf, a terminating newline
	// is added if missing.
	Warningf(format string, args ...interface{})

	// Error logs at the ERROR severity. Arguments are handled in the manner of fmt.Print, a terminating newline is
	// added if missing.
	Error(args ...interface{})
	// Errorf logs at the ERROR severity. Arguments are handled in the manner of fmt.Printf, a terminating newline is
	// added if missing.
	Errorf(format string, args ...interface{})

	// Fatal logs at the FATAL severity and aborts the current program. Arguments are handled in the manner of
	// fmt.Print, a terminating newline is added if missing.
	Fatal(args ...interface{})
	// Fatalf logs at the FATAL severity and aborts the current program. Arguments are handled in the manner of
	// fmt.Printf, a terminating newline is added if missing.
	Fatalf(format string, args ...interface{})

	// V returns a VerboseLeveledLogger at a given verbosity level. These verbosity levels can be dynamically set and
	// unset on a package-granular level by consumers of the LeveledLogger logs. The returned value represents whether
	// logging at the given verbosity level was active at that time, and as such should not be a long-lived object
	// in programs.
	// This construct is further refered to as 'V-logs'.
	V(level VerbosityLevel) VerboseLeveledLogger
}

// VerbosityLevel is a verbosity level defined for V-logs. This can be changed programmatically per Go package. When
// logging at a given VerbosityLevel V, the current level must be equal or higher to V for the logs to be recorded.
// Conversely, enabling a V-logging at a VerbosityLevel V also enables all logging at lower levels [Int32Min .. (V-1)].
type VerbosityLevel int32

type VerboseLeveledLogger interface {
	// Enabled returns if this level was enabled. If not enabled, all logging into this logger will be discarded
	// immediately.
	// Thus, Enabled() can be used to check the verbosity level before performing any logging:
	//    if l.V(3).Enabled() { l.Info("V3 is enabled") }
	// or, in simple cases, the convenience function .Info can be used:
	//    l.V(3).Info("V3 is enabled")
	// The second form is shorter and more convenient, but more expensive, as its arguments are always evaluated.
	Enabled() bool
	// Info is the equivalent of a LeveledLogger's Info call, guarded by whether this VerboseLeveledLogger is enabled.
	Info(args ...interface{})
	// Infof is the equivalent of a LeveledLogger's Infof call, guarded by whether this VerboseLeveledLogger is enabled.
	Infof(format string, args ...interface{})
}

// Severity is one of the severities as described in LeveledLogger.
type Severity string

const (
	INFO    Severity = "I"
	WARNING Severity = "W"
	ERROR   Severity = "E"
	FATAL   Severity = "F"
)

var (
	// SeverityAtLeast maps a given severity to a list of severities that at that severity or higher. In other words,
	// SeverityAtLeast[X] returns a list of severities that might be seen in a log at severity X.
	SeverityAtLeast = map[Severity][]Severity{
		INFO:    {INFO, WARNING, ERROR, FATAL},
		WARNING: {WARNING, ERROR, FATAL},
		ERROR:   {ERROR, FATAL},
		FATAL:   {FATAL},
	}
)

func (s Severity) AtLeast(other Severity) bool {
	for _, el := range SeverityAtLeast[other] {
		if el == s {
			return true
		}
	}
	return false
}

func SeverityFromProto(s apb.LeveledLogSeverity) (Severity, error) {
	switch s {
	case apb.LeveledLogSeverity_INFO:
		return INFO, nil
	case apb.LeveledLogSeverity_WARNING:
		return WARNING, nil
	case apb.LeveledLogSeverity_ERROR:
		return ERROR, nil
	case apb.LeveledLogSeverity_FATAL:
		return FATAL, nil
	default:
		return "", fmt.Errorf("unknown severity value %d", s)
	}
}

func (s Severity) ToProto() apb.LeveledLogSeverity {
	switch s {
	case INFO:
		return apb.LeveledLogSeverity_INFO
	case WARNING:
		return apb.LeveledLogSeverity_WARNING
	case ERROR:
		return apb.LeveledLogSeverity_ERROR
	case FATAL:
		return apb.LeveledLogSeverity_FATAL
	default:
		return apb.LeveledLogSeverity_INVALID
	}
}
