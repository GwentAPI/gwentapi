package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var FactionMedia = MediaType("application/vnd.gwentapi.faction+json", func() {
	Description("A faction")
	Attributes(func() {
		Attribute("id", String, "Unique faction ID")
		Attribute("href", String, "API href for making requests on the faction")
		Attribute("name", String, "Name of the faction")
		Required("id", "href", "name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})
