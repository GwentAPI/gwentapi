package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ArtType = Type("artType", func() {
	Description("Type of card art")
	Attribute("fullsizeImage", String, "Href to full size artwork", func() {
		Format("uri")
	})
	Attribute("artist", String, "Name of the artist", func() {
		Example("Marek Madej")
	})
	Attribute("thumbnailImage", String, "Href to thumbnail size artwork", func() {
		Format("uri")
	})

	Required("thumbnailImage")
})

var CostType = Type("costType", func() {
	Description("Type used to define the associated crafting/milling cost cost")
	Attribute("premium", Integer, "Premium cost", func() {
		Example(200)
	})
	Attribute("normal", Integer, "Normal cost", func() {
		Example(30)
	})

	Required("premium", "normal")
})
