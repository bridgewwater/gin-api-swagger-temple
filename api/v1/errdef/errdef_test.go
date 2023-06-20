package errdef

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCode(t *testing.T) {
	// mock NewCode

	t.Logf("~> mock NewCode")
	errCode := NewErr(ErrParams)
	_ = errCode.Add("params id not found")
	// do NewCode
	t.Logf("~> do NewCode")
	code := errCode.Code
	msg := errCode.Msg
	errorInfo := errCode.Error()
	// verify NewCode
	assert.Equal(t, 10003, code)
	assert.Equal(t, "Error params. Error params.. params id not found", msg)
	assert.Equal(t, "Err - code: 10003, message: Error params. Error params.. params id not found, error: Error params.", errorInfo)
}

func TestNew(t *testing.T) {
	// mock New

	t.Logf("~> mock New")
	paramsErr := fmt.Errorf("error params")
	err := New(ErrParams, paramsErr)
	// do New
	t.Logf("~> do New")
	code := err.Code
	errInfo := err.Error()
	// verify New
	assert.Equal(t, 10003, code)
	assert.Equal(t, "Err - code: 10003, message: Error params., error: error params", errInfo)
}

func TestErr_Add(t *testing.T) {
	// mock Err_Add

	t.Logf("~> mock Err_Add")
	paramsErr := fmt.Errorf("error params")
	addMessage := "add message"
	err := New(ErrParams, paramsErr)
	_ = err.Add(addMessage)
	// do Err_Add
	t.Logf("~> do Err_Add")
	code := err.Code
	msg := err.Msg
	// verify Err_Add
	assert.Equal(t, 10003, code)
	assert.Equal(t, "Error params. error params. add message", msg)
}

func TestIsErrUserNotFound(t *testing.T) {
	// mock IsErrUserNotFound

	t.Logf("~> mock IsErrUserNotFound")
	userNotFound := NewErr(ErrUserNotFound)
	bindErr := NewErr(ErrBind)

	// do IsErrUserNotFound
	foundOne := IsErrUserNotFound(userNotFound)
	foundTwo := IsErrUserNotFound(bindErr)
	t.Logf("~> do IsErrUserNotFound")
	// verify IsErrUserNotFound
	assert.True(t, foundOne)
	assert.False(t, foundTwo)
}
