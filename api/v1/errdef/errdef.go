package errdef

import "fmt"

type ErrDef struct {
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	HttpStatus int    `json:"-"`
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

// NewErr
// new error for use
// errcode errdef.ErrDef in $project/pkg/errdef/ file errcode.go, you can add more!
// use like
// errdef.NewErr(errdef.ErrParams)
// and you can add message
//
//	errdef.NewErr(errdef.ErrParams).Add("params id not found")
func NewErr(errcode *ErrDef) *Err {
	return &Err{Code: errcode.Code, Msg: errcode.Msg, Err: fmt.Errorf(errcode.Msg)}
}

// New
// new error for use
// errcode errdef.ErrDef in $project/pkg/errdef/ file errcode.go, you can add more!
// err error can use fmt.Errorf() to create
// use like
// errdef.New(errdef.InternalServerError, fmt.Errorf("server error, err: %v", err))
// and you can add message
// errdef.New(errdef.InternalServerError, fmt.Errorf("server error, err: %v", err)).Add("client can know error")
func New(errcode *ErrDef, err error) *Err {
	errDef := Err{Code: errcode.Code, Msg: errcode.Msg, Err: err}
	errDef.Msg = fmt.Sprintf("%s %s", errcode.Msg, err.Error())
	return &errDef
}

// Add
// to errdef.New().add("user message")
func (err *Err) Add(message string) error {
	//err.Msg = fmt.Sprintf("%s %s", err.Msg, message)
	//err.Msg += " " + message
	err.Msg = fmt.Sprintf("%v %v. %v", err.Msg, err.Err.Error(), message)
	return err
}

// Addf
// to errdef.New().addf("user message %v", args)
func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.Msg = fmt.Sprintf("%s %s", err.Msg, fmt.Sprintf(format, args...))
	err.Msg += " " + fmt.Sprintf(format, args...)
	return err
}

// Error
// errdef.New().Error() to print error message
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Msg, err.Err)
}

// DecodeErr
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

// IsErrUserNotFound
// asset error is ErrUserNotFound to use errdef.DecodeErr()
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}
