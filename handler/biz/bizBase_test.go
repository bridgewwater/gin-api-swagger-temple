package biz

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetString(t *testing.T) {
	convey.Convey("TestGetString", t, func() {
		// mock
		url := fmt.Sprintf("%v%v%v", baseURL, "/v1/biz", "/string")
		request := gorequest.New()
		// do
		resp, body, errs := request.Get(url).
			End()
		// verify
		if errs != nil {
			t.Fatalf("test error url: %v, why: %v", url, errs)
		}
		if resp.StatusCode != 200 {
			t.Errorf("url: %v fail, StatusCode %v, body: %v", url, resp.StatusCode, body)
		}
		t.Logf("url: %v ; resp body: %v", url, body)
		convey.So(body, convey.ShouldEqual, "this is biz message")
	})
}
