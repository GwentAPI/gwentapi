//go:generate goagen bootstrap -d github.com/tri125/gwentapi/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/gzip"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
)

func main() {
	// Create service
	service := goa.New("gwentapi")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	service.Use(gzip.Middleware(-1))
	// Mount "card" controller
	c := NewCardController(service)
	app.MountCardController(service, c)
	// Mount "faction" controller
	c2 := NewFactionController(service)
	app.MountFactionController(service, c2)
	// Mount "glyph" controller
	c3 := NewGlyphController(service)
	app.MountGlyphController(service, c3)
	// Mount "patch" controller
	c4 := NewPatchController(service)
	app.MountPatchController(service, c4)
	// Mount "phonebook" controller
	c5 := NewPhonebookController(service)
	app.MountPhonebookController(service, c5)
	// Mount "rarity" controller
	c6 := NewRarityController(service)
	app.MountRarityController(service, c6)
	// Mount "row" controller
	c7 := NewRowController(service)
	app.MountRowController(service, c7)
	// Mount "type" controller
	c8 := NewTypeController(service)
	app.MountTypeController(service, c8)

	//database
	var err error
	controllers.DBCon, err = controllers.NewDBConnection("client:thereisonlyonekinginthenorth@tcp/gwentapi?timeout=90s&strict=true&parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer controllers.DBCon.Close()

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
