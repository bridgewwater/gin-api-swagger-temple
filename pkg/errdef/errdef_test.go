package errdef

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_err_def_NewCode(t *testing.T) {
	convey.Convey("mock Test_err_def_NewCode", t, func() {
		// mock
		errCode := NewErr(ErrParams)
		_ = errCode.Add("params id not found")

		convey.Convey("do Test_err_def_NewCode", func() {
			// do
			code := errCode.Code
			msg := errCode.Msg
			errorInfo := errCode.Error()
			convey.Convey("verify Test_err_def_NewCode", func() {
				// verify
				convey.So(code, convey.ShouldEqual, 10003)
				convey.So(msg, convey.ShouldEqual, "Error params. params id not found")
				convey.So(errorInfo, convey.ShouldEqual, "Err - code: 10003, message: Error params. params id not found, error: Error params.")
			})
		})
	})
}

func Test_err_def_NEW(t *testing.T) {
	convey.Convey("Test", t, func() {
		// mock
		// do
		paramsErr := fmt.Errorf("error params")
		err := New(ErrParams, paramsErr)
		code := err.Code
		errInfo := err.Error()
		// verify
		convey.So(code, convey.ShouldEqual, 10003)
		convey.So(errInfo, convey.ShouldEqual, "Err - code: 10003, message: Error params., error: error params")
	})
}

func Test_err_def_add(t *testing.T) {
	convey.Convey("mock Test_err_def_add", t, func() {
		// mock
		paramsErr := fmt.Errorf("error params")
		addMessage := "add message"
		err := New(ErrParams, paramsErr)
		_ = err.Add(addMessage)
		convey.Convey("do Test_err_def_add", func() {
			// do
			code := err.Code
			msg := err.Msg
			convey.Convey("verify Test_err_def_add", func() {
				// verify
				convey.So(code, convey.ShouldEqual, 10003)
				convey.So(msg, convey.ShouldEqual, "Error params. add message")
			})
		})
	})
}
func Test_err_def_IsErrUserNotFound(t *testing.T) {
	convey.Convey("mock Test_err_def_IsErrUserNotFound", t, func() {
		// mock
		userNotFound := NewErr(ErrUserNotFound)
		bindErr := NewErr(ErrBind)
		convey.Convey("do Test_err_def_IsErrUserNotFound", func() {
			// do
			foundOne := IsErrUserNotFound(userNotFound)
			foundTwo := IsErrUserNotFound(bindErr)
			convey.Convey("verify Test_err_def_IsErrUserNotFound", func() {
				// verify
				convey.So(foundOne, convey.ShouldEqual, true)
				convey.So(foundTwo, convey.ShouldNotEqual, true)
			})
		})
	})
}
