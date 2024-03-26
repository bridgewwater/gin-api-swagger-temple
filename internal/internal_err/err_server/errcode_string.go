// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package err_server

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GoVersionError-20000]
	_ = x[ConfigFormatError-20101]
	_ = x[RedisConnectError-20201]
	_ = x[DBConnectError-20301]
	_ = x[DBInitializeError-20302]
}

const (
	_ErrCode_name_0 = "GoVersionError"
	_ErrCode_name_1 = "ConfigFormatError"
	_ErrCode_name_2 = "RedisConnectError"
	_ErrCode_name_3 = "DBConnectErrorDBInitializeError"
)

var (
	_ErrCode_index_3 = [...]uint8{0, 14, 31}
)

func (i ErrCode) String() string {
	switch {
	case i == 20000:
		return _ErrCode_name_0
	case i == 20101:
		return _ErrCode_name_1
	case i == 20201:
		return _ErrCode_name_2
	case 20301 <= i && i <= 20302:
		i -= 20301
		return _ErrCode_name_3[_ErrCode_index_3[i]:_ErrCode_index_3[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}