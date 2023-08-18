package biz_test

import (
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/handler/biz"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog"
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

const (
	EnvBaseUrl = "ENV_MOCK_BASE_URL"
)

var (
	baseURL     string
	basePath    = "/api/v1"
	basicRouter *gin.Engine
)

func init() {
	zlog.MockZapLoggerInit()
	mockTestEnv()
	basicRouter = setupTestRouter()
}

func setupTestRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(FetchGinRunMode())
	biz.Router(r, basePath)
	return r
}

func mockTestEnv() {
	baseURL = fetchOsEnvStr(EnvBaseUrl, "http://127.0.0.1:34565")
}

func TestMain(m *testing.M) {
	// setup
	mockTestEnv()
	os.Exit(m.Run())
	// teardown
}
