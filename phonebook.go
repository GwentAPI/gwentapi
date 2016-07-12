package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
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
		Factions: app.FactionHref(""),
		Glyphs:   app.GlyphHref(""),
		Rarities: app.RarityHref(""),
		Rows:     app.RowHref(""),
		Types:    app.TypeHref(""),
		Patches:  app.PatchHref(""),
		Cards:    app.CardHref(""),
	}
	return ctx.OK(res)
}
