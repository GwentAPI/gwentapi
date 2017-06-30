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

// RarityController implements the rarity resource.
type RarityController struct {
	*goa.Controller
}

// NewRarityController creates a rarity controller.
func NewRarityController(service *goa.Service) *RarityController {
	return &RarityController{Controller: service.NewController("RarityController")}
}

// List runs the list action.
func (c *RarityController) List(ctx *app.ListRarityContext) error {
	// RarityController_List: start_implement
	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalRarity(dataStore)

	rarities, err := dc.FetchAll()

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Rarity", "action", "List", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := make(app.GwentapiRarityCollection, len(*rarities))

	lastModified := time.Time{}
	for i, rarity := range *rarities {
		r, _ := factory.CreateRarity(&rarity)

		if lastModified.Before(rarity.Last_Modified) {
			lastModified = rarity.Last_Modified
		}

		res[i] = r
	}

	// RarityController_List: end_implement
	helpers.LastModified(ctx.ResponseData, lastModified)
	if ctx.IfModifiedSince != nil {
		if !helpers.IsModified(*ctx.IfModifiedSince, lastModified) {
			return ctx.NotModified()
		}
	}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *RarityController) Show(ctx *app.ShowRarityContext) error {
	// RarityController_Show: start_implement

	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalRarity(dataStore)
	uuid, err := helpers.DecodeUUID(ctx.RarityID)

	if err != nil {
		return ctx.NotFound()
	}

	rarity, err := dc.Fetch(uuid)

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Rarity", "action", "Show", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	// RarityController_Show: end_implement
	res, _ := factory.CreateRarity(rarity)
	helpers.LastModified(ctx.ResponseData, rarity.Last_Modified)
	if ctx.IfModifiedSince != nil {
		if !helpers.IsModified(*ctx.IfModifiedSince, rarity.Last_Modified) {
			return ctx.NotModified()
		}
	}
	return ctx.OK(res)
}
