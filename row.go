package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
	"log"
)

// RowController implements the row resource.
type RowController struct {
	*goa.Controller
}

// NewRowController creates a row controller.
func NewRowController(service *goa.Service) *RowController {
	return &RowController{Controller: service.NewController("RowController")}
}

// List runs the list action.
func (c *RowController) List(ctx *app.ListRowContext) error {
	// RowController_List: start_implement
	rows, err := controllers.FetchAllRows()
	if err != nil {
		log.Println(err)
		return ctx.NotFound()
	}
	res := make(app.GwentapiRowCollection, len(rows))

	for i, row := range rows {
		r := &app.GwentapiRow{
			ID:   row.ID,
			Name: row.Name,
			Href: app.RowHref(row.ID),
		}
		res[i] = r
	}
	// RowController_List: end_implement
	return ctx.OK(res)
}

// Show runs the show action.
func (c *RowController) Show(ctx *app.ShowRowContext) error {
	// RowController_Show: start_implement
	row, err := controllers.FetchRow(ctx.RowID)
	if err != nil {
		log.Println(err)
		return ctx.NotFound()
	}

	res := &app.GwentapiRow{
		ID:   row.ID,
		Name: row.Name,
		Href: app.RowHref(row.ID),
	}
	// RowController_Show: end_implement
	return ctx.OK(res)
}
