package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
	"github.com/tawaku/zbxcui/api"
	"log"
	"os"
	"strings"
	_time "time"
)

type Dashboard struct {
	gui              *gocui.Gui
	client           *api.Client
	eventWidget      *EventWidget
	navigationWidget *NavigationWidget
	done             chan struct{}
	logger           *log.Logger
}

func NewDashboard(c *api.Client, l *log.Logger) *Dashboard {
	// Initialize
	d := new(Dashboard)

	// Set client
	d.client = c

	// Set widget
	d.eventWidget = new(EventWidget)

	// Set channel
	d.done = make(chan struct{})

	// Set logger
	if l != nil {
		d.logger = l
	} else {
		d.logger = log.New(os.Stdout, "[ZBXCUI|GUI]", log.LstdFlags|log.LUTC)
	}

	return d
}

func (self *Dashboard) eventGet() ([]api.EventGetResult, error) {
	// Get event
	if r, err := self.client.EventGet(); err != nil {
		return nil, errors.Wrap(err, "Failed to get event.")
	} else {
		return r, nil
	}
}

func (self *Dashboard) Run() {
	// Set gui
	var err error
	self.gui, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(errors.Wrap(err, "Failed to create gui instance."))
	}
	defer self.gui.Close()

	// Set widget
	const navHeight = 3
	e, _ := self.eventGet()
	maxX, maxY := self.gui.Size()
	self.eventWidget = NewEventWidget("Event", 0, 0, maxX, maxY-navHeight, e)
	self.navigationWidget = NewNavigationWidget("Navigation", 0, maxY-navHeight, maxX, maxY)
	self.gui.SetManager(self.eventWidget, self.navigationWidget)

	// Enable cursor
	self.gui.Cursor = true

	// Set key bindigs
	if err := self.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, self.quit); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}
	if err := self.gui.SetKeybinding("Event", 'q', gocui.ModNone, self.quit); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}
	if err := self.gui.SetKeybinding("Event", gocui.KeyArrowUp, gocui.ModNone, self.up); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}
	if err := self.gui.SetKeybinding("Event", gocui.KeyArrowDown, gocui.ModNone, self.down); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}
	if err := self.gui.SetKeybinding("Event", 'k', gocui.ModNone, self.up); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}
	if err := self.gui.SetKeybinding("Event", 'j', gocui.ModNone, self.down); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}
	if err := self.gui.SetKeybinding("Event", gocui.KeySpace, gocui.ModNone, self.toggle); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}
	if err := self.gui.SetKeybinding("Event", 'r', gocui.ModNone, self.update); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}
	if err := self.gui.SetKeybinding("Event", '/', gocui.ModNone, self.inNav); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}
	if err := self.gui.SetKeybinding("Navigation", gocui.KeyEnter, gocui.ModNone, self.filter); err != nil {
		log.Panicln(errors.Wrap(err, "Failed to bind key."))
	}

	// Update view periodicaly
	go self.autoRefresh()

	// Start main loop
	if err := self.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(errors.Wrap(err, "Failed to start main loop."))
	}
}

func (self *Dashboard) autoRefresh() {
	for {
		select {
		case <-self.done:
			return
		case <-_time.After(5000 * _time.Millisecond):
			self.gui.Update(func(g *gocui.Gui) error {
				if v, err := g.View("Event"); err != nil {
					log.Panicln(errors.Wrap(err, "Failed to update view."))
				} else {
					self.update(g, v)
				}
				return nil
			})
		}
	}
}

func (self *Dashboard) quit(g *gocui.Gui, v *gocui.View) error {
	close(self.done)
	return gocui.ErrQuit
}

func (self *Dashboard) up(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		if y := bufferCursor(v); y > 1 {
			v.MoveCursor(0, -1, false)
		}
	}
	return nil
}

func (self *Dashboard) down(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		maxY := len(v.ViewBufferLines())
		if y := bufferCursor(v); y < maxY-2 {
			v.MoveCursor(0, 1, false)
		}
	}
	return nil
}

func (self *Dashboard) toggle(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		// Get line on cursor position
		y := bufferCursor(v)
		// Get corresponding event ID from cursor position
		i := self.eventWidget.findEventidByPosition(y)
		if i == "" {
			// Do nothing if the cursor doesn`t point except for event line.
			return nil
		}
		// Reverse unfolded flag if acknowledgement is configured.
		if e, err := self.eventWidget.findEventById(i); err != nil {
			return errors.Wrap(err, "Failed to find Event.")
		} else {
			if e.Acknowledged == 1 {
				self.eventWidget.properties[i].unfolded = !self.eventWidget.properties[i].unfolded
			}
		}
		// Update view
		self.eventWidget.render(v)
	}
	return nil
}

func (self *Dashboard) update(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		// Get event
		e, _ := self.eventGet()
		// Update the events
		self.eventWidget.mu.Lock()
		self.eventWidget.events = e
		self.eventWidget.mu.Unlock()
		// Update view with the result
		self.eventWidget.render(v)
	}
	return nil
}

func (self *Dashboard) inNav(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		// Switch to navigation widget
		vNav, _ := g.SetCurrentView(self.navigationWidget.name)
		// Clear previous filter key
		vNav.Clear()
		vNav.SetCursor(0, 0)
	}
	return nil
}

func (self *Dashboard) filter(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		// Get keyword to filter
		if k, err := v.Line(bufferCursor(v)); err != nil {
			// Remove filter
			for _, e := range self.eventWidget.events {
				self.eventWidget.properties[e.Eventid].filtered = false
			}
		} else {
			for _, e := range self.eventWidget.events {
				notFiltered :=
					strings.Contains(severity(e), k) ||
						strings.Contains(status(e), k) ||
						strings.Contains(info(e), k) ||
						strings.Contains(time(e), k) ||
						strings.Contains(age(e), k) ||
						strings.Contains(ack(e), k) ||
						strings.Contains(host(e), k) ||
						strings.Contains(name(e), k)
				self.eventWidget.properties[e.Eventid].filtered = !notFiltered
			}
		}
		// Update view with filtered event
		if ev, err := g.View(self.eventWidget.name); err != nil {
			return errors.Wrap(err, "Failed to get event view.")
		} else {
			self.eventWidget.render(ev)
		}
	}
	return self.outNav(g, v)
}

func (self *Dashboard) outNav(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		// Switch to event widget
		vEvt, _ := g.SetCurrentView(self.eventWidget.name)
		vEvt.SetCursor(0, 1)
	}
	return nil
}

func bufferCursor(v *gocui.View) int {
	// Get line on cursor position
	_, viewY := v.Cursor()
	if line, err := v.Line(viewY); err != nil {
		return -1
	} else {
		// Get position of the line in view's buffer
		for i, l := range v.ViewBufferLines() {
			if line == l {
				return i
			}
		}
		return -1
	}
}
