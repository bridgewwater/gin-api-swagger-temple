package parse_http

type FormContent struct {
	URL          string            `json:"url"`
	ParseTime    string            `json:"parse_time,omitempty"`
	Method       string            `json:"method,omitempty"`
	RemoteAddr   string            `json:"remote_addr,omitempty"`
	Host         string            `json:"host,omitempty"`
	Proto        string            `json:"proto,omitempty"`
	ProtoMajor   int               `json:"proto_major,omitempty"`
	ProtoMinor   int               `json:"proto_minor,omitempty"`
	Param        string            `json:"param,omitempty"`
	Referer      string            `json:"referer,omitempty"`
	Headers      map[string]string `json:"headers,omitempty"`
	PostFormData map[string]string `json:"postData,omitempty"`
	QueryStrings map[string]string `json:"queryString,omitempty"`
}
