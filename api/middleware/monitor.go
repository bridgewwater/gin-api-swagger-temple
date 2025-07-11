package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bar-counter/monitor/v2"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// monitorAPI
//
//	use monitor https://github.com/bar-counter/monitor
//
//	config.yml like
//
//	monitor: # monitor
//		status: true             # api status use {monitor.health}
//		health: /status/health   # api health
//		retryCount: 10           # ping api health retry count
//		hardware: true           # hardware true or false
//		status_hardware:
//			disk: /status/hardware/disk     # hardware api disk
//			cpu: /status/hardware/cpu       # hardware api cpu
//			ram: /status/hardware/ram       # hardware api ram
//		debug: true                         # debug true or false
//		pprof: true                         # security true or false
//		security: false                     # debug and security true or false
//		securityUser:
//			admin: pwd # admin:pwd
func monitorAPI(g *gin.Engine) {
	var monitorCfg *monitor.Cfg

	isSecurity := viper.GetBool("monitor.security")
	if isSecurity {
		monitorCfg = &monitor.Cfg{
			Status:         viper.GetBool("monitor.status"),
			StatusHardware: viper.GetBool("monitor.hardware"),
			Debug:          viper.GetBool("monitor.debug"),
			DebugMiddleware: gin.BasicAuth(gin.Accounts{
				"admin": viper.GetString("monitor.securityUser.admin"),
			}),
			PProf: viper.GetBool("monitor.pprof"),
		}
	} else {
		monitorCfg = &monitor.Cfg{
			Status:         viper.GetBool("monitor.status"),
			StatusHardware: viper.GetBool("monitor.status_hardware"),
			Debug:          viper.GetBool("monitor.debug"),
			PProf:          viper.GetBool("monitor.pprof"),
		}
	}

	err := monitor.Register(g, monitorCfg)
	if err != nil {
		zlog.S().Errorf("monitor Register err %v", err)
	}
}

// checkPingServer
//
//	ping the server to make sure the router is working.
//	use config.yml as
//
//	apiBaseURL load by github.com/spf13/viper
//
// viper config.yml
//
//	monitor: # monitor
//		status: true             # api status use {monitor.health}
//		health: /status/health   # api health
//		retryCount: 10           # ping api health retry count
func checkPingServer(apiBaseURL string) {
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(apiBaseURL, viper.GetString("monitor.health")); err != nil {
			zlog.S().
				Errorf("The router has no response, or it might took too long to start up. err %v", err)
		}

		zlog.S().Info("The router has been deployed successfully.")
	}()
}

// PingServer
//
//	ping server pings the http server
//
//	apiBaseURL load by github.com/spf13/viper
//	checkRouter monitor.health by github.com/spf13/viper
//
//	viper config.yml
//	monitor: # monitor
//		status: true             # api status use {monitor.health}
//		health: /status/health   # api health
//		retryCount: 10           # ping api health retry count
func pingServer(apiBaseURL, checkRouter string) error {
	pingApi := apiBaseURL + checkRouter
	zlog.S().Infof("pingServer test api : %v", pingApi)

	retryCount := viper.GetInt("monitor.retryCount")
	if retryCount > 0 {
		for range retryCount {
			// Ping the server by sending a GET request to `/health`.
			resp, err := http.Get(pingApi)
			if err == nil && resp.StatusCode == http.StatusOK {
				zlog.S().Infof("pingServer test pass api at: %v", pingApi)

				return nil
			}

			// sleep for a second to continue the next ping.
			zlog.S().Warnf("Waiting for the router, retry in 1 second. Check URL: %v", pingApi)
			time.Sleep(time.Second)
		}
	}

	//noinspection ALL
	return fmt.Errorf("Can not connect to the router %v.", pingApi)
}
