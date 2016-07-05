package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("gwentapi", func() {
	Title("Gwent API")
	Description("A consumption-only API of all Gwent data")
	BasePath("/")
	Host("localhost:8080")
	Scheme("https")
})

var _ = Resource("faction", func() {
	DefaultMedia(FactionMedia)
	BasePath("/factions")

	Action("show", func() {
		Description("Return faction with given id")
		Routing(GET("/:factionID"))
		Params(func() {
			Param("factionID", String, "Faction ID")
		})
		Response(OK)
		Response(NotFound)

	})
})

var FactionMedia = MediaType("application/vnd.gwentapi.faction+json", func() {
	Description("A faction")
	Attributes(func() {
		Attribute("id", Integer, "Unique faction ID")
		Attribute("href", String, "API href for making requests on the faction")
		Attribute("name", String, "Name of the faction")
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})
})
