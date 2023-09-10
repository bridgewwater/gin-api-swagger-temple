package zlog_access

import (
	"fmt"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/common"
	"go.uber.org/zap"
	"sort"
	"strings"
)

// A
// for access zap Sugar
func A() *zap.SugaredLogger {
	if zapLog == nil {
		panic(fmt.Errorf("please use zlog_access.InitByViper()"))
	}
	return zapLog.Sugar
}

var zapLog *zapLogger

type zapLogger struct {
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
}

// newAccessAsZap
// init zap by sugar
func newAccessAsZap(log *zap.Logger, sugar *zap.SugaredLogger) {
	// format skipPath
	skipPath = common.RemoveStringDuplicateNotCopy(skipPath)
	if zapLog == nil {
		zapLog = &zapLogger{
			Log:   log,
			Sugar: sugar,
		}
		sugar.Info("zap log init access logger as sugar now")
	}
}

var (
	skipPath = []string{
		"/favicon.ico",
		"/ping",
		"/status/health",
		"/status/hardware/disk",
		"/status/hardware/ram",
		"/status/hardware/cpu",
		"/status/hardware/cpu_info",
		"/debug/vars",
		"/debug/pprof/",
		"/swagger/editor/index.html",
		"/swagger/editor/swagger-ui.css",
		"/swagger/editor/swagger-ui-standalone-preset.js",
		"/swagger/editor/swagger-ui-bundle.js",
		"/swagger/v1_doc.json",
	}
)

func CheckPathIsSkip(target string) bool {
	index := sort.SearchStrings(skipPath, target)
	// Attention should be paid to the judgment here.
	// First, judge the condition on the left side of & &.
	//If it is not met, end the judgment here and will not be judged on the right side.
	if index < len(skipPath) && skipPath[index] == target {
		return true
	}
	return false

}

func AppendSkipPath(listPath ...string) {
	if len(listPath) == 0 {
		return
	}
	skipPath = append(skipPath, listPath...)
	skipPath = common.RemoveStringDuplicateNotCopy(skipPath)
}

var (
	skipPrefix []string
)

func CheckPrefixIsSkip(target string) bool {
	if len(skipPrefix) == 0 {
		return false
	}

	for _, prefix := range skipPrefix {
		if strings.HasPrefix(target, prefix) {
			return true
		}
	}
	return false
}

// AppendSkipPrefix
//
//	not dot append too much item
func AppendSkipPrefix(listPrefix ...string) {
	if len(listPrefix) == 0 {
		return
	}
	skipPrefix = append(skipPrefix, listPrefix...)
	skipPrefix = common.RemoveStringDuplicateNotCopy(skipPrefix)
}
