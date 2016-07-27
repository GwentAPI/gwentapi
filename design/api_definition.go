package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("gwentapi", func() {
	Title("Gwent API")
	Description("A consumption-only API of all Gwent data")
	Version("0.0")
	Contact(func() {
		Name("Tristan S.")
		Email("api.gwent@gmail.com")
		URL("https://twitter.com/GwentAPI")
	})
	License(func() {
		Name("The textual information presented through this API about GWENT: The Witcher Card Game is copyrighted by CD Projekt RED")
	})
	BasePath("/v0")
	Host("api.gwentapi.com")
	Scheme("https")
	Produces("application/json")
})
