package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ResourceMedia = MediaType("application/vnd.gwentapi.resource+json", func() {
	Description("Listing of all available resource endpoint")
	Attributes(func() {
		Attribute("cards", String, "API href for making requests on cards")
		Attribute("factions", String, "API href for making requests on factions")
		Attribute("glyphs", String, "API href for making requests on glyphs")
		Attribute("rarities", String, "API href for making requests on rarities")
		Attribute("rows", String, "API href for making requests on rows")
		Attribute("types", String, "API href for making requests on types")
		Attribute("patches", String, "API href for making requests on patches")

		Required("cards", "factions", "glyphs", "rarities", "rows", "types", "patches")
	})

	View("default", func() {
		Attribute("cards")
		Attribute("factions")
		Attribute("glyphs")
		Attribute("rarities")
		Attribute("rows")
		Attribute("types")
		Attribute("patches")
	})
})

var CardMedia = MediaType("application/vnd.gwentapi.card+json", func() {
	Description("A card")
	Attributes(func() {
		Attribute("id", String, "Unique card ID")
		Attribute("href", String, "API href for making requests on the card")
		Attribute("name", String, "Name of the card")
		Attribute("type", TypeMedia, "Type of the card")
		Attribute("faction", FactionMedia, "Faction of the card")

		Attribute("subtypes", CollectionOf(TypeMedia), "Subtypes of the card")
		Attribute("rows", CollectionOf(RowMedia), "Rows where the card can be played")

		Attribute("rarity", RarityMedia, "Rarity of the card")
		Attribute("strength", Integer, "Strength of the card")
		Attribute("text", String, "Text of the card detailing its abilities and how it plays")
		Attribute("flavor", String, "Flavor text of the card")

		Required("id", "href", "name", "type", "faction", "rarity")
	})

	Links(func() {
		Link("type")
		Link("faction")
		Link("subtypes")
		Link("rows")
		Link("rarity")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("type", func() {
			View("link")
		})
		Attribute("faction", func() {
			View("link")
		})
		Attribute("subtypes", func() {
			View("link")
		})
		Attribute("rows", func() {
			View("link")
		})
		Attribute("rarity", func() {
			View("link")
		})
		Attribute("strength")
		Attribute("text")
		Attribute("flavor")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

})

var FactionMedia = MediaType("application/vnd.gwentapi.faction+json", func() {
	Description("A card faction")
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

var GlyphMedia = MediaType("application/vnd.gwentapi.glyph+json", func() {
	Description("A glyph")
	Attributes(func() {
		Attribute("id", String, "Unique glyph ID")
		Attribute("href", String, "API href for making requests on the glyph")
		Attribute("name", String, "Name of the glyph")
		Attribute("isWeatherGlyph", Boolean, "Indicate whether or not the glyph is a weather glyph")
		Attribute("text", String, "Text of the glyph")

		Required("id", "href", "name", "isWeatherGlyph", "text")
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("isWeatherGlyph")
		Attribute("text")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})

var RarityMedia = MediaType("application/vnd.gwentapi.rarity+json", func() {
	Description("A card rarity")
	Attributes(func() {
		Attribute("id", String, "Unique rarity ID")
		Attribute("href", String, "API href for making requests on the rarity")
		Attribute("name", String, "Name of the rarity")

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

var TypeMedia = MediaType("application/vnd.gwentapi.type+json", func() {
	Description("A card type")
	Attributes(func() {
		Attribute("id", String, "Unique type ID")
		Attribute("href", String, "API href for making requests on the type")
		Attribute("name", String, "Name of the type")

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

var RowMedia = MediaType("application/vnd.gwentapi.row+json", func() {
	Description("A row where a card can be played to")
	Attributes(func() {
		Attribute("id", String, "Unique row ID")
		Attribute("href", String, "API href for making requests on the row")
		Attribute("name", String, "Name of the row")

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

	View("noModel", func() {
		Attribute("name")
	})
})

var PatchMedia = MediaType("application/vnd.gwentapi.patch+json", func() {
	Description("A game patch")
	Attributes(func() {
		Attribute("id", String, "Unique Patch ID")
		Attribute("href", String, "API href for making requests on the patch")
		Attribute("version", String, "Patch version")
		Attribute("releaseDate", DateTime, "Release date of the patch")
		Attribute("changelog", String, "Patch changelog")

		Required("id", "href", "version", "releaseDate")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("version")
		Attribute("releaseDate")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("version")
		Attribute("releaseDate")
		Attribute("changelog")
	})
})