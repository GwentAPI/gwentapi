//go:generate goagen bootstrap -d github.com/GwentAPI/gwentapi/design

package main

import (
	"flag"
	"fmt"
	"github.com/GwentAPI/gwentapi/app"
	"github.com/GwentAPI/gwentapi/configuration"
	"github.com/GwentAPI/gwentapi/dataLayer/dal"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/gzip"
	log "github.com/inconshreveable/log15"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var enableGzip bool = true
var gzipLevel int = -1

/*
Set this variable with go build with the -ldflags="-X main.version=<value>" parameter.
*/
var version = "undefined"

func init() {
	versionFlag := flag.Bool("v", false, "Prints current version")
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

	mux := http.NewServeMux()
	mountMedia(config.App.MediaPath, mux)
	mux.HandleFunc("/", service.Mux.ServeHTTP)
	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      mux,
		Addr:         ":8080",
	}
	// Close the main session
	defer dal.ShutDown()

	// Start service
	go func() {
		service.LogInfo("startup", "message", "Service is running.")
		if err := srv.ListenAndServe(); err != nil {
			service.LogError("startup", "err", err)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	service.LogInfo("shutdown", "message", "Server stopped ungracefully by killing all connections.")

}

func mountMedia(fileSystemPath string, mux *http.ServeMux) {
	fs := justFilesFilesystem{http.Dir("./data/input/media")}
	mux.Handle("/media/", http.StripPrefix("/media/", http.FileServer(fs)))
}

type justFilesFilesystem struct {
	Fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.Fs.Open(name)

	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}
