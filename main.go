//go:generate goagen bootstrap -d github.com/GwentAPI/gwentapi/design

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/GwentAPI/gwentapi/app"
	"github.com/GwentAPI/gwentapi/configuration"
	"github.com/GwentAPI/gwentapi/dataLayer/dal"
	"github.com/GwentAPI/gwentapi/serverService"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/gzip"
	log "github.com/inconshreveable/log15"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var daemonize *bool
var enableGzip bool = true
var gzipLevel int = -1

/*
Set this variable with go build with the -ldflags="-X main.version=<value>" parameter.
*/
var version = "undefined"

func init() {
	versionFlag := flag.Bool("v", false, "Prints current version")
	daemonize = flag.Bool("daemon", false, "Linux option only: Daemonize the program by using systemd socket activation feature. Be sure to set up the systemd service first.")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}
}

func main() {
	// Create service
	service := goa.New("gwentapi")

	//Create logger
	logger := log.New("module", "app/server")

	config := configuration.GetConfig()
	//Logger configuration
	logger.SetHandler(log.MultiHandler(
		log.LvlFilterHandler(log.LvlInfo, log.Must.FileHandler(config.App.LogInfoFile, log.LogfmtFormat())),
		log.LvlFilterHandler(log.LvlError, log.Must.FileHandler(config.App.LogErrorFile, log.LogfmtFormat()))))

	//Inject logger
	service.WithLogger(goalog15.New(logger))

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(config.App.Verbose))
	service.Use(middleware.ErrorHandler(service, config.App.Verbose))
	service.Use(middleware.Recover())
	if enableGzip {
		service.Use(gzip.Middleware(gzipLevel))
	}
	// Mount "card" controller
	c := NewCardController(service)
	app.MountCardController(service, c)
	// Mount "faction" controller
	c2 := NewFactionController(service)
	app.MountFactionController(service, c2)
	// Mount "index" controller
	c3 := NewIndexController(service)
	app.MountIndexController(service, c3)
	// Mount "category" controller
	c4 := NewCategoryController(service)
	app.MountCategoryController(service, c4)
	// Mount "rarity" controller
	c6 := NewRarityController(service)
	app.MountRarityController(service, c6)
	// Mount "group" controller
	c8 := NewGroupController(service)
	app.MountGroupController(service, c8)

	// Close the main session
	defer dal.ShutDown()

	// create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// a WaitGroup for the goroutines to tell us they've stopped
	wg := sync.WaitGroup{}
	wg.Add(1)
	if *daemonize {
		go serverService.SocketActivatedServer(ctx, &wg, service, config)
	} else {
		go serverService.Server(ctx, &wg, service, config)
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	//log.Debug("Received signal: shutting down.")
	//log.Debug("Telling goroutines to stop")
	cancel()

	wg.Wait()
	//log.Debug("All goroutines have told us they've finished.")
	service.LogInfo("shutdown", "message", "Server gracefully stopped.")
}
