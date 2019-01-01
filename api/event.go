package api

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type EventGet struct {
	*CommonGet
	Eventids            []string   `json:"eventids,omitempty"`
	Groupids            []string   `json:"groupids,omitempty"`
	Hostids             []string   `json:"hostids,omitempty"`
	Objectids           []string   `json:"objectids,omitempty"`
	Applicationids      []string   `json:"applicationids,omitempty"`
	Source              uint32     `json:"source,omitempty"`
	Object              uint32     `json:"object,omitempty"`
	Acknowledged        *bool      `json:"acknowledged,omitempty"`
	Severities          []byte     `json:"severities,omitempty"`
	Tags                []EventTag `json:"tags,omitempty"`
	EventidFrom         string     `json:"eventid_from,omitempty"`
	EventidTill         string     `json:"eventid_till,omitempty"`
	TimeFrom            int64      `json:"time_from,omitempty"`
	TimeTill            int64      `json:"time_till,omitempty"`
	Value               []uint32   `json:"value,omitempty"`
	SelectHosts         string     `json:"selectHosts,omitempty"`
	SelectRelatedObject string     `json:"selectRelatedObject,omitempty"`
	SelectAlerts        string     `json:"select_alerts,omitempty"`
	SelectAcknowledges  string     `json:"select_acknowledges,omitempty"`
	SelectTags          string     `json:"selectTags,omitempty"`
	Sortfield           []string   `json:"sortfield,omitempty"`
}

type EventGetResult struct {
	*Event
	Hosts         []Host        `json:"hosts,omitempty"`
	RelatedObject interface{}   `json:"relatedObject,omitempty"`
	Alerts        []Alert       `json:"alerts,omitempty"`
	Acknowledges  []Acknowledge `json:"acknowledges,omitempty"`
	Tags          []EventTag    `json:"tags,omitempty"`
}

type Event struct {
	Eventid       string `json:"eventid,omitempty"`
	Source        byte   `json:"source,string,omitempty"`
	Object        byte   `json:"object,string,omitempty"`
	Objectid      string `json:"objectid,omitempty"`
	Acknowledged  byte   `json:"acknowledged,string,omitempty"`
	Clock         int64  `json:"clock,string,omitempty"`
	Ns            uint32 `json:"ns,string,omitempty"`
	Value         byte   `json:"value,string,omitempty"`
	REventid      string `json:"r_eventid,omitempty"`
	CEventid      string `json:"c_eventid,omitempty"`
	Correlationid string `json:"correlationid,omitempty"`
	Userid        string `json:"userid,omitempty"`
}

type EventTag struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

type Acknowledge struct {
	Acknowledgeid string `json:"acknowledgeid,omitempty"`
	Userid        string `json:"userid,omitempty"`
	Eventid       string `json:"eventid,omitempty"`
	Clock         int64  `json:"clock,string,omitempty"`
	Message       string `json:"message,omitempty"`
	Action        byte   `json:"action,string,omitempty"`
	Alias         string `json:"alias,omitempty"`
	Name          string `json:"name,omitempty"`
	Surname       string `json:"surname,omitempty"`
}

func (self *Client) EventGet() ([]EventGetResult, error) {
	// Json data to send
	send := self.makeRequest(
		"event.get",
		EventGet{
			CommonGet: &CommonGet{
				Sortorder: []string{"DESC"},
			},
			SelectHosts:         "extend",
			SelectRelatedObject: "extend",
			SelectAlerts:        "extend",
			SelectAcknowledges:  "extend",
			SelectTags:          "extend",
			Sortfield:           []string{"clock"},
		},
	)
	// Prepare struct to receive
	recv := new(Response)
	// Do request
	self.request(send, recv)
	// Return
	if recv.Error != nil {
		return nil, recv.Error.Error()
	} else {
		if recv.Result == nil {
			return nil, errors.New("event.get API failed: result is not returned.")
		} else {
			if j, err := json.Marshal(recv.Result); err != nil {
				return nil, errors.Wrap(err, "event.get API failed: json.Marshal failed.")
			} else {
				r := []EventGetResult{}
				json.Unmarshal(j, &r)
				return r, nil
			}
		}
	}
}

func (self *Event) SourceValue() string {
	switch self.Source {
	case 0:
		return "event created by a trigger"
	case 1:
		return "event created by a discovery rule"
	case 2:
		return "event created by active agent auto-registration"
	case 3:
		return "internal event"
	default:
		return "undefined"
	}
}

func (self *Event) ObjectValue() string {
	switch self.Source {
	case 0:
		switch self.Object {
		case 0:
			return "trigger"
		default:
			return "undefined"
		}
	case 1:
		switch self.Object {
		case 1:
			return "discovered host"
		case 2:
			return "discovered service"
		default:
			return "undefined"
		}
	case 2:
		switch self.Object {
		case 3:
			return "auto-registered host"
		default:
			return "undefined"
		}
	case 3:
		switch self.Object {
		case 0:
			return "trigger"
		case 4:
			return "item"
		case 5:
			return "LLD rule"
		default:
			return "undefined"
		}
	default:
		return "undefined"
	}
}

func (self *Event) AcknowledgedValue() string {
	switch self.Acknowledged {
	case 0:
		return "no"
	case 1:
		return "yes"
	default:
		return "undefined"
	}
}

func (self *Event) ValueValue() string {
	switch self.Source {
	case 0:
		switch self.Value {
		case 0:
			return "OK"
		case 1:
			return "problem"
		default:
			return "undefined"
		}
	case 1:
		switch self.Value {
		case 0:
			return "host or service up"
		case 1:
			return "host or service down"
		case 2:
			return "host or service discovered"
		case 3:
			return "host or service lost"
		default:
			return "undefined"
		}
	case 3:
		switch self.Value {
		case 0:
			return "\"normal\" state"
		case 1:
			return "\"unknown\" or \"not supported\" state"
		default:
			return "undefined"
		}
	default:
		return "undefined"
	}
}

func (self *EventGetResult) RelatedObjectValue() interface{} {
	if j, err := json.Marshal(self.RelatedObject); err != nil {
		return nil
	} else {
		switch self.Event.Source {
		case 0:
			r := Trigger{}
			json.Unmarshal(j, &r)
			return r
		case 1:
			return nil
		case 2:
			return nil
		case 3:
			return nil
		default:
			return nil
		}
	}
}
