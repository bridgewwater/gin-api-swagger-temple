package middleware

import (
	"github.com/bridgewwater/gin-api-swagger-temple/pkg/pkgJson"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	headKeyApiVersion = "X-App-Version"
)

// AppVersion
//
//	usage
//	middleware.AppVersion(c)
//	return now api app version
func AppVersion(c *gin.Context) string {
	return c.Request.Header.Get(headKeyApiVersion)
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
	if strings.TrimSpace(apiVersion) == "" {
		packageJsonVersion := pkgJson.GetPackageJsonVersion()
		c.Request.Header.Add(headKeyApiVersion, packageJsonVersion)
		c.Header(headKeyApiVersion, packageJsonVersion)
	}
	c.Next()
}
