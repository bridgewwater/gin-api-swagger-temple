package errServer

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErr(t *testing.T) {
	// mock Error
	tests := []struct {
		name     string
		errCode  ErrCode
		wantJson string
		wantErr  error
	}{
		{
			name:     "Not Found",
			errCode:  10,
			wantJson: `{"code":10,"msg":"ErrCode(10)"}`,
		},
		{
			name:     "GoVersionError",
			errCode:  GoVersionError,
			wantJson: `{"code":20000,"msg":"GoVersionError"}`,
		},
		{
			name:     "ConfigFormatError",
			errCode:  ConfigFormatError,
			wantJson: `{"code":20101,"msg":"ConfigFormatError"}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do Error
			gotResult := Error(tc.errCode)

			// verify Error
			assert.NotNil(t, gotResult)
			var errParse *ServerError
			asErr := errors.As(gotResult, &errParse)
			assert.True(t, asErr)
			assert.Equal(t, tc.errCode, errParse.Code)
			assert.Equal(t, tc.wantJson, errParse.Json())
		})
	}
}

func TestNewErrf(t *testing.T) {
	// mock Errorf
	tests := []struct {
		name      string
		errCode   ErrCode
		errFormat string
		errArg    []any
		wantJson  string
		wantErr   error
	}{
		{
			name:      "Not Found",
			errCode:   10,
			errFormat: "not found",
			errArg:    nil,
			wantJson:  `{"code":10,"msg":"ErrCode(10) | not found"}`,
		},
		{
			name:      "GoVersionError",
			errCode:   GoVersionError,
			errFormat: "can not use this version %s",
			errArg:    []any{"1.0.0"},
			wantJson:  `{"code":20000,"msg":"GoVersionError | can not use this version 1.0.0"}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do Error
			gotResult := Errorf(tc.errCode, tc.errFormat, tc.errArg...)

			// verify Error
			assert.NotNil(t, gotResult)
			var errParse *ServerError
			asErr := errors.As(gotResult, &errParse)
			assert.True(t, asErr)
			assert.Equal(t, tc.errCode, errParse.Code)
			assert.Equal(t, tc.wantJson, errParse.Json())
		})
	}
}

func TestNewErrJson(t *testing.T) {
	// mock ErrJson
	tests := []struct {
		name     string
		errCode  ErrCode
		wantJson string
		wantErr  error
	}{
		{
			name:     "Not Found",
			errCode:  10,
			wantJson: `{"code":10,"msg":"ErrCode(10)"}`,
		},
		{
			name:     "GoVersionError",
			errCode:  GoVersionError,
			wantJson: `{"code":20000,"msg":"GoVersionError"}`,
		},
		{
			name:     "ConfigFormatError",
			errCode:  ConfigFormatError,
			wantJson: `{"code":20101,"msg":"ConfigFormatError"}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do ErrJson
			gotResult := ErrJson(tc.errCode)

			// verify ErrJson
			assert.Equal(t, tc.wantJson, gotResult)
		})
	}
}

func TestNewErrJsonf(t *testing.T) {
	// mock ErrJsonf
	tests := []struct {
		name      string
		errCode   ErrCode
		errFormat string
		errArg    []any
		wantJson  string
		wantErr   error
	}{
		{
			name:      "Not Found",
			errCode:   10,
			errFormat: "not found",
			errArg:    nil,
			wantJson:  `{"code":10,"msg":"ErrCode(10) | not found"}`,
		},
		{
			name:      "GoVersionError",
			errCode:   GoVersionError,
			errFormat: "can not use this version %s",
			errArg:    []any{"1.0.0"},
			wantJson:  `{"code":20000,"msg":"GoVersionError | can not use this version 1.0.0"}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do ErrJsonf
			gotResult := ErrJsonf(tc.errCode, tc.errFormat, tc.errArg...)

			// verify ErrJsonf
			assert.Equal(t, tc.wantJson, gotResult)
		})
	}
}

func TestIsEqualCode(t *testing.T) {
	// mock IsEqualCode
	tests := []struct {
		name    string
		err     error
		code    ErrCode
		wantRes bool
	}{
		{
			name:    "foo",
			err:     errors.New("foo"),
			code:    GoVersionError,
			wantRes: false,
		},
		{
			name:    "GoVersionError",
			err:     Error(GoVersionError),
			code:    GoVersionError,
			wantRes: true,
		},
		{
			name:    "ConfigFormatError",
			err:     Errorf(ConfigFormatError, "some config error"),
			code:    ConfigFormatError,
			wantRes: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do IsEqualCode
			gotResult := IsEqualCode(tc.err, tc.code)

			// verify IsEqualCode
			assert.Equal(t, tc.wantRes, gotResult)
		})
	}
}
