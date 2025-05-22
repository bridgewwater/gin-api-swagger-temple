package biz_test

import (
	"fmt"
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/handler/biz"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/config"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog"
	"github.com/gin-gonic/gin"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	KeyEnvMockBaseUrl = "ENV_MOCK_BASE_URL"

	mockBasePath = "/api/v1"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	testGoldenKit      *unittest_file_kit.TestGoldenKit

	valEnvBaseURL string
	valBasePath   = mockBasePath
	basicRouter   *gin.Engine
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()
	testGoldenKit = unittest_file_kit.NewTestGoldenKit(testBaseFolderPath)

	zlog.MockZapLoggerInit()
	mockTestEnv()

	basicRouter = setupTestRouter()
}

func mockTestEnv() {
	valEnvBaseURL = env_kit.FetchOsEnvStr(KeyEnvMockBaseUrl, "http://0.0.0.0:34565")
}

func setupTestRouter() *gin.Engine {
	r := config.MockInitSample(valEnvBaseURL, valBasePath, false)
	biz.Router(r, valBasePath)
	return r
}

func TestMain(m *testing.M) {
	// setup
	mockTestEnv()
	os.Exit(m.Run())
	// teardown
}

// test case basic tools start
// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// test case basic tools end
