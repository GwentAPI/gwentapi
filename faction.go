package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
	"log"
)

// FactionController implements the faction resource.
type FactionController struct {
	*goa.Controller
}

// NewFactionController creates a faction controller.
func NewFactionController(service *goa.Service) *FactionController {
	return &FactionController{Controller: service.NewController("FactionController")}
}

// List runs the list action.
func (c *FactionController) List(ctx *app.ListFactionContext) error {
	// FactionController_List: start_implement
	factions, err := controllers.FetchAllFactions()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError()
	}
	res := make(app.GwentapiFactionCollection, len(factions))

	for i, faction := range factions {
		f := &app.GwentapiFaction{
			ID:   faction.ID,
			Href: app.FactionHref(faction.ID),
			Name: faction.Name,
		}
		res[i] = f
	}
	// FactionController_List: end_implement
	//res := app.GwentapiFactionCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *FactionController) Show(ctx *app.ShowFactionContext) error {
	// FactionController_Show: start_implement
	faction, err := controllers.FetchFaction(ctx.FactionID)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError()
	}
	// FactionController_Show: end_implement
	res := &app.GwentapiFaction{
		ID:   faction.ID,
		Name: faction.Name,
		Href: app.FactionHref(faction.ID),
	}
	return ctx.OK(res)
}
