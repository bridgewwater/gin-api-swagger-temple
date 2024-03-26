package parse_http

type HeaderContent struct {
	ParseTime  string            `json:"parse_time,omitempty"`
	URL        string            `json:"url"`
	Method     string            `json:"method,omitempty"`
	RemoteAddr string            `json:"remote_addr,omitempty"`
	Host       string            `json:"host,omitempty"`
	Proto      string            `json:"proto,omitempty"`
	ProtoMajor int               `json:"proto_major,omitempty"`
	ProtoMinor int               `json:"proto_minor,omitempty"`
	Param      string            `json:"param,omitempty"`
	Referer    string            `json:"referer,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Header     string            `json:"header,omitempty"`
}
