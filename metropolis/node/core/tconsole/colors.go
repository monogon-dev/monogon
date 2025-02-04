// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package tconsole

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// Terminal is a supported terminal kind. This is a simplistic alternative to
// terminfo.
type Terminal string

const (
	// TerminalLinux is a linux VTY console.
	TerminalLinux Terminal = "linux"
	// TerminalGeneric is any other terminal that is supported by tcell.
	TerminalGeneric Terminal = "generic"
)

// color type separate from tcell.Color. This is necessary because tcell does not
// support color remapping in Linux VTYs, only static colors, and we have to
// build our own color system on top of it.
//
// Our color type is abstract, it denotes a given colour in terms of human
// perception, not an actual colour code.
type color string

const (
	// colorBlack is a deep black, suitable for use as a foreground.
	colorBlack color = "black"
	// colorPink is a light pink, suitable for use as a background.
	colorPink color = "pink"
	// colorBlue is a light blue, suitable for use as a background.
	colorBlue color = "blue"
)

// colorDef is an entry for a palette and defines a tcell color to be used from
// tcell, and an optional RGB colour to remap the given tcell.Color into.
type colorDef struct {
	ansiColor tcell.Color
	rgbColor  uint32
}

type palette map[color]colorDef

var (
	// paletteLinux is the full-colour palette used on Linux consoles, with exact RGB
	// color definitions.
	paletteLinux = palette{
		colorBlack: colorDef{tcell.ColorBlack + 0, 0x000000},
		colorPink:  colorDef{tcell.ColorBlack + 1, 0xeedfda},
		colorBlue:  colorDef{tcell.ColorBlack + 2, 0x919ba7},
	}
	// paletteGeneric is a fallback palette used on systems which do not support
	// colour remapping and only have 2 colours: black and white.
	paletteGeneric = palette{
		colorBlack: colorDef{tcell.ColorBlack, 0},
		colorPink:  colorDef{tcell.ColorWhite, 0},
		colorBlue:  colorDef{tcell.ColorWhite, 0},
	}
)

// linuxConsoleOverrideColor uses Linux-specific ANSI OSC codes to remap a given
// colour to an RGB value. See console_codes(4) for more information.
//
// The function returns a string which must be sent to the terminal.
func linuxConsoleOverrideColor(co tcell.Color, rgb uint32) string {
	return fmt.Sprintf("\x1b]P%x%06x", int(co-tcell.ColorBlack), rgb)
}

// setupLinuxConsole configures a given palette with rgcColor date to be
// displayed on the user screen.
//
// The function returns a string which must be sent to the terminal.
func (p *palette) setupLinuxConsole() string {
	res := ""
	for _, v := range *p {
		res += linuxConsoleOverrideColor(v.ansiColor, v.rgbColor)
	}
	return res
}

// color retrieves a tcell.Color to use for a given color, respecting the
// currently set palette.
func (c *Console) color(col color) tcell.Color {
	return c.palette[col].ansiColor
}
