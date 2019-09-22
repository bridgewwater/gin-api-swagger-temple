package router

import (
	"github.com/bar-counter/monitor"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

// use monitor https://github.com/bar-counter/monitor
func monitorAPI(g *gin.Engine) {
	monitorCfg := &monitor.Cfg{
		Status:         viper.GetBool("monitor.status"),
		StatusHardware: viper.GetBool("monitor.status_hardware"),
		Debug:          viper.GetBool("monitor.debug"),
		DebugMiddleware: gin.BasicAuth(gin.Accounts{
			"admin": viper.GetString("monitor.admin.pwd"),
		}),
		PProf: viper.GetBool("monitor.pprof"),
	}
	err := monitor.Register(g, monitorCfg)
	if err != nil {
		log.Errorf(err, "monitor Register err %v", err)
	}
}
