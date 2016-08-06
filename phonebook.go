package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
)

// PhonebookController implements the phonebook resource.
type PhonebookController struct {
	*goa.Controller
}

// NewPhonebookController creates a phonebook controller.
func NewPhonebookController(service *goa.Service) *PhonebookController {
	return &PhonebookController{Controller: service.NewController("PhonebookController")}
}

// Show runs the show action.
func (c *PhonebookController) Show(ctx *app.ShowPhonebookContext) error {
	// PhonebookController_Show: start_implement

	// PhonebookController_Show: end_implement
	res := &app.GwentapiResource{
		Factions: controllers.FactionURL(""),
		Glyphs:   controllers.GlyphURL(""),
		Rarities: controllers.RarityURL(""),
		Types:    controllers.TypeURL(""),
		Patches:  controllers.PatchURL(""),
		Cards:    controllers.CardURL(""),
	}
	return ctx.OK(res)
}
