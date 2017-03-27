package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/dal"
	"github.com/tri125/gwentapi/dataLayer/factory"
	"github.com/tri125/gwentapi/helpers"
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

	for i, rarity := range *rarities {
		r, _ := factory.CreateRarity(&rarity)
		res[i] = r
	}

	// RarityController_List: end_implement
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
	uuid, err := helpers.Base64ToUUID(ctx.RarityID)

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
	return ctx.OK(res)
}
