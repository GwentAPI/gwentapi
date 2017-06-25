// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "gwentapi": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/GwentAPI/gwentapi/design
// --out=$(GOPATH)\src\github.com\GwentAPI\gwentapi
// --version=v1.2.0

package app

import (
	"github.com/goadesign/goa"
)

// A card (default view)
//
// Identifier: application/vnd.gwentapi.card+json; view=default
type GwentapiCard struct {
	// Categories of the card
	Categories GwentapiCategoryLinkCollection `form:"categories,omitempty" json:"categories,omitempty" xml:"categories,omitempty"`
	// Faction of the card
	Faction *GwentapiFactionLink `form:"faction" json:"faction" xml:"faction"`
	// Flavor text of the card
	Flavor *string `form:"flavor,omitempty" json:"flavor,omitempty" xml:"flavor,omitempty"`
	// Group of the card
	Group *GwentapiGroupLink `form:"group" json:"group" xml:"group"`
	// API href for making requests on the card
	Href string `form:"href" json:"href" xml:"href"`
	// Text of the card detailing its abilities and how it plays
	Info *string `form:"info,omitempty" json:"info,omitempty" xml:"info,omitempty"`
	// Name of the card
	Name string `form:"name" json:"name" xml:"name"`
	// Positions where the card can be played
	Positions []string `form:"positions,omitempty" json:"positions,omitempty" xml:"positions,omitempty"`
	// Strength of the card
	Strength *int `form:"strength,omitempty" json:"strength,omitempty" xml:"strength,omitempty"`
	// Unique card UUID
	UUID string `form:"uuid" json:"uuid" xml:"uuid"`
	// Variations of the card
	Variations GwentapiVariationLinkCollection `form:"variations" json:"variations" xml:"variations"`
}

