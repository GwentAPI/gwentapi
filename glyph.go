package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
	"log"
)

// GlyphController implements the glyph resource.
type GlyphController struct {
	*goa.Controller
}

// NewGlyphController creates a glyph controller.
func NewGlyphController(service *goa.Service) *GlyphController {
	return &GlyphController{Controller: service.NewController("GlyphController")}
}

// List runs the list action.
func (c *GlyphController) List(ctx *app.ListGlyphContext) error {
	// GlyphController_List: start_implement
	glyphs, err := controllers.FetchAllGlyphs()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError()
	}
	// GlyphController_List: end_implement
	res := make(app.GwentapiGlyphCollection, len(glyphs))

	for i, glyph := range glyphs {
		g := &app.GwentapiGlyph{
			ID:             glyph.ID,
			Href:           controllers.GlyphURL(glyph.ID),
			IsWeatherGlyph: glyph.IsWeatherGlyph,
			Text:           glyph.Text,
			Name:           glyph.Name,
		}
		res[i] = g
	}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *GlyphController) Show(ctx *app.ShowGlyphContext) error {
	// GlyphController_Show: start_implement
	glyph, err := controllers.FetchGlyph(ctx.GlyphID)
	if controllers.NotFound(err) {
		return ctx.NotFound()
	}
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError()
	}

	res := &app.GwentapiGlyph{
		ID:             glyph.ID,
		Href:           controllers.GlyphURL(glyph.ID),
		IsWeatherGlyph: glyph.IsWeatherGlyph,
		Text:           glyph.Text,
		Name:           glyph.Name,
	}

	// GlyphController_Show: end_implement
	return ctx.OK(res)
}
