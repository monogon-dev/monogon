// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package tconsole

import "github.com/gdamore/tcell/v2"

// pageLogsData encompasses all data to be shown within the logs page.
type pageLogsData struct {
	// log lines, simple deque with the newest log lines appended to the end.
	lines []string
}

func (p *pageLogsData) appendLine(s string) {
	p.lines = append(p.lines, s)
}

// compactData ensures that there's no more lines stored than maxlines by
// discarding the oldest lines.
func (p *pageLogsData) compactData(maxlines int) {
	if extra := len(p.lines) - maxlines; extra > 0 {
		p.lines = p.lines[extra:]
	}
}

// pageLogs renders the logs page to the user given pageLogsData.
func (c *Console) pageLogs(data *pageLogsData) {
	c.screen.Clear()
	sty1 := tcell.StyleDefault.Background(c.color(colorPink)).Foreground(c.color(colorBlack))
	sty2 := tcell.StyleDefault.Background(c.color(colorBlue)).Foreground(c.color(colorBlack))

	// Draw frame.
	c.fillRectangle(0, c.width, 0, c.height, sty2)
	c.fillRectangle(1, c.width-1, 1, c.height-2, sty1)

	// Inner log area size.
	nlines := (c.height - 2) - 1
	linelen := (c.width - 1) - 1

	// Compact and draw log lines.
	data.compactData(nlines)
	for y := 0; y < nlines; y++ {
		if y < len(data.lines) {
			line := data.lines[y]
			if len(line) > linelen {
				line = line[:linelen]
			}
			c.drawText(1, 1+y, line, sty1)
		}
	}
}
