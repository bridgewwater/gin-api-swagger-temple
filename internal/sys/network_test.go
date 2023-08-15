package sys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkLocalIP(t *testing.T) {
	// mock NetworkLocalIP

	t.Logf("~> mock NetworkLocalIP")
	// do NetworkLocalIP
	ipv4, err := NetworkLocalIP()
	if err != nil {
		t.Fatalf("TestNetworkLocalIP test error: %v", err)
	}
	t.Logf("~> do NetworkLocalIP")
	t.Logf("NetworkLocalIP get: %v", ipv4)
	// verify NetworkLocalIP
	assert.NotNil(t, ipv4)
}
