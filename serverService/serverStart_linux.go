package serverService

import (
	"github.com/coreos/go-systemd/activation"
	"log"
	"net/http"
)

func startService(server *http.Server) {
	listeners, err := activation.Listeners(true)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Web server is starting.")
	// No socket passed from systemd, we aren't running with socket activation.
	if len(listeners) == 0 {
		// Fallback
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Error on startup: ", err)
		}
	} else {
		// Socket activation code
		log.Println("Listening to sockets.")
		if len(listeners) != 1 {
			log.Fatal("Unexpected number of socket activation fds. Number: ", len(listeners))
		}
		if err := server.Serve(listeners[0]); err != nil {
			log.Fatal("Error starting socket activated web server: ", err)
		}
	}
}
