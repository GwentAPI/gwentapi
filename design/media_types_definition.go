package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ResourceMedia = MediaType("application/vnd.gwentapi.resource+json", func() {
	Description("Listing of all available resource endpoint and a link to the api definition")
	Attributes(func() {
		Attribute("cards", String, "API href for making requests on cards")
		Attribute("factions", String, "API href for making requests on factions")
		Attribute("rarities", String, "API href for making requests on rarities")
		Attribute("categories", String, "API href for making requests on categories")
		Attribute("groups", String, "API href for making requests on groups")
		Attribute("swagger", String, "Href linking to the swagger definition")

		Required("cards", "factions", "rarities", "categories", "groups", "swagger")
	})

	View("default", func() {
		Attribute("cards")
		Attribute("factions")
		Attribute("rarities")
		Attribute("categories")
		Attribute("groups")
		Attribute("swagger")
	})
})

var PageCard = MediaType("application/vnd.gwentapi.pageCard+json", func() {
	Description("Paginated card")
	Attributes(func() {
		Attribute("count", Integer, "Total number of cards stored in the database")
		Attribute("previous", String, "Href to the previous page")
		Attribute("next", String, "Href to the next page")
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
		Attribute("uuid", String, "Unique artwork UUID")
		Attribute("href", String, "API href for making requests on the artwork")
		//Attribute("card", CardMedia, "Card referred to by the artwork")
		Attribute("art", ArtType, "Artworks of the card variation.")
		Attribute("mill", CostType, "Milling cost of the card variation.")
		Attribute("craft", CostType, "Crafting cost of the card variation.")
		Attribute("availability", String, "Describe from which set the variation comes from and its general availability")
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
		Attribute("uuid", String, "Unique card UUID")
		Attribute("href", String, "API href for making requests on the card")
		Attribute("name", String, "Name of the card")
		Attribute("group", GroupMedia, "Group of the card")
		Attribute("faction", FactionMedia, "Faction of the card")

		Attribute("categories", CollectionOf(CategoryMedia), "Categories of the card")
		Attribute("positions", ArrayOf(String), "Positions where the card can be played")

		Attribute("strength", Integer, "Strength of the card")
		Attribute("info", String, "Text of the card detailing its abilities and how it plays")
		Attribute("flavor", String, "Flavor text of the card")
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
		Attribute("uuid", String, "Unique faction UUID")
		Attribute("href", String, "API href for making requests on the faction")
		Attribute("name", String, "Name of the faction")

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
		Attribute("uuid", String, "Unique rarity UUID")
		Attribute("href", String, "API href for making requests on the rarity")
		Attribute("name", String, "Name of the rarity")

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
		Attribute("uuid", String, "Unique group UUID")
		Attribute("href", String, "API href for making requests on the group")
		Attribute("name", String, "Name of the group")

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
		Attribute("uuid", String, "Unique category UUID")
		Attribute("href", String, "API href for making requests on the category")
		Attribute("name", String, "Name of the category")

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
