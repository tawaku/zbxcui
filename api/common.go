package api

type Request struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    *string     `json:"auth"`
	Id      uint32      `json:"id"`
}

type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      uint32      `json:"id,string"`
	Result  interface{} `json:"result,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

type CommonGet struct {
	CountOutput            *bool       `json:"countOutput,omitempty"`
	Editable               *bool       `json:"editable,omitempty"`
	ExcludeSearch          *bool       `json:"excludeSearch,omitempty"`
	Filter                 interface{} `json:"filter,omitempty"`
	Limit                  uint32      `json:"limit,omitempty"`
	Nodeids                string      `json:"nodeids,omitempty"`
	Output                 []string    `json:"output,omitempty"`
	Preservekeys           *bool       `json:"preservekeys,omitempty"`
	Search                 interface{} `json:"search,omitempty"`
	SearchByAny            *bool       `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled *bool       `json:"searchWildcardsEnabled,omitempty"`
	Sortfield              []string    `json:"sortfield,omitempty"`
	Sortorder              []string    `json:"sortorder,omitempty"`
	StartSearch            *bool       `json:"startSearch,omitempty"`
}

func MakeSearch(k string, v string) interface{} {
	s := map[string]interface{}{k: v}
	return s
}

func MakeFilter(k string, v []string) interface{} {
	f := map[string]interface{}{k: v}
	return f
}
