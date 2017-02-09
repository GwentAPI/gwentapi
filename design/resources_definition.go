package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("phonebook", func() {
	DefaultMedia(ResourceMedia)
	BasePath("/v0")

	Action("show", func() {
		Description("Listing of all supported resources endpoint")
		Routing(GET(""))
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("faction", func() {

	DefaultMedia(FactionMedia)
	BasePath("/v0/factions")

	Response(InternalServerError)

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

var _ = Resource("rarity", func() {
	DefaultMedia(RarityMedia)
	BasePath("/v0/rarities")

	Response(InternalServerError)

	Action("list", func() {
		Routing(GET(""))
		Description("Return all rarities.")

		Response(OK, CollectionOf(RarityMedia))
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return rarity with given id.")
		Routing(GET("/:rarityID"))
		Params(func() {
			Param("rarityID", String, "Rarity ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("glyph", func() {
	DefaultMedia(GlyphMedia)
	BasePath("/v0/glyphs")

	Response(InternalServerError)

	Action("list", func() {
		Routing(GET(""))
		Description("Return all glyphs.")

		Response(OK, CollectionOf(GlyphMedia))
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return glyph with given id.")
		Routing(GET("/:glyphID"))
		Params(func() {
			Param("glyphID", String, "Glyph ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("type", func() {
	DefaultMedia(TypeMedia)
	BasePath("/v0/types")

	Response(InternalServerError)

	Action("list", func() {
		Routing(GET(""))
		Description("Return all card types.")

		Response(OK, CollectionOf(TypeMedia))
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return card type with given id.")
		Routing(GET("/:typeID"))
		Params(func() {
			Param("typeID", String, "Card type ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("patch", func() {
	DefaultMedia(PatchMedia)
	BasePath("/v0/patches")

	Response(InternalServerError)

	Action("list", func() {
		Routing(GET(""))
		Description("Return all patches.")

		Response(OK, CollectionOf(PatchMedia))
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return patch with given id.")
		Routing(GET("/:patchID"))
		Params(func() {
			Param("patchID", String, "Patch ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("latest", func() {
		Description("Return latest patch.")
		Routing(GET("/latest"))
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("card", func() {
	DefaultMedia(CardMedia)
	BasePath("/v0/cards")

	Response(InternalServerError)

	Params(func() {
		Param("limit", Integer, "Number of cards to receive")
		Param("offset", Integer, "Offset of the starting count")
	})

	Action("list", func() {
		Routing(GET(""))
		Description("Return a page of cards.")

		Response(OK, PageCard)
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return card with given id.")
		Routing(GET("/:cardID"))

		Params(func() {
			Param("cardID", String, "Card ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("cardFaction", func() {
		Description("Return all cards with given faction id.")
		Routing(GET("/factions/:factionID"))

		Params(func() {
			Param("factionID", String, "Faction ID")
		})
		Response(OK, PageCard)
		Response(NotFound)
	})

	Action("cardRarity", func() {
		Description("Return all cards with given rarity id.")
		Routing(GET("/rarities/:rarityID"))

		Params(func() {
			Param("rarityID", String, "Rarity ID")
		})
		Response(OK, PageCard)
		Response(NotFound)
	})

	Action("cardLeader", func() {
		Description("Return all leader cards.")
		Routing(GET("/leaders"))

		Response(OK, PageCard)
		Response(NotFound)
	})

	Action("cardArtworks", func() {
		Description("Return artwork with given id.")
		Routing(GET("/:cardID/artworks/"))

		Params(func() {
			Param("cardID", String, "Card ID")
		})
		Response(OK, VariationMedia)
		Response(NotFound)
	})
})
