package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bridgewwater/gin-api-swagger-temple/config"
	"github.com/bridgewwater/gin-api-swagger-temple/router"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "api server config file path.")
)

// @termsOfService http://github.com/

// @contact.name API Support
// @contact.url http://github.com/
// @contact.email support@sinlov.cn

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		fmt.Printf("Error, run service not use -c or config yaml error, more info: %v\n", err)
		panic(err)
	}
	fmt.Printf("%s -> %v at time: %v\n", "start service", viper.GetString("name"), time.Now().String())

	// Set gin mode.
	runMode := viper.GetString("runmode")
	gin.SetMode(runMode)

	// Create the Gin engine.
	g := gin.New()

	var middlewareList []gin.HandlerFunc
	// Routes.
	router.Load(
		// Cores.
		g,

		// middlewareList.
		middlewareList...,
	)

	log.Infof("Start to listening the incoming requests on http address: %v", viper.GetString("addr"))
	log.Infof("Sever name: %v , has start!", viper.GetString("name"))
	err := http.ListenAndServe(viper.GetString("addr"), g)
	if err != nil {
		log.Errorf(err, "server run error %v", err)
	} else {
		log.Infof("server run success!")
	}
}
