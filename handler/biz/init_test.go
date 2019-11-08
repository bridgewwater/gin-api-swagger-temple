package biz

import (
	"os"
	"testing"
)

var baseURL string

func mockTestByEnv() {
	baseURL = "http://192.168.111.214:39000"
}

func TestMain(m *testing.M) {
	// setup
	mockTestByEnv()
	os.Exit(m.Run())
	// teardown
}
