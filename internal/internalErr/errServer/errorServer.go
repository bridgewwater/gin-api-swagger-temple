package errServer

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

//go:generate stringer -type ErrCode -linecomment
const (
	GoVersionError    ErrCode = 20000
	ConfigFormatError ErrCode = 20101

	RedisConnectError ErrCode = 20201

	DBConnectError    ErrCode = 20301
	DBInitializeError ErrCode = 20302
)

// ErrCode
// error code
type ErrCode int

type ServerError struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"msg"`
}

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

func Error(code ErrCode) error {
	return errors.Wrap(&ServerError{
		Code:    code,
		Message: code.String(),
	}, "")
}

func Errorf(code ErrCode, format string, arg ...any) error {
	var msg string
	if len(arg) > 1 {
		msg = format
	} else {
		msg = fmt.Sprintf(format, arg...)
	}
	return errors.Wrap(&ServerError{
		Code:    code,
		Message: fmt.Sprintf("%s | %s", code.String(), msg),
	}, msg)
}

func IsEqualCode(err error, target ErrCode) bool {
	var errParse *ServerError
	asErr := errors.As(err, &errParse)
	if !asErr {
		return false
	}
	return errParse.Code == target
}

func ErrJson(code ErrCode) string {
	err := Error(code)
	var errParse *ServerError
	asErr := errors.As(err, &errParse)
	if !asErr {
		return ""
	}
	return errParse.Json()
}

func ErrJsonf(code ErrCode, format string, arg ...any) string {
	err := Errorf(code, format, arg...)
	var errParse *ServerError
	asErr := errors.As(err, &errParse)
	if !asErr {
		return ""
	}
	return errParse.Json()
}
