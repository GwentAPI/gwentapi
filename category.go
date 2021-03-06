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

// CategoryController implements the category resource.
type CategoryController struct {
	*goa.Controller
}

// NewCategoryController creates a category controller.
func NewCategoryController(service *goa.Service) *CategoryController {
	return &CategoryController{Controller: service.NewController("CategoryController")}
}

// List runs the list action.
func (c *CategoryController) List(ctx *app.ListCategoryContext) error {
	// CategoryController_List: start_implement
	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalCategory(dataStore)

	categories, err := dc.FetchAll()

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Category", "action", "List", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := make(app.GwentapiCategoryCollection, len(*categories))

	lastModified := time.Time{}

	for i, category := range *categories {
		c, _ := factory.CreateCategory(&category)

		if lastModified.Before(category.Last_Modified) {
			lastModified = category.Last_Modified
		}

		res[i] = c
	}
	// CategoryController_List: end_implement
	helpers.LastModified(ctx.ResponseData, lastModified)
	if ctx.IfModifiedSince != nil {
		if !helpers.IsModified(*ctx.IfModifiedSince, lastModified) {
			return ctx.NotModified()
		}
	}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *CategoryController) Show(ctx *app.ShowCategoryContext) error {
	// CategoryController_Show: start_implement

	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalCategory(dataStore)
	uuid, err := helpers.DecodeUUID(ctx.CategoryID)

	if err != nil {
		return ctx.NotFound()
	}

	category, err := dc.Fetch(uuid)

	if helpers.IsNotFoundError(err) {
		return ctx.NotFound()
	} else if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Category", "action", "Show", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}
	// CategoryController_Show: end_implement
	res, _ := factory.CreateCategory(category)
	helpers.LastModified(ctx.ResponseData, category.Last_Modified)
	if ctx.IfModifiedSince != nil {
		if !helpers.IsModified(*ctx.IfModifiedSince, category.Last_Modified) {
			return ctx.NotModified()
		}
	}
	return ctx.OK(res)
}
