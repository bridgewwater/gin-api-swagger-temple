package err_server

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

	ServerErrorFunc ServerErrorFunc `json:"-"`
}

type ServerErrorFunc interface {
	Json() string
}
