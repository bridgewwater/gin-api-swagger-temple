package v1

import (
	"fmt"
	_ "github.com/bridgewwater/gin-api-swagger-temple/docs" // docs generated by swag CLI, you have to import it. use [ swag init ]
	"github.com/bridgewwater/gin-api-swagger-temple/internal/config"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swaggerInfo
//
//	want load by config.yml but https://github.com/swaggo/swag v2 not support
//
// load at config.yml
func swaggerInfo(baseURL string) {
	//genDocs.SwaggerInfo.Version = pkg_kit.GetPackageJsonVersionGoStyle()
	//genDocs.SwaggerInfo.Schemes = []string{"http", "https"}
	zlog.S().Info("=== In debug mode,you can use swagger ===")
	zlog.S().Infof("baseURL at: %v", baseURL)
	zlog.S().Infof("api MAJOR version: %v", viper.GetString("swagger.api_major_version"))
	zlog.S().Infof("swagger.security status: %v", viper.GetBool("swagger.security"))
	zlog.S().Infof("api base_path: %v", viper.GetString("base_path"))
}

// swaggerGroup
//
//	swagger:                                # swagger not show at release
//		api_major_version: v1               # api MAJOR version as https://semver.org/
//		root: /swagger                      # swagger root
//		ui_root: /editor                    # swagger ui root
//		index: /swagger/index.html          # swagger index
//		security: false                     # swagger security true or false
//		user:                               # swagger user setting of BasicAuth
//			admin: 36116f7c73bc9acb2a7a26   # admin:pwd
//			user: e2236a11aceac4de          # user:pwd
func swaggerGroup(baseURL string, g *gin.Engine) {
	// https://github.com/swaggo/swag/issues/194#issuecomment-475853710
	apiMajorVersion := viper.GetString("swagger.api_major_version")
	if apiMajorVersion == "" {
		panic("must set [ swagger.api_major_version ] at config.yml")
	}
	swaggerRoot := viper.GetString("swagger.root")
	if swaggerRoot == "" {
		swaggerRoot = "/swagger"
	}
	swaggerUiRoot := viper.GetString("swagger.ui_root")
	if swaggerUiRoot == "" {
		swaggerUiRoot = "/editor"
	}

	swaggerApiRoot := fmt.Sprintf("%s/%s", swaggerRoot, apiMajorVersion)

	configSwagger := &ginSwagger.Config{
		URL: fmt.Sprintf("%s%s_%s", config.BaseURL(), swaggerApiRoot, "doc.json"), //The url pointing to API definition
	}

	docJsonRouter := fmt.Sprintf("/%s_%s", apiMajorVersion, "doc.json")
	zlog.S().Debugf("- docJsonRouter %v", docJsonRouter)

	docJsonLocalPath := fmt.Sprintf("./docs/%s_%s", apiMajorVersion, "swagger.json")
	zlog.S().Debugf("- docJsonLocalPath %v", docJsonLocalPath)
	if viper.GetBool("swagger.security") {
		// just use swagger as swagger_user
		swaggerGroup := g.Group(swaggerRoot, gin.BasicAuth(gin.Accounts{
			"admin": viper.GetString("swagger.user.admin"),
			"user":  viper.GetString("swagger.user.user"),
		}))
		{
			swaggerGroup.StaticFile(
				docJsonRouter,
				docJsonLocalPath)
		}
		//noinspection GoTypesCompatibility
		swaggerGroup.GET(swaggerUiRoot+"/*any", ginSwagger.CustomWrapHandler(configSwagger, swaggerFiles.Handler))
	} else {
		swaggerGroup := g.Group(swaggerRoot)
		{
			swaggerGroup.StaticFile(
				docJsonRouter,
				docJsonLocalPath)
		}

		//noinspection GoTypesCompatibility
		//g.GET(swaggerRoot+"/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, envName))
		swaggerGroup.GET(swaggerUiRoot+"/*any", ginSwagger.CustomWrapHandler(configSwagger, swaggerFiles.Handler))
	}

	zlog.S().Debugf("configSwagger.URL %v", configSwagger.URL)
	zlog.S().Infof("== swagger.link at: %v%v", baseURL, viper.GetString("swagger.index"))
}
