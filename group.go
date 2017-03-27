package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/dal"
	"github.com/tri125/gwentapi/dataLayer/factory"
	"github.com/tri125/gwentapi/helpers"
)

// GroupController implements the group resource.
type GroupController struct {
	*goa.Controller
}

// NewGroupController creates a group controller.
func NewGroupController(service *goa.Service) *GroupController {
	return &GroupController{Controller: service.NewController("GroupController")}
}

// List runs the list action.
func (c *GroupController) List(ctx *app.ListGroupContext) error {
	// GroupController_List: start_implement

	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalGroup(dataStore)

	groups, err := dc.FetchAll()

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Group", "action", "List", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := make(app.GwentapiGroupCollection, len(*groups))

	for i, group := range *groups {
		g, _ := factory.CreateGroup(&group)
		res[i] = g
	}

	// GroupController_List: end_implement
	return ctx.OK(res)
}

// Show runs the show action.
func (c *GroupController) Show(ctx *app.ShowGroupContext) error {
	// GroupController_Show: start_implement

	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalGroup(dataStore)
	uuid, err := helpers.Base64ToUUID(ctx.GroupID)

	if err != nil {
		return ctx.NotFound()
	}

	group, err := dc.Fetch(uuid)

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Group", "action", "Show", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	// GroupController_Show: end_implement
	res, _ := factory.CreateGroup(group)
	return ctx.OK(res)
}
