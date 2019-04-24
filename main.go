package main

import (
	"fmt"
	"net/http"
	"time"

	"git.sinlov.cn/bridgewwater/temp-gin-api-self/config"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/docs"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/router"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "api server config file path.")
)

// @termsOfService http://git.sinlov.cn/
// @contact.name API Support
// @contact.url http://git.sinlov.cn/
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
	fmt.Printf("%s \n", "start dev app at here")

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

	var apiBase = config.BaseURL()
	if "debug" == runMode || "test" == runMode {
		// set swagger info
		docs.SwaggerInfo.Title = viper.GetString("swagger.title")
		docs.SwaggerInfo.Description = viper.GetString("swagger.description")
		docs.SwaggerInfo.Version = viper.GetString("swagger.version")
		docs.SwaggerInfo.Host = viper.GetString("swagger.host")
		docs.SwaggerInfo.BasePath = viper.GetString("base_path")
		log.Infof("In debug mode,you can use swagger")
		log.Infof("swagger.link at: %v%v", apiBase, viper.GetString("swagger.index"))
		log.Infof("swagger.security status: %v", viper.GetBool("swagger.security"))
		// Ping the server to make sure the router is working.
		go func() {
			if err := pingServer(apiBase); err != nil {
				log.Error("The router has no response, or it might took too long to start up.", err)
			}
			log.Info("The router has been deployed successfully.")
		}()
	} else if "test" == runMode {
		log.Infof("In test mode, you can use swagger.link at: %v%v", apiBase, viper.GetString("swagger.index"))
	} else {
	}
	log.Infof("Start to listening the incoming requests on http address: %v", viper.GetString("addr"))
	log.Infof("Sever name: %v , has start!", viper.GetString("name"))
	err := http.ListenAndServe(viper.GetString("addr"), g)
	if err != nil {
		log.Errorf(err, "server run error %v", err)
	} else {
		log.Infof("server run success!")
	}
}

// pingServer pings the http server to make sure the router is working.
func pingServer(api string) error {
	pingApi := api + viper.GetString("base_path") + viper.GetString("ssc.health")
	log.Infof("pingServer test api as: %v", pingApi)

	for i := 0; i < viper.GetInt("ssc.count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(pingApi)
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Warn("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	//noinspection ALL
	return fmt.Errorf("Can not connect to the router.")
}