// Validate validates the GwentapiCard media type instance.
func (mt *GwentapiCard) Validate() (err error) {
	if mt.UUID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "uuid"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Group == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "group"))
	}
	if mt.Faction == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "faction"))
	}
	if mt.Variations == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "variations"))
	}
	if err2 := mt.Categories.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if mt.Faction != nil {
		if err2 := mt.Faction.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Group != nil {
		if err2 := mt.Group.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	if err2 := mt.Variations.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// A card (link view)
//
// Identifier: application/vnd.gwentapi.card+json; view=link
type GwentapiCardLink struct {
	// API href for making requests on the card
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the card
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiCardLink media type instance.
func (mt *GwentapiCardLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	return
}

// GwentapiCardCollection is the media type for an array of GwentapiCard (default view)
//
// Identifier: application/vnd.gwentapi.card+json; type=collection; view=default
type GwentapiCardCollection []*GwentapiCard

// Validate validates the GwentapiCardCollection media type instance.
func (mt GwentapiCardCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// GwentapiCardCollection is the media type for an array of GwentapiCard (link view)
//
// Identifier: application/vnd.gwentapi.card+json; type=collection; view=link
type GwentapiCardLinkCollection []*GwentapiCardLink

// Validate validates the GwentapiCardLinkCollection media type instance.
func (mt GwentapiCardLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A card category (default view)
//
// Identifier: application/vnd.gwentapi.category+json; view=default
type GwentapiCategory struct {
	// API href for making requests on the category
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the category
	Name string `form:"name" json:"name" xml:"name"`
	// Unique category UUID
	UUID string `form:"uuid" json:"uuid" xml:"uuid"`
}

// Validate validates the GwentapiCategory media type instance.
func (mt *GwentapiCategory) Validate() (err error) {
	if mt.UUID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "uuid"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	return
}

// A card category (link view)
//
// Identifier: application/vnd.gwentapi.category+json; view=link
type GwentapiCategoryLink struct {
	// API href for making requests on the category
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the category
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiCategoryLink media type instance.
func (mt *GwentapiCategoryLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	return
}

// GwentapiCategoryCollection is the media type for an array of GwentapiCategory (default view)
//
// Identifier: application/vnd.gwentapi.category+json; type=collection; view=default
type GwentapiCategoryCollection []*GwentapiCategory

// Validate validates the GwentapiCategoryCollection media type instance.
func (mt GwentapiCategoryCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// GwentapiCategoryCollection is the media type for an array of GwentapiCategory (link view)
//
// Identifier: application/vnd.gwentapi.category+json; type=collection; view=link
type GwentapiCategoryLinkCollection []*GwentapiCategoryLink

// Validate validates the GwentapiCategoryLinkCollection media type instance.
func (mt GwentapiCategoryLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A card faction (default view)
//
// Identifier: application/vnd.gwentapi.faction+json; view=default
type GwentapiFaction struct {
	// API href for making requests on the faction
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the faction
	Name string `form:"name" json:"name" xml:"name"`
	// Unique faction UUID
	UUID string `form:"uuid" json:"uuid" xml:"uuid"`
}

// Validate validates the GwentapiFaction media type instance.
func (mt *GwentapiFaction) Validate() (err error) {
	if mt.UUID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "uuid"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	return
}

// A card faction (link view)
//
// Identifier: application/vnd.gwentapi.faction+json; view=link
type GwentapiFactionLink struct {
	// API href for making requests on the faction
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the faction
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiFactionLink media type instance.
func (mt *GwentapiFactionLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	return
}

// GwentapiFactionCollection is the media type for an array of GwentapiFaction (default view)
//
// Identifier: application/vnd.gwentapi.faction+json; type=collection; view=default
type GwentapiFactionCollection []*GwentapiFaction

// Validate validates the GwentapiFactionCollection media type instance.
func (mt GwentapiFactionCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// GwentapiFactionCollection is the media type for an array of GwentapiFaction (link view)
//
// Identifier: application/vnd.gwentapi.faction+json; type=collection; view=link
type GwentapiFactionLinkCollection []*GwentapiFactionLink

// Validate validates the GwentapiFactionLinkCollection media type instance.
func (mt GwentapiFactionLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A card group (default view)
//
// Identifier: application/vnd.gwentapi.group+json; view=default
type GwentapiGroup struct {
	// API href for making requests on the group
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the group
	Name string `form:"name" json:"name" xml:"name"`
	// Unique group UUID
	UUID string `form:"uuid" json:"uuid" xml:"uuid"`
}

// Validate validates the GwentapiGroup media type instance.
func (mt *GwentapiGroup) Validate() (err error) {
	if mt.UUID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "uuid"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	return
}

// A card group (link view)
//
// Identifier: application/vnd.gwentapi.group+json; view=link
type GwentapiGroupLink struct {
	// API href for making requests on the group
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the group
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiGroupLink media type instance.
func (mt *GwentapiGroupLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	return
}

// GwentapiGroupCollection is the media type for an array of GwentapiGroup (default view)
//
// Identifier: application/vnd.gwentapi.group+json; type=collection; view=default
type GwentapiGroupCollection []*GwentapiGroup

// Validate validates the GwentapiGroupCollection media type instance.
func (mt GwentapiGroupCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// GwentapiGroupCollection is the media type for an array of GwentapiGroup (link view)
//
// Identifier: application/vnd.gwentapi.group+json; type=collection; view=link
type GwentapiGroupLinkCollection []*GwentapiGroupLink

// Validate validates the GwentapiGroupLinkCollection media type instance.
func (mt GwentapiGroupLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Paginated card (default view)
//
// Identifier: application/vnd.gwentapi.pagecard+json; view=default
type GwentapiPagecard struct {
	// Total number of cards stored in the database
	Count int `form:"count" json:"count" xml:"count"`
	// Href to the next page
	Next *string `form:"next,omitempty" json:"next,omitempty" xml:"next,omitempty"`
	// Href to the previous page
	Previous *string `form:"previous,omitempty" json:"previous,omitempty" xml:"previous,omitempty"`
	// Results of the page containing cards
	Results GwentapiCardLinkCollection `form:"results" json:"results" xml:"results"`
}

// Validate validates the GwentapiPagecard media type instance.
func (mt *GwentapiPagecard) Validate() (err error) {

	if mt.Results == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "results"))
	}
	if mt.Next != nil {
		if err2 := goa.ValidateFormat(goa.FormatURI, *mt.Next); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.next`, *mt.Next, goa.FormatURI, err2))
		}
	}
	if mt.Previous != nil {
		if err2 := goa.ValidateFormat(goa.FormatURI, *mt.Previous); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.previous`, *mt.Previous, goa.FormatURI, err2))
		}
	}
	if err2 := mt.Results.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// A card rarity (default view)
//
// Identifier: application/vnd.gwentapi.rarity+json; view=default
type GwentapiRarity struct {
	// API href for making requests on the rarity
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the rarity
	Name string `form:"name" json:"name" xml:"name"`
	// Unique rarity UUID
	UUID string `form:"uuid" json:"uuid" xml:"uuid"`
}

// Validate validates the GwentapiRarity media type instance.
func (mt *GwentapiRarity) Validate() (err error) {
	if mt.UUID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "uuid"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	return
}

// A card rarity (link view)
//
// Identifier: application/vnd.gwentapi.rarity+json; view=link
type GwentapiRarityLink struct {
	// API href for making requests on the rarity
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the rarity
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiRarityLink media type instance.
func (mt *GwentapiRarityLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	return
}

// GwentapiRarityCollection is the media type for an array of GwentapiRarity (default view)
//
// Identifier: application/vnd.gwentapi.rarity+json; type=collection; view=default
type GwentapiRarityCollection []*GwentapiRarity

// Validate validates the GwentapiRarityCollection media type instance.
func (mt GwentapiRarityCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// GwentapiRarityCollection is the media type for an array of GwentapiRarity (link view)
//
// Identifier: application/vnd.gwentapi.rarity+json; type=collection; view=link
type GwentapiRarityLinkCollection []*GwentapiRarityLink

// Validate validates the GwentapiRarityLinkCollection media type instance.
func (mt GwentapiRarityLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Listing of all available resource endpoint and a link to the api definition (default view)
//
// Identifier: application/vnd.gwentapi.resource+json; view=default
type GwentapiResource struct {
	// API href for making requests on cards
	Cards string `form:"cards" json:"cards" xml:"cards"`
	// API href for making requests on categories
	Categories string `form:"categories" json:"categories" xml:"categories"`
	// API href for making requests on factions
	Factions string `form:"factions" json:"factions" xml:"factions"`
	// API href for making requests on groups
	Groups string `form:"groups" json:"groups" xml:"groups"`
	// API href for making requests on rarities
	Rarities string `form:"rarities" json:"rarities" xml:"rarities"`
	// Href linking to the swagger definition
	Swagger string `form:"swagger" json:"swagger" xml:"swagger"`
	// Semver of the software that is currently running
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the GwentapiResource media type instance.
func (mt *GwentapiResource) Validate() (err error) {
	if mt.Cards == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "cards"))
	}
	if mt.Factions == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "factions"))
	}
	if mt.Rarities == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rarities"))
	}
	if mt.Categories == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "categories"))
	}
	if mt.Groups == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "groups"))
	}
	if mt.Swagger == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "swagger"))
	}
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Cards); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.cards`, mt.Cards, goa.FormatURI, err2))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Categories); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.categories`, mt.Categories, goa.FormatURI, err2))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Factions); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.factions`, mt.Factions, goa.FormatURI, err2))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Groups); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.groups`, mt.Groups, goa.FormatURI, err2))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Rarities); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.rarities`, mt.Rarities, goa.FormatURI, err2))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Swagger); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.swagger`, mt.Swagger, goa.FormatURI, err2))
	}
	return
}

