package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type HelpWidget struct {
	name string
	x, y int
	w, h int
}

func NewHelpWidget(name string, x, y, w, h int) *HelpWidget {
	return &HelpWidget{name: name, x: x, y: y, w: w, h: h}
}

func (self *HelpWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(self.name, self.x, self.y, self.w, self.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set view parameter
		v.Frame = false
		v.BgColor = gocui.ColorBlue
		v.FgColor = gocui.ColorBlack

		// Initialize view
		v.Clear()

		// Display help information
		fmt.Fprintf(v, " k: Down, j: Up, /: Search, space: Toggle ack")
	}
	return nil
}
