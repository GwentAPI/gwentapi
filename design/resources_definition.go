package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("faction", func() {

	DefaultMedia(FactionMedia)
	BasePath("/factions")

	Action("list", func() {
		Routing(GET(""))

		Description("Return all factions.")
		Response(OK, CollectionOf(FactionMedia))
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return faction with given id.")
		Routing(GET("/:factionID"))
		Params(func() {
			Param("factionID", String, "Faction ID")
		})
		Response(OK)
		Response(NotFound)

	})
})
