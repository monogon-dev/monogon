//go:build tools
// +build tools

package analysis

import (
	_ "4d63.com/gocheckcompilerdirectives/checkcompilerdirectives"
	_ "github.com/corverroos/commentwrap/cmd/commentwrap"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
