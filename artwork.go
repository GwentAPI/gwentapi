package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
)

// ArtworkController implements the artwork resource.
type ArtworkController struct {
	*goa.Controller
}

// NewArtworkController creates a artwork controller.
func NewArtworkController(service *goa.Service) *ArtworkController {
	return &ArtworkController{Controller: service.NewController("ArtworkController")}
}

// Show runs the show action.
func (c *ArtworkController) Show(ctx *app.ShowArtworkContext) error {
	// ArtworkController_Show: start_implement

	// Put your logic here
	_, err := controllers.FetchCard(ctx.CardID)
	if err != nil {
		return ctx.InternalServerError()
	}

	// ArtworkController_Show: end_implement
	res := createArtwork(ctx.CardID)
	return ctx.OK(res)
}

func createArtwork(id string) *app.GwentapiArtwork {
	//var artist string = "test"
	//var artistP *string = &artist

	at := &app.ArtworkType{
		//Artist: artistP,
		Full:  controllers.MediaURL(id) + ".png",
		Small: controllers.MediaURL(id) + "_small.png",
	}

	a := &app.GwentapiArtwork{
		ID:      id,
		Href:    controllers.ArtworkURL(id),
		Artwork: at,
	}

	return a
}
