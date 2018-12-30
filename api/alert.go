package api

type AlertGet struct {
	*CommonGet
	Alertids         []string `json:"alertids,omitempty"`
	Actionids        []string `json:"actionids,omitempty"`
	Eventids         []string `json:"eventids,omitempty"`
	Groupids         []string `json:"groupids,omitempty"`
	Hostids          []string `json:"hostids,omitempty"`
	Mediatypeids     []string `json:"mediatypeids,omitempty"`
	Objectids        []string `json:"objectids,omitempty"`
	Userids          []string `json:"userids,omitempty"`
	EventObject      byte     `json:"eventobject,omitempty"`
	EventSource      byte     `json:"eventsource,omitempty"`
	TimeFrom         int64    `json:"time_from,omitempty"`
	TimeTill         int64    `json:"time_till,omitempty"`
	SelectHosts      []string `json:"selectHosts,omitempty"`
	SelectMediatypes []string `json:"selectMediatypes,omitempty"`
	SelectUsers      []string `json:"selectUsers,omitempty"`
	Sortfield        []string `json:"sortfield,omitempty"`
}

type Alert struct {
	Alertid      string `json:"alertid"`
	Actionid     string `json:"actionid"`
	AlertType    byte   `json:"alerttype,string"`
	Clock        int64  `json:"clock,string"`
	Error        string `json:"error"`
	EscStep      uint32 `json:"esc_step,string"`
	Eventid      string `json:"eventid"`
	Mediatypeid  string `json:"mediatypeid"`
	Message      string `json:"message"`
	Retries      uint32 `json:"retries,string"`
	Sendto       string `json:"sendto"`
	Status       byte   `json:"status,string"`
	Subject      string `json:"subject"`
	Userid       string `json:"userid"`
	PEventid     string `json:"p_eventid"`
	Acknowledged string `json:"acknowledged"`
}

func (self *Alert) AlertTypeValue() string {
	switch self.AlertType {
	case 0:
		return "message"
	case 1:
		return "remote command"
	default:
		return "undefined"
	}
}

func (self *Alert) StatusValue() string {
	switch self.AlertType {
	case 0:
		switch self.Status {
		case 0:
			return "message not sent."
		case 1:
			return "message sent."
		case 2:
			return "faild after a number of retries."
		default:
			return "undefined"
		}
	case 1:
		switch self.Status {
		case 0:
			return "command not run."
		case 1:
			return "command run."
		case 2:
			return "tried to run the command on the Zabbix agent but it was unavailable."
		default:
			return "undefined"
		}
	default:
		return "undefined"
	}
}
