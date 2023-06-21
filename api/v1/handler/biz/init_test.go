package biz

import (
	"os"
	"testing"
)

var baseURL string

func mockTestByEnv() {
	baseURL = "http://127.0.0.1:39000"
}

func TestMain(m *testing.M) {
	// setup
	mockTestByEnv()
	os.Exit(m.Run())
	// teardown
}