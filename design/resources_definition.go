package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("index", func() {
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

		Response(OK, CollectionOf(FactionMedia), func() {
			Headers(func() {
				Header("Last-Modified", String)
				Required("Last-Modified")
			})
		})
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

var _ = Resource("group", func() {
	DefaultMedia(GroupMedia)
	BasePath("/v0/groups")

	Response(InternalServerError)

	Action("list", func() {
		Routing(GET(""))
		Description("Return all card groups.")

		Response(OK, CollectionOf(GroupMedia))
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return card group with given id.")
		Routing(GET("/:groupID"))
		Params(func() {
			Param("groupID", String, "Card group ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("category", func() {
	DefaultMedia(CategoryMedia)
	BasePath("/v0/categories")

	Response(InternalServerError)

	Action("list", func() {
		Routing(GET(""))
		Description("Return all card categories.")

		Response(OK, CollectionOf(CategoryMedia))
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return card category with given id.")
		Routing(GET("/:categoryID"))
		Params(func() {
			Param("categoryID", String, "Card category ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("card", func() {
	DefaultMedia(CardMedia)
	BasePath("/v0/cards")

	Response(InternalServerError)

	Params(func() {
		Param("limit", Integer, func() {
			Description("Number of cards to receive")
			Minimum(1)
			Default(20)
		})
		Param("offset", Integer, func() {
			Description("Offset of the starting count")
			Default(0)
			Minimum(0)
		})
	})

	Action("list", func() {
		Routing(GET(""))
		Description("Return a page of cards.")
		Params(func() {
			Param("name", String, func() {
				Description("Query to search for cards with the name starting by the entered value")
				MinLength(3)
				MaxLength(10)
			})

		})
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

	/*Action("cardRarity", func() {
		Description("Return all cards with given rarity id.")
		Routing(GET("/rarities/:rarityID"))

		Params(func() {
			Param("rarityID", String, "Rarity ID")
		})
		Response(OK, PageCard)
		Response(NotFound)
	})*/

	Action("cardLeader", func() {
		Description("Return all leader cards.")
		Routing(GET("/leaders"))

		Response(OK, PageCard)
		Response(NotFound)
	})

	Action("cardVariations", func() {
		Description("Return the variations of a card with the given id.")
		Routing(GET("/:cardID/variations/"))

		Params(func() {
			Param("cardID", String, "Card ID")
		})
		Response(OK, CollectionOf(VariationMedia))
		Response(NotFound)
	})

	Action("cardVariation", func() {
		Description("Return the variation of a given id of a card with the given id.")
		Routing(GET("/:cardID/variations/:variationID"))

		Params(func() {
			Param("cardID", String, "Card ID")
			Param("variationID", String, "Variation ID")
		})
		Response(OK, VariationMedia)
		Response(NotFound)
	})

})
