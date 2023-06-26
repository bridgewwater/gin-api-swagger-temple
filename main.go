//go:build !test

package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/bridgewwater/gin-api-swagger-temple/api/middleware"
	"github.com/bridgewwater/gin-api-swagger-temple/pkg/pkgJson"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// done
	//	gracefully exit http server
	//	Used to synchronize main and handle Exit Signal threads
	done = make(chan bool, 1)
	// quit
	//	used to receive semaphores
	quit = make(chan os.Signal, 1)
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
	server := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}
	go handleExitSignal(server)
	err := server.ListenAndServe()
	if err != nil {
		slog.Errorf(err, "server run error %v", err)
	} else {
		slog.Infof("server run success!")
	}
}

// handleExitSignal
//
//	Handle exit signal and gracefully shut down the server
//	When the server is shut down, the main thread will be notified that the exit signal is handled
//	syscall.SIGTERM will exit 0 and syscall.SIGINT will exit 0
//	operator is kill pid or ctrl + c
func handleExitSignal(s *http.Server) {
	// listen for the following two semaphores
	signal.Notify(quit, syscall.SIGTERM) // kill
	signal.Notify(quit, syscall.SIGINT)  // ctrl + c
	// blocking wait semaphore
	<-quit

	// shuts down the server, causing the Listen And Serve function to return
	if err := s.Shutdown(context.Background()); err != nil {
		fmt.Printf("ShutDown Error: %v", err)
	}
	// notify the main thread handle exit signal is over
	close(done)
}
