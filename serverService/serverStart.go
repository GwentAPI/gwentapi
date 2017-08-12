// +build !linux

package serverService

import (
	"log"
	"net/http"
)

func startService(server *http.Server) {
	log.Println("Web server is starting.")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error on startup: ", err)
	}
}
