//go:build !test

package main

import (
	_ "embed"
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/bridgewwater/gin-api-swagger-temple/api/middleware"
	"github.com/bridgewwater/gin-api-swagger-temple/pkg/pkgJson"
	"net/http"
	"time"

	"github.com/bridgewwater/gin-api-swagger-temple/api/v1"
	"github.com/bridgewwater/gin-api-swagger-temple/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

//go:embed package.json
var packageJson string

var (
	cfg = pflag.StringP("config", "c", "", "api server config file path.")
)

func main() {
	pflag.Parse()
	pkgJson.InitPkgJsonContent(packageJson)

	// init config
	if err := config.Init(*cfg); err != nil {
		fmt.Printf("Error, run service not use -c or config yaml error, more info: %v\n", err)
		panic(err)
	}
	fmt.Printf("=> config init success, now api [ %s ] version: [ %v ]\n", pkgJson.GetPackageJsonName(), pkgJson.GetPackageJsonVersionGoStyle())
	fmt.Printf("-> start service %v at time: %v\n", viper.GetString("name"), time.Now().String())

	// Create the Gin engine.
	g := gin.New()

	var middlewareList []gin.HandlerFunc

	// usage middleware.
	middleware.Usage(g, middlewareList...)

	// Routes.
	v1.Register(
		// Cores.
		g,

		// middlewareList.
		middlewareList...,
	)

	slog.Warnf("-> Sever name: [ %s ], try ListenAndServe address: %s", viper.GetString("name"), viper.GetString("addr"))
	err := http.ListenAndServe(viper.GetString("addr"), g)
	if err != nil {
		slog.Errorf(err, "server run error %v", err)
	} else {
		slog.Infof("server run success!")
	}
}
