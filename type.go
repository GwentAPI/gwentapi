package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
)

// TypeController implements the type resource.
type TypeController struct {
	*goa.Controller
}

// NewTypeController creates a type controller.
func NewTypeController(service *goa.Service) *TypeController {
	return &TypeController{Controller: service.NewController("TypeController")}
}

// List runs the list action.
func (c *TypeController) List(ctx *app.ListTypeContext) error {
	// TypeController_List: start_implement
	cardTypes, err := controllers.FetchAllTypes()
	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Type", "action", "List", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := make(app.GwentapiTypeCollection, len(cardTypes))

	for i, cardType := range cardTypes {
		t := &app.GwentapiType{
			ID:   cardType.ID,
			Href: controllers.TypeURL(cardType.ID),
			Name: cardType.Name,
		}
		res[i] = t
	}
	// TypeController_List: end_implement
	return ctx.OK(res)
}

// Show runs the show action.
func (c *TypeController) Show(ctx *app.ShowTypeContext) error {
	// TypeController_Show: start_implement
	cardType, err := controllers.FetchType(ctx.TypeID)
	if controllers.NotFound(err) {
		return ctx.NotFound()
	}
	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Type", "action", "Show", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := &app.GwentapiType{
		ID:   cardType.ID,
		Href: controllers.TypeURL(cardType.ID),
		Name: cardType.Name,
	}
	// TypeController_Show: end_implement
	return ctx.OK(res)
}
