package v1

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title gin-api-swagger-temple
// @version v1.x.x
// @description This is a sample server
// @termsOfService http://github.com/

// @contact.name API Support
// @contact.url http://github.com/
// @contact.email support@sinlov.cn

// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey		WithToken
// @in								header
// @name							Authorization
// @description					Please set the token of the API, note that it starts with "Bearer"

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func Register(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(mw...)
	// api base path at config base_path
	basePath := viper.GetString("base_path")

	var env = config.GinRunMode()
	if env == "debug" || env == "test" {
		// set swagger info
		swaggerInfo(config.BaseURL())
		swaggerGroup(config.BaseURL(), g)
	}

	bizApi(g, basePath)

	// TODO: other router bind here

	return g
}
