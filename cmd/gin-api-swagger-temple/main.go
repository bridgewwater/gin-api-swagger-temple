//go:build !test

package main

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	gin_api_swagger_temple "github.com/bridgewwater/gin-api-swagger-temple"
	"github.com/bridgewwater/gin-api-swagger-temple/api/middleware"
	v1 "github.com/bridgewwater/gin-api-swagger-temple/api/v1"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/config"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/constant"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/pkg/pkg_kit"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	help        = pflag.BoolP("help", "h", false, "help info.")
	versionFlag = pflag.BoolP("version", "v", false, "show version")
	cfg         = pflag.StringP("config", "c", "", "api server config file path.")

	// done
	//	gracefully exit http server
	//	Used to synchronize main and handle Exit Signal threads
	done = make(chan bool, 1)
	// quit
	//	used to receive semaphores
	quit = make(chan os.Signal, 1)
)

//nolint:gochecknoglobals
var (
	// Populated by goreleaser during build.
	version    = "unknown"
	rawVersion = "unknown"
	buildID    string
	commit     = "?"
	date       = ""
)

func init() {
	if buildID == "" {
		buildID = "unknown"
	}
}

func main() {
	pflag.Parse()
	pkg_kit.InitPkgJsonContent(gin_api_swagger_temple.PackageJson)

	bdInfo := pkg_kit.NewBuildInfo(
		pkg_kit.GetPackageJsonName(),
		pkg_kit.GetPackageJsonDescription(),
		version,
		rawVersion,
		buildID,
		commit,
		date,
		pkg_kit.GetPackageJsonAuthor().Name,
		constant.CopyrightStartYear,
	)
	pkg_kit.SaveBuildInfo(&bdInfo)

	versionStr := bdInfo.VersionString()
	versionInfoStr := bdInfo.String()

	if *help {
		pflag.Usage()
		fmt.Printf("\n%s\n", versionInfoStr)

		return
	}

	if *versionFlag {
		fmt.Printf(
			"%s version %s - %s\n",
			bdInfo.PkgName,
			pkg_kit.FetchNowVersion(),
			pkg_kit.FetchNowBuildCode(),
		)

		return
	}

	// init config
	if err := config.Init(*cfg, bdInfo); err != nil {
		fmt.Printf("Error, run service not use -c or config yaml error, more info: %v\n", err)
		panic(err)
	}

	fmt.Printf("=> config init success, now %s\n",
		versionStr,
	)
	fmt.Printf("-> by: %s, run on %s %s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
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

	zlog.S().
		Warnf("-> Sever name: [ %s ], try ListenAndServe address: %s", viper.GetString("name"), config.Addr())

	server := &http.Server{
		Addr:    config.Addr(),
		Handler: g,
	}
	go handleExitSignal(server)

	err := server.ListenAndServe()
	if err != nil {
		zlog.S().Errorf("server run error %v", err)
	} else {
		zlog.S().Infof("server run success!")
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
