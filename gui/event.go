package gui

import (
	"bytes"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
	"github.com/tawaku/zbxcui/api"
	"math"
	"strings"
	"sync"
	"text/tabwriter"
	_time "time"
)

type EventWidget struct {
	name       string
	x, y       int
	w, h       int
	mu         sync.Mutex
	events     []api.EventGetResult
	properties map[string]*EventProperty // Event ID is a key
}

type EventProperty struct {
	position int  // Position on event table
	filtered bool // If event is filtered or not
	unfolded bool // If acknowledgement is unfolded or not
}

func NewEventWidget(name string, x, y, w, h int, events []api.EventGetResult) *EventWidget {
	properties := make(map[string]*EventProperty)
	for _, e := range events {
		// Position property is initialized by 0
		properties[e.Eventid] = NewEventProperty(0, false, false)
	}
	return &EventWidget{name: name, x: x, y: y, w: w, h: h, events: events, properties: properties}
}

func NewEventProperty(position int, filtered, unfolded bool) *EventProperty {
	return &EventProperty{position: position, filtered: filtered, unfolded: unfolded}
}

func (self *EventWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(self.name, self.x, self.y, self.w-1, self.h-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set view parameter
		v.Title = self.name
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.SetCursor(0, 1)
		g.SetCurrentView(self.name)

		self.render(v)
	}
	return nil
}

func (self *EventWidget) render(v *gocui.View) {
	// Create event table
	buf := &bytes.Buffer{}
	pos := 0
	w := tabwriter.NewWriter(buf, 10, 4, 1, ' ', tabwriter.TabIndent|tabwriter.Debug)
	fmt.Fprintf(w, "%s", strings.Repeat(" ", 1))
	fmt.Fprintf(w, "Severity\tStatus\tInfo\tTime\tAge\tAck\tHost\tName\n")
	pos += 1
	for _, e := range self.events {
		if !self.properties[e.Eventid].filtered {
			// Print event items
			fmt.Fprintf(w, "%s", strings.Repeat(" ", 1))
			// fmt.Fprintf(w, "%d:", pos)
			fmt.Fprintf(w, "%s\t", severity(e))
			fmt.Fprintf(w, "%s\t", status(e))
			fmt.Fprintf(w, "%s\t", info(e))
			fmt.Fprintf(w, "%s\t", time(e))
			fmt.Fprintf(w, "%s\t", age(e))
			fmt.Fprintf(w, "%s\t", ack(e))
			fmt.Fprintf(w, "%s\t", host(e))
			fmt.Fprintf(w, "%s", name(e))
			// Store position
			self.properties[e.Eventid].position = pos
			// Unfold acknowledgement if unfolded property is flagged
			if self.properties[e.Eventid].unfolded {
				fmt.Fprintf(w, "\f")
				pos += 2
				fmt.Fprintf(w, "%s", strings.Repeat(" ", 2))
				fmt.Fprintf(w, "Time\tUser\tMessage\tUser action")
				acks := e.Acknowledges
				for _, a := range acks {
					fmt.Fprintf(w, "\n")
					pos += 1
					// Print acknowledgement items
					fmt.Fprintf(w, "%s", strings.Repeat(" ", 2))
					fmt.Fprintf(w, "%s\t", time(a))
					fmt.Fprintf(w, "%s\t", user(a))
					fmt.Fprintf(w, "%s\t", message(a))
					fmt.Fprintf(w, "%s", userAction(a))
				}
				fmt.Fprintf(w, "\f")
				pos += 2
			} else {
				fmt.Fprintf(w, "\n")
				pos += 1
			}
		}
	}
	w.Flush()

	// Render event table
	v.Clear()
	fmt.Fprintf(v, "%s", buf)
}

func severity(e api.EventGetResult) string {
	switch e.Event.Source {
	case 0:
		t := e.RelatedObjectValue().(api.Trigger)
		return t.PriorityValue()
	case 1:
		return "undefined"
	case 2:
		return "undefined"
	case 3:
		return "undefined"
	default:
		return "undefined"
	}
}
func status(e api.EventGetResult) string {
	return e.Event.ValueValue()
}
func info(e api.EventGetResult) string {
	switch e.Event.Source {
	case 0:
		t := e.RelatedObjectValue().(api.Trigger)
		return t.Error
	case 1:
		return "undefined"
	case 2:
		return "undefined"
	case 3:
		return "undefined"
	default:
		return "undefined"
	}
}
func time(i interface{}) string {
	switch v := i.(type) {
	case api.EventGetResult:
		return _time.Unix(v.Event.Clock, 0).Format("2006-01-02 15:04:05")
	case api.Acknowledge:
		return _time.Unix(v.Clock, 0).Format("2006-01-02 15:04:05")
	default:
		return "undefined"
	}
}
func age(e api.EventGetResult) string {
	d := _time.Since(_time.Unix(e.Event.Clock, 0))
	// Seconds
	sec := math.Mod(d.Seconds(), 60.0)
	r := fmt.Sprintf("%2.0fs", sec)
	// Minutes
	if d.Minutes() >= 1 {
		min := math.Mod(d.Minutes(), 60.0)
		r = fmt.Sprintf("%2.0fm", min) + r
	} else {
		r = fmt.Sprintf("%s", strings.Repeat(" ", 3)) + r
	}
	// Hours
	if d.Hours() >= 1 {
		hour := math.Mod(d.Hours(), 24.0)
		r = fmt.Sprintf("%2.0fh", hour) + r
	} else {
		r = fmt.Sprintf("%s", strings.Repeat(" ", 3)) + r
	}
	// Days
	days := d.Hours() / 24.0
	day := math.Mod(days, 7.0)
	if days >= 1 {
		r = fmt.Sprintf("%2.0fd", day) + r
	} else {
		r = fmt.Sprintf("%s", strings.Repeat(" ", 3)) + r
	}
	// Weeks
	week := int64(days / 7.0)
	if week >= 1 {
		r = fmt.Sprintf("%3dw", week) + r
	} else {
		r = fmt.Sprintf("%s", strings.Repeat(" ", 4)) + r
	}
	return r
}
func ack(e api.EventGetResult) string {
	return e.Event.AcknowledgedValue()
}
func host(e api.EventGetResult) string {
	if len(e.Hosts) > 0 {
		return e.Hosts[0].Name
	} else {
		return ""
	}
}
func name(e api.EventGetResult) string {
	if len(e.Alerts) > 0 {
		return e.Alerts[0].Subject
	} else {
		return ""
	}
}
func user(a api.Acknowledge) string {
	u := fmt.Sprintf("%s (%s %s)", a.Alias, a.Name, a.Surname)
	return u
}
func message(a api.Acknowledge) string {
	return a.Message
}
func userAction(a api.Acknowledge) string {
	switch a.Action {
	case 0:
		return ""
	case 1:
		return "Close problem"
	default:
		return "undefined"
	}
}

func (self *EventWidget) findEventidByPosition(p int) string {
	r := ""
	for i, e := range self.properties {
		if p == e.position {
			r = i
			break
		}
	}
	return r
}

func (self *EventWidget) findEventById(i string) (api.EventGetResult, error) {
	var q api.EventGetResult
	for _, e := range self.events {
		if e.Eventid == i {
			return e, nil
		}
	}
	return q, errors.Errorf("Corresponding Event is not found.(Event ID: %s)", i)
}
