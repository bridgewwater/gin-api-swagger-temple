package errdef

type ErrDef struct {
	// Code
	// error code default 0 is OK
	Code int `json:"code"   validate:"required"  example:"0"`
	// Msg
	// error message, if OK is empty
	Msg string `json:"msg"  validate:"optional"  example:"msg of error code 0 is empty"`
	// HttpStatus
	// http status code, this is set by config
	HttpStatus int `json:"-"`
}

func (err *ErrDef) Error() string {
	return err.Msg
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

// IsEqualByCode
// asset error is code to use errdef.DecodeErr()
func IsEqualByCode(err error, code int) bool {
	codeErr, _ := DecodeErr(err)
	return codeErr == code
}

// IsErrUserNotFound
// asset error is ErrUserNotFound to use errdef.DecodeErr()
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}
