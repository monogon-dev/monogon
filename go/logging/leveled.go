// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package logging

// Leveled is a generic interface for glog-style logging. There are four
// hardcoded log severities, in increasing order: INFO, WARNING, ERROR, FATAL.
// Logging at a certain severity level logs not only to consumers expecting data
// at that severity level, but also all lower severity levels. For example, an
// ERROR log will also be passed to consumers looking at INFO or WARNING logs.
type Leveled interface {
	// Info logs at the INFO severity. Arguments are handled in the manner of
	// fmt.Print, a terminating newline is added if missing.
	Info(args ...any)
	// Infof logs at the INFO severity. Arguments are handled in the manner of
	// fmt.Printf, a terminating newline is added if missing.
	Infof(format string, args ...any)

	// Warning logs at the WARNING severity. Arguments are handled in the manner of
	// fmt.Print, a terminating newline is added if missing.
	Warning(args ...any)
	// Warningf logs at the WARNING severity. Arguments are handled in the manner of
	// fmt.Printf, a terminating newline is added if missing.
	Warningf(format string, args ...any)

	// Error logs at the ERROR severity. Arguments are handled in the manner of
	// fmt.Print, a terminating newline is added if missing.
	Error(args ...any)
	// Errorf logs at the ERROR severity. Arguments are handled in the manner of
	// fmt.Printf, a terminating newline is added if missing.
	Errorf(format string, args ...any)

	// Fatal logs at the FATAL severity and aborts the current program. Arguments are
	// handled in the manner of fmt.Print, a terminating newline is added if missing.
	Fatal(args ...any)
	// Fatalf logs at the FATAL severity and aborts the current program. Arguments are
	// handled in the manner of fmt.Printf, a terminating newline is added if missing.
	Fatalf(format string, args ...any)

	// V returns a VerboseLeveledLogger at a given verbosity level. These verbosity
	// levels can be dynamically set and unset on a package-granular level by consumers
	// of the LeveledLogger logs. The returned value represents whether logging at the
	// given verbosity level was active at that time, and as such should not be a long-
	// lived object in programs. This construct is further refered to as 'V-logs'.
	V(level VerbosityLevel) VerboseLeveled

	// WithAddedStackDepth returns the same LeveledLogger, but adjusted with an
	// additional 'extra stack depth' which will be used to skip a given number of
	// stack/call frames when determining the location where the error originated.
	// For example, WithStackDepth(1) will return a logger that will skip one
	// stack/call frame. Then, with function foo() calling function helper() which
	// in turns call l.Infof(), the log line will be emitted with the call site of
	// helper() within foo(), instead of the default behaviour of logging the
	// call site of Infof() within helper().
	//
	// This is useful for functions which somehow wrap loggers in helper functions,
	// for example to expose a slightly different API.
	WithAddedStackDepth(depth int) Leveled
}

// VerbosityLevel is a verbosity level defined for V-logs. This can be changed
// programmatically per Go package. When logging at a given VerbosityLevel V, the
// current level must be equal or higher to V for the logs to be recorded.
// Conversely, enabling a V-logging at a VerbosityLevel V also enables all logging
// at lower levels [Int32Min .. (V-1)].
type VerbosityLevel int32

type VerboseLeveled interface {
	// Enabled returns if this level was enabled. If not enabled, all logging into this
	// logger will be discarded immediately. Thus, Enabled() can be used to check the
	// verbosity level before performing any logging:
	//    if l.V(3).Enabled() { l.Info("V3 is enabled") }
	// or, in simple cases, the convenience function .Info can be used:
	//    l.V(3).Info("V3 is enabled")
	// The second form is shorter and more convenient, but more expensive, as its
	// arguments are always evaluated.
	Enabled() bool
	// Info is the equivalent of a LeveledLogger's Info call, guarded by whether this
	// VerboseLeveledLogger is enabled.
	Info(args ...any)
	// Infof is the equivalent of a LeveledLogger's Infof call, guarded by whether this
	// VerboseLeveledLogger is enabled.
	Infof(format string, args ...any)
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
	// SeverityAtLeast maps a given severity to a list of severities that at that
	// severity or higher. In other words, SeverityAtLeast[X] returns a list of
	// severities that might be seen in a log at severity X.
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

// Valid returns whether true if this severity is one of the known levels
// (INFO, WARNING, ERROR or FATAL), false otherwise.
func (s Severity) Valid() bool {
	switch s {
	case INFO, WARNING, ERROR, FATAL:
		return true
	default:
		return false
	}
}
