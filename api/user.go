package api

type UserLogin struct {
	Password string `json:"password"`
	User     string `json:"user"`
	Userdata *bool  `json:"userData,omitempty"`
}

type UserLogout struct{}

type User struct {
	Userid        string `json:"userid"`
	Alias         string `json:"alias "`
	AttemptClock  int64  `json:"attempt_clock,string"`
	AttemptFailed uint32 `json:"attempt_failed,string"`
	AttemptIp     string `json:"attempt_ip"`
	Autologin     byte   `json:"autologin,string"`
	Autologout    string `json:"autologout"`
	Lang          string `json:"lang"`
	Name          string `json:"name"`
	Refresh       string `json:"refresh"`
	RowsPerPage   uint32 `json:"rows_per_page,string"`
	Surname       string `json:"surname"`
	Theme         string `json:"theme"`
	Type          string `json:"type"`
	Url           string `json:"url"`
	DebugMode     bool   `json:"debug_mode"`
	GuiAccess     byte   `json:"gui_access"`
	Sessionid     string `json:"sessionid"`
	Userip        string `json:"userip"`
}

func (self *Client) userLogin(usr, pwd string, dat bool) *Response {
	// Json data to send
	send := self.makeRequest(
		"user.login",
		UserLogin{
			Password: pwd,
			User:     usr,
			Userdata: &dat,
		},
	)
	// Prepare struct to receive
	recv := new(Response)
	// Do request
	self.request(send, recv)
	// Return based on Error exist
	return recv
}

func (self *Client) userLogout() *Response {
	// Json data to send
	send := self.makeRequest(
		"user.logout",
		UserLogout{},
	)
	// Prepare struct to receive
	recv := new(Response)
	// Do request
	self.request(send, recv)
	// Return based on Error exist
	return recv
}
