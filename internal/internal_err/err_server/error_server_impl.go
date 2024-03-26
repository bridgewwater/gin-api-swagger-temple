package err_server

import "encoding/json"

func (e *ServerError) Error() string {
	return e.Code.String()
}

func (e *ServerError) Json() string {
	marshal, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(marshal)
}
