package errdef

import "fmt"

type ErrDef struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (err ErrDef) Error() string {
	return err.Msg
}

// Err represents an error only show in log, but not to client
type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  error  `json:"-"`
}

// can return errdef.Err use fmt.Errorf
//	errcode errdef.ErrDef in $project/pkg/errdef/ file errcode.go, you can add more!
//	message string error message format
//	args interface error message args
// use like
// 	errdef.NewSfp(errdef.ErrParams, "pl error, now %v", pl)
//
func NewSfp(errcode *ErrDef, message string, args ...interface{}) *Err {
	if message == "" {
		return &Err{
			Code: errcode.Code,
			Msg:  errcode.Msg,
		}
	} else {
		return &Err{
			Code: errcode.Code,
			Msg:  fmt.Sprintf(message, args),
			Err:  fmt.Errorf(message, args),
		}
	}
}

// new error for use
//	errcode errdef.ErrDef in $project/pkg/errdef/ file errcode.go, you can add more!
//	err error can use fmt.Errorf() to create
// use like
//	errdef.New(errdef.InternalServerError, fmt.Errorf("server error, err: %v", err)).Add("client can know error")
//
func New(errcode *ErrDef, err error) *Err {
	return &Err{Code: errcode.Code, Msg: errcode.Msg, Err: err}
}

// to errdef.New().add("user message")
func (err *Err) Add(message string) error {
	//err.Msg = fmt.Sprintf("%s %s", err.Msg, message)
	err.Msg += " " + message
	return err
}

// to errdef.New().addf("user message %v", args)
func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.Msg = fmt.Sprintf("%s %s", err.Msg, fmt.Sprintf(format, args...))
	err.Msg += " " + fmt.Sprintf(format, args...)
	return err
}

// errdef.New().Error() to print error message
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Msg, err.Err)
}

// asset error is ErrUserNotFound to use errdef.DecodeErr()
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

// decode error
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Msg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Msg
	case *ErrDef:
		return typed.Code, typed.Msg
	default:
	}

	return InternalServerError.Code, err.Error()
}
