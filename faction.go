package main

import (
	"github.com/GwentAPI/gwentapi/app"
	"github.com/GwentAPI/gwentapi/dataLayer/dal"
	"github.com/GwentAPI/gwentapi/dataLayer/factory"
	"github.com/GwentAPI/gwentapi/helpers"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"time"
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
	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalFaction(dataStore)

	factions, err := dc.FetchAll()

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Faction", "action", "List", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := make(app.GwentapiFactionCollection, len(*factions))

	lastModified := time.Time{}

	for i, faction := range *factions {
		f, _ := factory.CreateFaction(&faction)

		if lastModified.Before(faction.Last_Modified) {
			lastModified = faction.Last_Modified
		}

		res[i] = f
	}

	helpers.LastModified(ctx.ResponseData, lastModified)
	if ctx.IfModifiedSince != nil {
		if !helpers.IsModified(*ctx.IfModifiedSince, lastModified) {
			return ctx.NotModified()
		}
	}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *FactionController) Show(ctx *app.ShowFactionContext) error {
	// FactionController_Show: start_implement

	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalFaction(dataStore)
	uuid, err := helpers.DecodeUUID(ctx.FactionID)

	if err != nil {
		return ctx.NotFound()
	}

	faction, err := dc.Fetch(uuid)

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Faction", "action", "Show", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	// FactionController_Show: end_implement
	res, _ := factory.CreateFaction(faction)
	helpers.LastModified(ctx.ResponseData, faction.Last_Modified)
	if ctx.IfModifiedSince != nil {
		if !helpers.IsModified(*ctx.IfModifiedSince, faction.Last_Modified) {
			return ctx.NotModified()
		}
	}
	return ctx.OK(res)
}
