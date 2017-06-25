package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ResourceMedia = MediaType("application/vnd.gwentapi.resource+json", func() {
	Description("Listing of all available resource endpoint and a link to the api definition")
	Attributes(func() {
		Attribute("cards", String, "API href for making requests on cards", func() {
			Format("uri")
		})
		Attribute("factions", String, "API href for making requests on factions", func() {
			Format("uri")
		})
		Attribute("rarities", String, "API href for making requests on rarities", func() {
			Format("uri")
		})
		Attribute("categories", String, "API href for making requests on categories", func() {
			Format("uri")
		})
		Attribute("groups", String, "API href for making requests on groups", func() {
			Format("uri")
		})
		Attribute("swagger", String, "Href linking to the swagger definition", func() {
			Format("uri")
		})
		Attribute("version", String, "Semver of the software that is currently running", func() {
			Example("v0.5.5-rc.2")
		})

		Required("cards", "factions", "rarities", "categories", "groups", "swagger", "version")
	})

	View("default", func() {
		Attribute("cards")
		Attribute("factions")
		Attribute("rarities")
		Attribute("categories")
		Attribute("groups")
		Attribute("swagger")
		Attribute("version")
	})
})

var PageCard = MediaType("application/vnd.gwentapi.pageCard+json", func() {
	Description("Paginated card")
	Attributes(func() {
		Attribute("count", Integer, "Total number of cards stored in the database", func() {
			Example(280)
		})
		Attribute("previous", String, "Href to the previous page", func() {
			Format("uri")
		})
		Attribute("next", String, "Href to the next page", func() {
			Format("uri")
		})
		Attribute("results", CollectionOf(CardMedia), "Results of the page containing cards")

		Required("count", "results")
	})

	Links(func() {
		Link("results")
	})

	View("default", func() {
		Attribute("count")
		Attribute("previous")
		Attribute("next")
		Attribute("results", func() {
			View("link")
		})
	})
})

var VariationMedia = MediaType("application/vnd.gwentapi.variation+json", func() {
	Description("Variation of a card containing artworks, crafting/milling cost, set availability, and rarity.")
	Attributes(func() {
		Attribute("uuid", String, "Unique artwork UUID", func() {
			Example("pcN4QMTlTAaIOwicgNwtKA")
		})
		Attribute("href", String, "API href for making requests on the artwork", func() {
			Format("uri")
		})
		Attribute("art", ArtType, "Artworks of the card variation.")
		Attribute("mill", CostType, "Milling cost of the card variation.")
		Attribute("craft", CostType, "Crafting cost of the card variation.")
		Attribute("availability", String, "Describe from which set the variation comes from and its general availability", func() {
			Example("BaseSet")
		})
		Attribute("rarity", RarityMedia, "Rarity of the card")

		Required("uuid", "href", "availability", "rarity")
	})

	//Links(func() {
	//	Link("card")
	//})

	View("default", func() {
		Attribute("uuid")
		Attribute("href")
		Attribute("art")
		Attribute("mill")
		Attribute("craft")
		Attribute("availability")
		Attribute("rarity", func() {
			View("link")
		})
	})

	View("link", func() {
		Attribute("href")
		Attribute("availability")
		Attribute("rarity", func() {
			View("link")
		})
	})
})

var CardMedia = MediaType("application/vnd.gwentapi.card+json", func() {
	Description("A card")
	Attributes(func() {
		Attribute("uuid", String, "Unique card UUID", func() {
			Example("oe6UPiaDSNyI-630fYz4LA")
		})
		Attribute("href", String, "API href for making requests on the card", func() {
			Format("uri")
		})
		Attribute("name", String, "Name of the card", func() {
			Example("Arachas")
		})
		Attribute("group", GroupMedia, "Group of the card")
		Attribute("faction", FactionMedia, "Faction of the card")

		Attribute("categories", CollectionOf(CategoryMedia), "Categories of the card")
		Attribute("positions", ArrayOf(String), "Positions where the card can be played", func() {
			Example([...]string{"Ranged"})
		})

		Attribute("strength", Integer, "Strength of the card", func() {
			Example(3)
		})
		Attribute("info", String, "Text of the card detailing its abilities and how it plays", func() {
			Example("Deploy: Play all Arachasae from your Deck.")
		})
		Attribute("flavor", String, "Flavor text of the card", func() {
			Example("Ugly – Nature's Way of saying stay away.")
		})
		Attribute("variations", CollectionOf(VariationMedia), "Variations of the card")

		Required("uuid", "href", "name", "group", "faction", "variations")
	})

	Links(func() {
		Link("group")
		Link("faction")
		Link("categories")
		Link("variations")
	})

	View("default", func() {
		Attribute("uuid")
		Attribute("href")
		Attribute("name")
		Attribute("group", func() {
			View("link")
		})
		Attribute("faction", func() {
			View("link")
		})
		Attribute("categories", func() {
			View("link")
		})
		Attribute("positions")
		Attribute("strength")
		Attribute("info")
		Attribute("flavor")
		Attribute("variations", func() {
			View("link")
		})
	})

	View("link", func() {
		Attribute("name")
		Attribute("href")
	})

})

var FactionMedia = MediaType("application/vnd.gwentapi.faction+json", func() {
	Description("A card faction")
	Attributes(func() {
		Attribute("uuid", String, "Unique faction UUID", func() {
			Example("9wM9vGWiRzCvEwSLnLfY1w")
		})
		Attribute("href", String, "API href for making requests on the faction", func() {
			Format("uri")
		})
		Attribute("name", String, "Name of the faction", func() {
			Example("Monster")
		})

		Required("uuid", "href", "name")
	})
	View("default", func() {
		Attribute("uuid")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("name")
		Attribute("href")
	})
})

var RarityMedia = MediaType("application/vnd.gwentapi.rarity+json", func() {
	Description("A card rarity")
	Attributes(func() {
		Attribute("uuid", String, "Unique rarity UUID", func() {
			Example("TPCvIPOjRjO0s7Jfeo1NtA")
		})
		Attribute("href", String, "API href for making requests on the rarity", func() {
			Format("uri")
		})
		Attribute("name", String, "Name of the rarity", func() {
			Example("Common")
		})

		Required("uuid", "href", "name")
	})

	View("default", func() {
		Attribute("uuid")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("name")
		Attribute("href")
	})
})

var GroupMedia = MediaType("application/vnd.gwentapi.group+json", func() {
	Description("A card group")
	Attributes(func() {
		Attribute("uuid", String, "Unique group UUID", func() {
			Example("GbmwHbkcQniDKJq6rKz-bQ")
		})
		Attribute("href", String, "API href for making requests on the group", func() {
			Format("uri")
		})
		Attribute("name", String, "Name of the group", func() {
			Example("Bronze")
		})

		Required("uuid", "href", "name")
	})

	View("default", func() {
		Attribute("uuid")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("name")
		Attribute("href")
	})
})

var CategoryMedia = MediaType("application/vnd.gwentapi.category+json", func() {
	Description("A card category")
	Attributes(func() {
		Attribute("uuid", String, "Unique category UUID", func() {
			Example("0PcjdpZ6QR2NKutLFGx-oQ")
		})
		Attribute("href", String, "API href for making requests on the category", func() {
			Format("uri")
		})
		Attribute("name", String, "Name of the category", func() {
			Example("Insectoid")
		})

		Required("uuid", "href", "name")
	})

	View("default", func() {
		Attribute("uuid")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("name")
		Attribute("href")
	})
})
