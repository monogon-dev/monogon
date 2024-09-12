package logging

import (
	"fmt"
	"io"
	"runtime"
	"strings"
	"time"
)

// FunctionBackend is the simplest Backend (Leveled implementation). It
// synchronously forwards logged entries to a function.
type FunctionBackend struct {
	depth     int
	fn        func(severity Severity, msg string)
	verbosity VerbosityLevel
}

// NewFunctionBackend returns a FunctionBackend (Leveled implementation) which
// will call the given function on every log entry synchronously.
func NewFunctionBackend(fn func(severity Severity, msg string)) *FunctionBackend {
	return &FunctionBackend{
		fn: fn,
	}
}

func (w *FunctionBackend) Info(args ...any) {
	w.fn(INFO, fmt.Sprint(args...))
}
func (w FunctionBackend) Infof(format string, args ...any) {
	w.fn(INFO, fmt.Sprintf(format, args...))
}
func (w *FunctionBackend) Warning(args ...any) {
	w.fn(WARNING, fmt.Sprint(args...))
}
func (w *FunctionBackend) Warningf(format string, args ...any) {
	w.fn(WARNING, fmt.Sprintf(format, args...))
}
func (w *FunctionBackend) Error(args ...any) {
	w.fn(ERROR, fmt.Sprint(args...))
}
func (w *FunctionBackend) Errorf(format string, args ...any) {
	w.fn(ERROR, fmt.Sprintf(format, args...))
}
func (w *FunctionBackend) Fatal(args ...any) {
	w.fn(FATAL, fmt.Sprint(args...))
}
func (w *FunctionBackend) Fatalf(format string, args ...any) {
	w.fn(FATAL, fmt.Sprintf(format, args...))
}

type verboseFunctionBackend struct {
	backend *FunctionBackend
	enabled bool
}

func (w *FunctionBackend) V(level VerbosityLevel) VerboseLeveled {
	return &verboseFunctionBackend{
		backend: w,
		enabled: level > w.verbosity,
	}
}

func (w *FunctionBackend) WithAddedStackDepth(depth int) Leveled {
	w2 := *w
	w.depth += depth
	return &w2
}

func (v *verboseFunctionBackend) Enabled() bool {
	return v.enabled
}

func (v *verboseFunctionBackend) Info(args ...any) {
	if !v.enabled {
		return
	}
	v.backend.fn(INFO, fmt.Sprint(args...))
}

func (v *verboseFunctionBackend) Infof(format string, args ...any) {
	if !v.enabled {
		return
	}
	v.backend.fn(INFO, fmt.Sprintf(format, args...))
}

// WriterBackend is a Backend (Leveled implementation) which outputs log entries
// to a writer using a given Formatter.
type WriterBackend struct {
	FunctionBackend
	// Formatter is used to turn a log entry alongside metadata into a string. By
	// default, it's set to LeveledFormatterGlog.
	Formatter       LeveledFormatter
	out             io.Writer
	MinimumSeverity Severity
}

// LeveledFormatter is a function to turn a leveled log entry into a string which
// can be output to a user.
type LeveledFormatter func(file string, line int, time time.Time, severity Severity, msg string) string

// LeveledFormatterGlog implements LeveledFormatter in a glog/klog-compatible
// way.
func LeveledFormatterGlog(file string, line int, ts time.Time, severity Severity, msg string) string {
	// TODO(q3k): unify with //osbase/logtree.LeveledPayload.String.
	_, month, day := ts.Date()
	hour, minute, second := ts.Clock()
	nsec := ts.Nanosecond() / 1000

	res := fmt.Sprintf("%s%02d%02d %02d:%02d:%02d.%06d %s:%d] ", severity, month, day, hour, minute, second, nsec, file, line)
	res += msg
	return res
}

// NewWriterBackend constructs a WriterBackend (Leveled implementation) which
// writes glog/klog-style entries to the given Writer.
func NewWriterBackend(w io.Writer) *WriterBackend {
	res := &WriterBackend{
		Formatter:       LeveledFormatterGlog,
		out:             w,
		MinimumSeverity: INFO,
	}
	res.FunctionBackend.fn = res.log
	return res
}

func (w *WriterBackend) log(severity Severity, msg string) {
	if !severity.AtLeast(w.MinimumSeverity) {
		return
	}
	_, file, line, ok := runtime.Caller(2 + w.depth)
	if !ok {
		file = "???"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	res := w.Formatter(file, line, time.Now(), severity, msg)
	w.out.Write([]byte(res + "\n"))
}
