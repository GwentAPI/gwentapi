package factory

import (
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/dal"
	"github.com/tri125/gwentapi/dataLayer/models"
	"github.com/tri125/gwentapi/helpers"
)

func CreateCard(c *models.Card, ds *dal.DataStore) (*app.GwentapiCard, error) {
	uuid := helpers.EncodeUUID(c.UUID)
	dalCat := dal.NewDalCategory(ds)
	dalV := dal.NewDalVariation(ds)
	dalG := dal.NewDalGroup(ds)
	dalF := dal.NewDalFaction(ds)
	dalR := dal.NewDalRarity(ds)

	variations, errV := dalV.FetchFromCardID(c.ID)
	if errV != nil {
		return nil, errV
	}
	categories, errC := dalCat.FetchFromArrayID(c.Categories_id)
	if errC != nil {
		return nil, errC
	}
	group, errG := dalG.FetchWithName(c.Group)
	if errG != nil {
		return nil, errG
	}
	faction, errF := dalF.FetchWithName(c.Faction)
	if errF != nil {
		return nil, errF
	}

	categoriesLinkMedia := make(app.GwentapiCategoryLinkCollection, len(*categories))
	variationsLinkMedia := make(app.GwentapiVariationLinkCollection, len(*variations))

	for i, category := range *categories {
		cl, _ := CreateCategoryLink(&category)
		categoriesLinkMedia[i] = cl
	}

	for i, variation := range *variations {
		rarity, errR := dalR.FetchWithName(variation.Rarity)
		if errR != nil {
			return nil, errR
		}

		r, _ := CreateRarityLink(rarity)

		variationUUID := helpers.EncodeUUID(variation.UUID)
		cv := &app.GwentapiVariationLink{
			Availability: variation.Availability,
			Href:         helpers.VariationURL(uuid, variationUUID),
			Rarity:       r,
		}
		variationsLinkMedia[i] = cv
	}

	factionMedia, _ := CreateFactionLink(faction)
	groupMedia, _ := CreateGroupLink(group)

	result := &app.GwentapiCard{
		Name:       c.Name,
		Categories: categoriesLinkMedia,
		Flavor:     c.Flavor,
		Info:       c.Info,
		Strength:   c.Strength,
		Positions:  c.Positions,
		Faction:    factionMedia,
		Group:      groupMedia,
		Variations: variationsLinkMedia,
		Href:       helpers.CardURL(uuid),
		UUID:       uuid,
	}

	return result, nil
}

func CreatePageCard(c *[]models.Card, url string, resultCount int, limit int, offset int) (*app.GwentapiPagecard, error) {
	results := make(app.GwentapiCardLinkCollection, len(*c))
	for i, result := range *c {
		uuid := helpers.EncodeUUID(result.UUID)
		cl := &app.GwentapiCardLink{
			Href: helpers.CardURL(uuid),
			Name: result.Name,
		}
		results[i] = cl
	}

	prev, next := helpers.GeneratePrevNextPageHref(resultCount, limit, offset, helpers.CardURL(url))

	page := &app.GwentapiPagecard{
		Count:    resultCount,
		Next:     next,
		Previous: prev,
		Results:  results,
	}
	return page, nil
}
