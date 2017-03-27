package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/helpers"
)

// IndexController implements the index resource.
type IndexController struct {
	*goa.Controller
}

// NewIndexController creates a index controller.
func NewIndexController(service *goa.Service) *IndexController {
	return &IndexController{Controller: service.NewController("IndexController")}
}

// Show runs the show action.
func (c *IndexController) Show(ctx *app.ShowIndexContext) error {
	// IndexController_Show: start_implement

	// IndexController_Show: end_implement

	res := &app.GwentapiResource{
		Cards:      helpers.CardURL(""),
		Factions:   helpers.FactionURL(""),
		Rarities:   helpers.RarityURL(""),
		Categories: helpers.CategoryURL(""),
		Groups:     helpers.GroupURL(""),
		Swagger:    helpers.SwaggerURL(),
	}
	return ctx.OK(res)
}
