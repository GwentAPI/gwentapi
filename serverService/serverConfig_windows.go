package serverService

import (
	"context"
	"github.com/GwentAPI/gwentapi/configuration"
	"github.com/goadesign/goa"
	"log"
	"sync"
)

func SocketActivatedServer(ctx context.Context, wg *sync.WaitGroup, service *goa.Service, config configuration.GwentConfig) {
	log.Println("Socket activation isn't supported on windows.")
	log.Println("Falling back to default option.")
	Server(ctx, wg, service, config)
}
