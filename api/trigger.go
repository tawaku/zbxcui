package api

type Trigger struct {
	Triggerid          string `json:"triggerid"`
	Descripttion       string `json:"description"`
	Expression         string `json:"expression"`
	Comments           string `json:"comments"`
	Error              string `json:"error"`
	Flags              byte   `json:"flags,string"`
	Lastchange         int64  `json:"lastchange,string"`
	Priority           byte   `json:"priority,string"`
	State              byte   `json:"state,string"`
	Status             byte   `json:"status,string"`
	Templateid         string `json:"templateid"`
	Type               byte   `json:"type,string"`
	Url                string `json:"url"`
	Value              byte   `json:"value,string"`
	RecoveryMode       byte   `json:"recovery_mode,string"`
	RecoveryExpression string `json:"recovery_expression"`
	CorrelationMode    byte   `json:"correlation_mode,string"`
	CorrelationTag     string `json:"correlation_tag"`
	ManualClose        byte   `json:"manual_close,string"`
}

func (self *Trigger) FlagsValue() string {
	switch self.Flags {
	case 0:
		return "a plain trigger"
	case 4:
		return "a discoered trigger"
	default:
		return "undefined"
	}
}

func (self *Trigger) PriorityValue() string {
	switch self.Priority {
	case 0:
		return "not classified"
	case 1:
		return "information"
	case 2:
		return "warning"
	case 3:
		return "average"
	case 4:
		return "high"
	case 5:
		return "disaster"
	default:
		return "undefined"
	}
}

func (self *Trigger) StateValue() string {
	switch self.State {
	case 0:
		return "trigger state is up to date"
	case 1:
		return "current trigger state is unknown"
	default:
		return "undefined"
	}
}

func (self *Trigger) StatusValue() string {
	switch self.Status {
	case 0:
		return "enabled"
	case 1:
		return "disabled"
	default:
		return "undefined"
	}
}

func (self *Trigger) TypeValue() string {
	switch self.Type {
	case 0:
		return "do not generate multiple events"
	case 1:
		return "generate multiple events"
	default:
		return "undefined"
	}
}

func (self *Trigger) ValueValue() string {
	switch self.Value {
	case 0:
		return "OK"
	case 1:
		return "problem"
	default:
		return "undefined"
	}
}

func (self *Trigger) RecoveryModeValue() string {
	switch self.RecoveryMode {
	case 0:
		return "Expression"
	case 1:
		return "Recovery expression"
	case 2:
		return "None"
	default:
		return "undefined"
	}
}

func (self *Trigger) CorrelationModeValue() string {
	switch self.CorrelationMode {
	case 0:
		return "All problems"
	case 1:
		return "All problems if tag values match"
	default:
		return "undefined"
	}
}

func (self *Trigger) ManualCloseValue() string {
	switch self.ManualClose {
	case 0:
		return "No"
	case 1:
		return "Yes"
	default:
		return "undefined"
	}
}
