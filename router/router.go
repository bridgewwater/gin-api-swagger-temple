package router

import (
	"fmt"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/config"
	_ "git.sinlov.cn/bridgewwater/temp-gin-api-self/docs" // docs is generated by Swag CLI, you have to import it. use [ swag init ]
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/handler/ssc"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// api base path at config base_path
	basePath := viper.GetString("base_path")

	// swagger api docs
	var env = viper.GetString("runmode")
	var envName = ""
	if env == "debug" || env == "test" {
		swaggerGroup(envName, g)
	}

	// The health check handlers
	sscRouteGroup := g.Group(basePath + "/ssc")
	{
		sscRouteGroup.GET("/health", ssc.HealthCheck)
		sscRouteGroup.GET("/disk", ssc.DiskCheck)
		sscRouteGroup.GET("/cpu", ssc.CPUCheck)
		sscRouteGroup.GET("/ram", ssc.RAMCheck)
	}

	// TODO other router

	return g
}

func swaggerGroup(envName string, g *gin.Engine) {
	envName = "NAME_OF_ENV_VARIABLE"
	log.Infof("envName %v", envName)
	// https://github.com/swaggo/swag/issues/194#issuecomment-475853710
	configSwagger := &ginSwagger.Config{
		URL: fmt.Sprintf("%v%v", config.BaseURL(), "/swagger/doc.json"), //The url pointing to API definition
	}
	swaggerRoot := viper.GetString("swagger.root")
	if viper.GetBool("swagger.security") {
		// just use swagger as swagger_user
		swaggerGroup := g.Group(swaggerRoot, gin.BasicAuth(gin.Accounts{
			"admin": viper.GetString("swagger.user.admin"),
			"user":  viper.GetString("swagger.user.user"),
		}))
		//noinspection GoTypesCompatibility
		swaggerGroup.GET("/*any", ginSwagger.CustomWrapHandler(configSwagger, swaggerFiles.Handler))
	} else {
		//noinspection GoTypesCompatibility
		//g.GET(swaggerRoot+"/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, envName))
		g.GET(swaggerRoot+"/*any", ginSwagger.CustomWrapHandler(configSwagger, swaggerFiles.Handler))
	}
}
