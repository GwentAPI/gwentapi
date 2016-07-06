package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("gwentapi", func() {
	Title("Gwent API")
	Description("A consumption-only API of all Gwent data")
	Contact(func() {
		Name("Tristan S.")
		Email("gwentapi@gmail.com")
		URL("https://twitter.com/GwentAPI")
	})
	License(func() {
		Name("The textual information presented through this API about GWENT: The Witcher Card Game is copyrighted by CD Projekt RED")
	})
	BasePath("/")
	Host("localhost:8080")
	Scheme("https")
	Produces("application/json")

	Origin("localhost", func() {
		Methods("GET")
		MaxAge(600)
	})

	ResponseTemplate("static", func(description string) {
		Description(description)
		Status(404)
		Media("application/json")
	})
})
