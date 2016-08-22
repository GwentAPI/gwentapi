package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
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
	rarities, err := controllers.FetchAllRarities()
	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Rarity", "action", "List", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}
	res := make(app.GwentapiRarityCollection, len(rarities))

	for i, rarity := range rarities {
		r := &app.GwentapiRarity{
			ID:   rarity.ID,
			Href: controllers.RarityURL(rarity.ID),
			Name: rarity.Name,
		}
		res[i] = r
	}
	// RarityController_List: end_implement
	return ctx.OK(res)
}

// Show runs the show action.
func (c *RarityController) Show(ctx *app.ShowRarityContext) error {
	// RarityController_Show: start_implement
	rarity, err := controllers.FetchRarity(ctx.RarityID)
	if controllers.NotFound(err) {
		return ctx.NotFound()
	}
	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Rarity", "action", "Show", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := &app.GwentapiRarity{
		ID:   rarity.ID,
		Href: controllers.RarityURL(rarity.ID),
		Name: rarity.Name,
	}
	// RarityController_Show: end_implement
	return ctx.OK(res)
}
