package err_server

import (
	"fmt"
	"github.com/pkg/errors"
)

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
