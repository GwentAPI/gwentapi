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
	Response(NotModified)

	Headers(func() {
		Header("If-Modified-Since", func() {
			Description("DateTime in RFC1123 format.")
			Example("Mon, 02 Jan 2006 15:04:05 GMT")
		})
	})

	Action("list", func() {
		Routing(GET(""))
		Description("Return all factions.")

		Response(OK, CollectionOf(FactionMedia, func() {
			ContentType("application/json; type=collection; charset=utf-8")
		}), func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
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
		Response(OK, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)

	})
})

var _ = Resource("rarity", func() {
	DefaultMedia(RarityMedia)
	BasePath("/v0/rarities")

	Response(InternalServerError)
	Response(NotModified)

	Headers(func() {
		Header("If-Modified-Since", func() {
			Description("DateTime in RFC1123 format.")
			Example("Mon, 02 Jan 2006 15:04:05 GMT")
		})
	})

	Action("list", func() {
		Routing(GET(""))
		Description("Return all rarities.")

		Response(OK, CollectionOf(RarityMedia, func() {
			ContentType("application/json; type=collection; charset=utf-8")
		}), func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return rarity with given id.")
		Routing(GET("/:rarityID"))
		Params(func() {
			Param("rarityID", String, "Rarity ID")
		})
		Response(OK, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})
})

var _ = Resource("group", func() {
	DefaultMedia(GroupMedia)
	BasePath("/v0/groups")

	Response(InternalServerError)
	Response(NotModified)

	Headers(func() {
		Header("If-Modified-Since", func() {
			Description("DateTime in RFC1123 format.")
			Example("Mon, 02 Jan 2006 15:04:05 GMT")
		})
	})

	Action("list", func() {
		Routing(GET(""))
		Description("Return all card groups.")

		Response(OK, CollectionOf(GroupMedia, func() {
			ContentType("application/json; type=collection; charset=utf-8")
		}), func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return card group with given id.")
		Routing(GET("/:groupID"))
		Params(func() {
			Param("groupID", String, "Card group ID")
		})
		Response(OK, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})
})

var _ = Resource("category", func() {
	DefaultMedia(CategoryMedia)
	BasePath("/v0/categories")

	Response(InternalServerError)
	Response(NotModified)

	Headers(func() {
		Header("If-Modified-Since", func() {
			Description("DateTime in RFC1123 format.")
			Example("Mon, 02 Jan 2006 15:04:05 GMT")
		})
	})

	Action("list", func() {
		Routing(GET(""))
		Description("Return all card categories.")

		Response(OK, CollectionOf(CategoryMedia, func() {
			ContentType("application/json; type=collection; charset=utf-8")
		}), func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return card category with given id.")
		Routing(GET("/:categoryID"))
		Params(func() {
			Param("categoryID", String, "Card category ID")
		})
		Response(OK, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})
})

var _ = Resource("card", func() {
	DefaultMedia(CardMedia)
	BasePath("/v0/cards")

	Response(InternalServerError)
	Response(NotModified)

	Headers(func() {
		Header("If-Modified-Since", func() {
			Description("DateTime in RFC1123 format.")
			Example("Mon, 02 Jan 2006 15:04:05 GMT")
		})
	})

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
				MaxLength(100)
			})

		})
		Response(OK, PageCard, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

	Action("show", func() {
		Description("Return card with given id.")
		Routing(GET("/:cardID"))

		Params(func() {
			Param("cardID", String, "Card ID")
		})
		Response(OK, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

	Action("cardFaction", func() {
		Description("Return all cards with given faction id.")
		Routing(GET("/factions/:factionID"))

		Params(func() {
			Param("factionID", String, "Faction ID")
		})
		Response(OK, PageCard, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

	Action("cardRarity", func() {
		Description("Return all cards with given rarity id.")
		Routing(GET("/rarities/:rarityID"))

		Params(func() {
			Param("rarityID", String, "Rarity ID")
		})
		Response(OK, PageCard, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

	Action("cardLeader", func() {
		Description("Return all leader cards.")
		Routing(GET("/leaders"))

		Response(OK, PageCard, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

	Action("cardVariations", func() {
		Description("Return the variations of a card with the given id.")
		Routing(GET("/:cardID/variations/"))

		Params(func() {
			Param("cardID", String, "Card ID")
		})
		Response(OK, CollectionOf(VariationMedia, func() {
			ContentType("application/json; type=collection; charset=utf-8")
		}), func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

	Action("cardVariation", func() {
		Description("Return the variation of a given id of a card with the given id.")
		Routing(GET("/:cardID/variations/:variationID"))

		Params(func() {
			Param("cardID", String, "Card ID")
			Param("variationID", String, "Variation ID")
		})
		Response(OK, VariationMedia, func() {
			Headers(func() {
				Header("Last-Modified", func() {
					Description("DateTime in RFC1123 format.")
					Example("Mon, 02 Jan 2006 15:04:05 GMT")
				})
			})
		})
		Response(NotFound)
	})

})
