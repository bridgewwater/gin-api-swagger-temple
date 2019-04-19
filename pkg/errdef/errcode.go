package errdef

var (
	// Common errors
	OK                  = &ErrDef{Code: 0, Msg: "OK"}
	InternalServerError = &ErrDef{Code: 10001, Msg: "Internal server error."}
	ErrBind             = &ErrDef{Code: 10002, Msg: "Error occurred while binding the request body to the struct."}
	ErrParams           = &ErrDef{Code: 10003, Msg: "Error params."}

	// database errors
	ErrValidation = &ErrDef{Code: 20001, Msg: "Validation failed."}
	ErrDatabase   = &ErrDef{Code: 20002, Msg: "Database error."}
	ErrToken      = &ErrDef{Code: 20003, Msg: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &ErrDef{Code: 20101, Msg: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &ErrDef{Code: 20102, Msg: "The user was not found."}
	ErrTokenInvalid      = &ErrDef{Code: 20103, Msg: "The token was invalid."}
	ErrPasswordIncorrect = &ErrDef{Code: 20104, Msg: "The password was incorrect."}

	// other errors
)
