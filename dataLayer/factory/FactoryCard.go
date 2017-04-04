package factory

import (
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/models"
	"github.com/tri125/gwentapi/helpers"
)

func CreateCard(c *models.Card, v *[]models.Variation) (*app.GwentapiCard, error) {
	uuid := helpers.UUIDToURLBase64(c.UUID)
	categories := make(app.GwentapiCategoryLinkCollection, len(c.Categories))
	variations := make(app.GwentapiVariationLinkCollection, len(*v))

	for i, category := range c.Categories {
		cl := &app.GwentapiCategoryLink{
			Name: category,
			Href: "",
		}
		categories[i] = cl
	}

	for i, variation := range *v {
		variationUUID := helpers.UUIDToURLBase64(variation.UUID)
		r := &app.GwentapiRarityLink{
			Name: variation.Rarity,
			Href: "",
		}

		cv := &app.GwentapiVariationLink{
			Availability: variation.Availability,
			Href:         helpers.VariationURL(uuid, variationUUID),
			Rarity:       r,
		}
		variations[i] = cv
	}

	faction := &app.GwentapiFactionLink{
		Href: "",
		Name: c.Faction,
	}

	group := &app.GwentapiGroupLink{
		Href: "",
		Name: c.Group,
	}

	result := &app.GwentapiCard{
		Name:       c.Name,
		Categories: categories,
		Flavor:     c.Flavor,
		Info:       c.Info,
		Strength:   c.Strength,
		Positions:  c.Positions,
		Faction:    faction,
		Group:      group,
		Variations: variations,
		Href:       helpers.CardURL(uuid),
		UUID:       uuid,
	}

	return result, nil
}

func CreatePageCard(c *[]models.Card, url string, resultCount int, limit int, offset int) (*app.GwentapiPagecard, error) {
	results := make(app.GwentapiCardLinkCollection, len(*c))
	for i, result := range *c {
		uuid := helpers.UUIDToURLBase64(result.UUID)
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
