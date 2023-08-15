package middleware

import (
	"github.com/bar-counter/gin-correlation-id/gin_correlation_id_snowflake"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Usage(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found.")
	})

	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(noCache)
	g.Use(options)
	g.Use(secure)

	g.Use(xAppVersionTracking())
	gin_correlation_id_snowflake.SetSnowflakeNameSpace(gin_correlation_id_snowflake.SnowflakeModeBase58)
	g.Use(gin_correlation_id_snowflake.Middleware())

	// monitor API
	monitorAPI(g)

	checkPingServer(config.BaseURL())
	g.Use(mw...)

	return g
}
