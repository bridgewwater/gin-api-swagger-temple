package router

import (
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/config"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found.")
	})

	// monitor API
	monitorAPI(g)

	// api base path at config base_path
	basePath := viper.GetString("base_path")

	var env = viper.GetString("runmode")
	var envName = ""
	if env == "debug" || env == "test" {
		checkPingServer(config.BaseURL())
		// set swagger info
		swaggerInfo(config.BaseURL())
		swaggerGroup(envName, g)
	}

	bizApi(g, basePath)

	// TODO other router

	return g
}
