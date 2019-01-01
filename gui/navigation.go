package gui

import (
	"github.com/jroimartin/gocui"
)

type NavigationWidget struct {
	name string
	x, y int
	w, h int
}

func NewNavigationWidget(name string, x, y, w, h int) *NavigationWidget {
	return &NavigationWidget{name: name, x: x, y: y, w: w, h: h}
}

func (self *NavigationWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(self.name, self.x, self.y, self.w, self.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set view parameter
		v.Editable = true

		// Initialize view
		v.Clear()
	}
	return nil
}
