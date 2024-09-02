package errdef

import "fmt"

// Err represents an error only show in log, but not to client
type Err struct {
	// Code
	// error code default 0 is OK
	Code int `json:"code"   validate:"required"  example:"0"`
	// Msg
	// error message, if OK is empty
	Msg string `json:"msg"  validate:"optional"  example:"msg of error code 0 is empty"`
	Err error  `json:"-"`
}

// NewErr
// new error for use
// e errdef.ErrDef in $project/pkg/errdef/ file e.go, you can add more!
// use like
// errdef.NewErr(errdef.ErrParams)
// and you can add message
//
//	errdef.NewErr(errdef.ErrParams).Add("params id not found")
func NewErr(e *ErrDef) *Err {
	// nolint:govet
	return &Err{Code: e.Code, Msg: e.Msg, Err: fmt.Errorf(e.Msg)}
}

// New
// new error for use
// e errdef.ErrDef in $project/pkg/errdef/ file e.go, you can add more!
// err error can use fmt.Errorf() to create
// use like
// errdef.New(errdef.InternalServerError, fmt.Errorf("server error, err: %v", err))
// and you can add message
// errdef.New(errdef.InternalServerError, fmt.Errorf("server error, err: %v", err)).Add("client can know error")
func New(e *ErrDef, err error) *Err {
	errDef := Err{Code: e.Code, Msg: e.Msg, Err: err}
	errDef.Msg = fmt.Sprintf("%s %s", e.Msg, err.Error())
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
