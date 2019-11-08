package sys

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNetworkLocalIP(t *testing.T) {
	convey.Convey("TestNetworkLocalIP", t, func() {
		// mock
		// do
		ipv4, err := NetworkLocalIP()
		if err != nil {
			t.Fatalf("TestNetworkLocalIP test error: %v", err)
		}
		// verify
		t.Logf("NetworkLocalIP get: %v", ipv4)
		convey.So(ipv4, convey.ShouldNotBeNil)
	})
}
