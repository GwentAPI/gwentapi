package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
	"log"
)

// PatchController implements the patch resource.
type PatchController struct {
	*goa.Controller
}

// NewPatchController creates a patch controller.
func NewPatchController(service *goa.Service) *PatchController {
	return &PatchController{Controller: service.NewController("PatchController")}
}

// Latest runs the latest action.
func (c *PatchController) Latest(ctx *app.LatestPatchContext) error {
	// PatchController_Latest: start_implement
	patch, err := controllers.FetchLatestPatch()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError()
	}

	res := &app.GwentapiPatch{
		Version:     patch.Version,
		ID:          patch.ID,
		Href:        controllers.PatchURL(patch.ID),
		ReleaseDate: patch.ReleaseDate,
	}
	// PatchController_Latest: end_implement
	return ctx.OK(res)
}

// List runs the list action.
func (c *PatchController) List(ctx *app.ListPatchContext) error {
	// PatchController_List: start_implement
	patches, err := controllers.FetchAllPatches()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError()
	}
	res := make(app.GwentapiPatchCollection, len(patches))
	for i, patch := range patches {
		p := &app.GwentapiPatch{
			Version:     patch.Version,
			ID:          patch.ID,
			Href:        controllers.PatchURL(patch.ID),
			ReleaseDate: patch.ReleaseDate,
		}
		res[i] = p
	}
	// PatchController_List: end_implement
	return ctx.OK(res)
}

// Show runs the show action.
func (c *PatchController) Show(ctx *app.ShowPatchContext) error {
	// PatchController_Show: start_implement
	patch, err := controllers.FetchPatch(ctx.PatchID)
	if controllers.NotFound(err) {
		return ctx.NotFound()
	}
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError()
	}

	res := &app.GwentapiPatch{
		Version:     patch.Version,
		ID:          patch.ID,
		Href:        controllers.PatchURL(patch.ID),
		ReleaseDate: patch.ReleaseDate,
	}
	// PatchController_Show: end_implement
	return ctx.OK(res)
}
