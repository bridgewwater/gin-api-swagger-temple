package middleware

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/pkg/pkg_kit"
	"github.com/bridgewwater/gin-api-swagger-temple/zymosis"
	"github.com/gin-gonic/gin"
	"strings"
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
	var apiVersion = AppVersion(c)
	if len(strings.TrimSpace(apiVersion)) == 0 {
		packageJsonVersion := pkg_kit.GetPackageJsonVersion()
		c.Request.Header.Add(headKeyApiVersion, packageJsonVersion)
		c.Header(headKeyApiVersion, packageJsonVersion)

	}
	var appMainRes = AppMainRes(c)
	if len(strings.TrimSpace(appMainRes)) == 0 {
		mainProgramRes := zymosis.MainProgramRes()
		c.Request.Header.Add(headKeyMainRes, mainProgramRes)
		c.Header(headKeyMainRes, mainProgramRes)
	}
	c.Next()
}
