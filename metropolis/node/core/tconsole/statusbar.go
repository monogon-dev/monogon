// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package tconsole

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

// draw button at coordinates containing text, with the left side of the button
// at (x, y). The number of columns used up by the button is returned.
func (c *Console) button(x, y int, caption string, selected bool, sty tcell.Style) int {
	fg, bg, _ := sty.Decompose()
	styInv := sty.Background(fg).Foreground(bg)

	xi := 1
	if selected {
		c.screen.SetContent(x+xi, y, tcell.RuneBlock, nil, sty)
		xi += 1
		xi += c.drawText(x+xi, y, caption, styInv)
		c.screen.SetContent(x+xi, y, tcell.RuneBlock, nil, sty)
		xi += 1
	} else {
		c.screen.SetContent(x+xi, y, ' ', nil, sty)
		xi += 1
		xi += c.drawText(x+xi, y, caption, sty)
		c.screen.SetContent(x+xi, y, ' ', nil, sty)
		xi += 1
	}
	return xi
}

// statusBar draw the main status bar at the bottom of the screen, containing
// page switching buttons and a clock.
func (c *Console) statusBar(active int, opts ...string) {
	sty1 := tcell.StyleDefault.Background(c.color(colorBlue)).Foreground(c.color(colorBlack))
	sty2 := tcell.StyleDefault.Background(c.color(colorPink)).Foreground(c.color(colorBlack))
	x := 0
	x += c.drawText(x, c.height-1, " Page (tab to switch): ", sty1)
	for i, opt := range opts {
		x += c.button(x, c.height-1, opt, i == active, sty2)
	}

	c.drawText(c.width-len(time.DateTime)-1, c.height-1, time.Now().Format(time.DateTime), sty1)
}
