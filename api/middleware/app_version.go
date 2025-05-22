package middleware

import (
	"strings"

	"github.com/bridgewwater/gin-api-swagger-temple/internal/pkg/pkg_kit"
	"github.com/gin-gonic/gin"
)

const (
	headKeyApiVersion = "X-App-Version"
	headKeyMainRes    = "X-App-Main-Res"
)

// AppVersion
//
//	usage
//	middleware.AppVersion(c)
//	return now api app version
func AppVersion(c *gin.Context) string {
	return c.Request.Header.Get(headKeyApiVersion)
}

func AppMainRes(c *gin.Context) string {
	return c.Request.Header.Get(headKeyMainRes)
}

// xAppVersionTracking
//
//	usage
//	g.Use(xAppVersionTracking())
func xAppVersionTracking() gin.HandlerFunc {
	return appVersionTracking
}

func appVersionTracking(c *gin.Context) {
	apiVersion := AppVersion(c)
	if len(strings.TrimSpace(apiVersion)) == 0 {
		packageJsonVersion := pkg_kit.FetchNowVersion()
		c.Request.Header.Add(headKeyApiVersion, packageJsonVersion)
		c.Header(headKeyApiVersion, packageJsonVersion)
	}

	appMainRes := AppMainRes(c)
	if len(strings.TrimSpace(appMainRes)) == 0 {
		mainProgramRes := pkg_kit.FetchNowBuildCode()
		if mainProgramRes == "" {
			mainProgramRes = pkg_kit.InfoUnknown
		}
		c.Request.Header.Add(headKeyMainRes, mainProgramRes)
		c.Header(headKeyMainRes, mainProgramRes)
	}

	c.Next()
}
