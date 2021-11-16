package cluster

import (
	"fmt"
	"io"
	"strings"

	"source.monogon.dev/metropolis/pkg/logbuffer"
)

// prefixedStdio is a io.ReadWriter which splits written bytes into lines,
// prefixes them with some known prefix, and spits them to os.Stdout.
//
// io.Reader is implemented for compatibility with code which expects an
// io.ReadWriter, but always returns EOF.
type prefixedStdio struct {
	*logbuffer.LineBuffer
}

// newPrefixedStdio returns a prefixedStdio that prefixes all lines with <num>|,
// used to distinguish different VMs used within the launch codebase.
func newPrefixedStdio(num int) prefixedStdio {
	return prefixedStdio{
		logbuffer.NewLineBuffer(1024, func(l *logbuffer.Line) {
			s := strings.TrimSpace(l.String())
			// TODO(q3k): don't just skip lines containing escape sequences, strip the
			// sequences out. Or stop parsing qemu logs and instead dial log endpoint in
			// spawned nodes.
			if strings.Contains(s, "\u001b") {
				return
			}
			fmt.Printf("%02d| %s\n", num, s)
		}),
	}
}

func (p prefixedStdio) Read(_ []byte) (int, error) {
	return 0, io.EOF
}
