package serverService

import (
	"context"
	"github.com/GwentAPI/gwentapi/configuration"
	"github.com/coreos/go-systemd/activation"
	"github.com/goadesign/goa"
	"log"
	"net/http"
	"sync"
	"time"
)

func SocketActivatedServer(ctx context.Context, wg *sync.WaitGroup, service *goa.Service, config configuration.GwentConfig) {
	defer wg.Done()

	listeners, err := activation.Listeners(true)
	if err != nil {
		log.Panic(err)
	}

	if len(listeners) != 1 {
		log.Panic("Unexpected number of socket activation fds. Number: ", len(listeners))
	}

	mux := http.NewServeMux()
	mountMedia(config.App.MediaPath, mux)
	mux.HandleFunc("/", service.Mux.ServeHTTP)
	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      mux,
		Addr:         config.App.Port,
	}

	// Start service
	go func() {
		service.LogInfo("startup", "message", "Web server is starting.")
		if err := srv.Serve(listeners[0]); err != nil {
			service.LogError("startup", "err", err)
		}
	}()
	<-ctx.Done()
	//log.Debug("Shutdown in progress.")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.Shutdown(shutdownCtx)
}
