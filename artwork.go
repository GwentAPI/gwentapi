package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
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

	// ArtworkController_Show: end_implement
	res := &app.GwentapiArtwork{}
	return ctx.OK(res)
}
