// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package tconsole

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"

	mversion "source.monogon.dev/metropolis/version"
	"source.monogon.dev/version"
)

//go:embed build/copyright_line.txt
var copyrightLine string

// pageStatusData encompasses all data to be shown within the status page.
type pageStatusData struct {
	netAddr     string
	roles       string
	id          string
	fingerprint string
}

// pageStatus renders the status page to the user given pageStatusData.
func (c *Console) pageStatus(d *pageStatusData) {
	c.screen.Clear()
	sty1 := tcell.StyleDefault.Background(c.color(colorPink)).Foreground(c.color(colorBlack))
	sty2 := tcell.StyleDefault.Background(c.color(colorBlue)).Foreground(c.color(colorBlack))

	logoWidth := len(strings.Split(logo, "\n")[1])
	logoHeight := len(strings.Split(logo, "\n"))

	// Vertical split between top copyright string and main display part.
	splitV := split(c.height, 4, logoHeight)

	// Colour the split.
	c.fillRectangle(0, c.width, 0, splitV, sty1)
	c.fillRectangle(0, c.width, splitV, c.height, sty2)

	// Draw the top part.
	c.drawTextCentered(c.width/2, splitV/2, "Monogon Cluster Operating System", sty1)
	c.drawTextCentered(c.width/2, splitV/2+1, copyrightLine, sty1)

	// Horizontal split between left logo and right status lines, a la 'fetch'.
	splitH := split(c.width, logoWidth, 60)

	// Status lines.
	lines := []string{
		fmt.Sprintf("Version: %s", version.Semver(mversion.Version)),
		fmt.Sprintf("Node ID: %s", d.id),
		fmt.Sprintf("CA fingerprint: %s", d.fingerprint),
		fmt.Sprintf("Management address: %s", d.netAddr),
		fmt.Sprintf("Roles: %s", d.roles),
	}
	// Calculate longest line.
	maxLine := 0
	for _, l := range lines {
		if len(l) > maxLine {
			maxLine = len(l)
		}
	}

	// If logo wouldn't fit, don't bother, save space for important data.
	drawLogo := true
	if splitH < logoWidth {
		drawLogo = false
		splitH = center(c.width, maxLine)
	}

	// Draw lines.
	for i, line := range lines {
		c.drawText(splitH, splitV+center(c.height-splitV, len(lines))+i, line, sty2)
	}

	// Draw logo.
	if drawLogo {
		c.drawLogo(splitH-logoWidth, splitV+center(c.height-splitV, logoHeight), sty2)
	}
}
