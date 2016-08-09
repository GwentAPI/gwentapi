package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ArtworkType = Type("artworkType", func() {
	Description("Type of card artwork")
	Attribute("full_size", String, "Href to full size artwork")
	Attribute("artist", String, "Name of the artist")
	Attribute("normal_size", String, "Href to normal size artwork")

	Required("normal_size")
})