// Variation of a card containing artworks, crafting/milling cost, set availability, and rarity. (default view)
//
// Identifier: application/vnd.gwentapi.variation+json; view=default
type GwentapiVariation struct {
	// Artworks of the card variation.
	Art *ArtType `form:"art,omitempty" json:"art,omitempty" xml:"art,omitempty"`
	// Describe from which set the variation comes from and its general availability
	Availability string `form:"availability" json:"availability" xml:"availability"`
	// Crafting cost of the card variation.
	Craft *CostType `form:"craft,omitempty" json:"craft,omitempty" xml:"craft,omitempty"`
	// API href for making requests on the artwork
	Href string `form:"href" json:"href" xml:"href"`
	// Milling cost of the card variation.
	Mill *CostType `form:"mill,omitempty" json:"mill,omitempty" xml:"mill,omitempty"`
	// Rarity of the card
	Rarity *GwentapiRarityLink `form:"rarity" json:"rarity" xml:"rarity"`
	// Unique artwork UUID
	UUID string `form:"uuid" json:"uuid" xml:"uuid"`
}

// Validate validates the GwentapiVariation media type instance.
func (mt *GwentapiVariation) Validate() (err error) {
	if mt.UUID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "uuid"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Availability == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "availability"))
	}
	if mt.Rarity == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rarity"))
	}
	if mt.Art != nil {
		if err2 := mt.Art.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	if mt.Rarity != nil {
		if err2 := mt.Rarity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Variation of a card containing artworks, crafting/milling cost, set availability, and rarity. (link view)
//
// Identifier: application/vnd.gwentapi.variation+json; view=link
type GwentapiVariationLink struct {
	// Describe from which set the variation comes from and its general availability
	Availability string `form:"availability" json:"availability" xml:"availability"`
	// API href for making requests on the artwork
	Href string `form:"href" json:"href" xml:"href"`
	// Rarity of the card
	Rarity *GwentapiRarityLink `form:"rarity" json:"rarity" xml:"rarity"`
}

// Validate validates the GwentapiVariationLink media type instance.
func (mt *GwentapiVariationLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Availability == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "availability"))
	}
	if mt.Rarity == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rarity"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Href); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.href`, mt.Href, goa.FormatURI, err2))
	}
	if mt.Rarity != nil {
		if err2 := mt.Rarity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// GwentapiVariationCollection is the media type for an array of GwentapiVariation (default view)
//
// Identifier: application/vnd.gwentapi.variation+json; type=collection; view=default
type GwentapiVariationCollection []*GwentapiVariation

// Validate validates the GwentapiVariationCollection media type instance.
func (mt GwentapiVariationCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// GwentapiVariationCollection is the media type for an array of GwentapiVariation (link view)
//
// Identifier: application/vnd.gwentapi.variation+json; type=collection; view=link
type GwentapiVariationLinkCollection []*GwentapiVariationLink

// Validate validates the GwentapiVariationLinkCollection media type instance.
func (mt GwentapiVariationLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
