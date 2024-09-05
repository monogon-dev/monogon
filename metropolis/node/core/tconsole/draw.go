package tconsole

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/uniseg"
)

// drawText draws a single line of text from left to right, starting at x and y.
func (c *Console) drawText(x, y int, text string, style tcell.Style) int {
	g := uniseg.NewGraphemes(text)
	xi := 0
	for g.Next() {
		runes := g.Runes()
		c.screen.SetContent(x+xi, y, runes[0], runes[1:], style)
		xi += 1
	}
	return xi
}

// drawTextCentered draw a single line of text from left to right, starting at a
// position so that the center of the line ends up at x and y.
func (c *Console) drawTextCentered(x, y int, text string, style tcell.Style) {
	g := uniseg.NewGraphemes(text)
	var runes [][]rune
	for g.Next() {
		runes = append(runes, g.Runes())
	}

	x -= len(runes) / 2

	for _, r := range runes {
		c.screen.SetContent(x, y, r[0], r[1:], style)
		x += 1
	}
}

// fillRectangle fills a given rectangle [x0,x1) [y0,y1) with empty space of a
// given style.
func (c *Console) fillRectangle(x0, x1, y0, y1 int, style tcell.Style) {
	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			c.screen.SetContent(x, y, ' ', nil, style)
		}
	}
}

const logo = `
             _g@@@@g_         _g@@@@g_             
           _@@@@@@@@@@a     g@@@@@@@@@@b           
          @@@@@@@@@@@@@@___@@@@@@@@@@@@@@          
         @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@         
        @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@        
       g@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@g       
      g@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@g      
     g@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@g     
    ;@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@,    
    |@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@"    
     @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@     
      %@@@@@@@@@@P<@@@@@@@@@@@@@>%@@@@@@@@@@P      
        "+B@@B>'    "@@@@@@@@B"    '<B@@BP"        
                        '"                         
`

// drawLogo draws the Monogon logo so that its top left corner is at x, y.
func (c *Console) drawLogo(x, y int, style tcell.Style) {
	for i, line := range strings.Split(logo, "\n") {
		c.drawText(x, y+i, line, style)
	}
}

// split calculates a mid-point in the [0, capacity) domain so that it splits it
// into two parts fairly with minA and minB used as minimum size hints for each
// section.
func split(capacity, minA, minB int) int {
	slack := capacity - (minA + minB)
	propA := float64(minA) / float64(minA+minB)
	slackA := int(propA * float64(slack))
	return minA + slackA
}

// center calculates a point at which to start drawing a 'size'-sized element in
// a 'capacity'-sized container so that it ends in the middle of said container.
func center(capacity, size int) int {
	return (capacity - size) / 2
}
