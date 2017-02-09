package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ArtType = Type("artType", func() {
	Description("Type of card art")
	Attribute("fullsizeImage", String, "Href to full size artwork")
	Attribute("artist", String, "Name of the artist")
	Attribute("thumbnailImage", String, "Href to thumbnail size artwork")

	//Required("thumbnailImage")
})

var CostType = Type("costType", func() {
	Description("Type used to define the associated crafting/milling cost cost")
	Attribute("premium", Integer, "Premium cost")
	Attribute("normal", Integer, "Normal cost")

	Required("premium", "normal")
})
