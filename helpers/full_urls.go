package helpers

import (
	"fmt"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/configuration"
)

// ArtworkHref returns the resource href.
func ArtworkURL(cardID interface{}) string {
	return fmt.Sprintf("%v/v0/cards/%v/artworks", configuration.Conf.Server.BaseURL, cardID)
}

// CardHref returns the resource href.
func CardURL(cardID interface{}) string {
	return configuration.Conf.Server.BaseURL + app.CardHref(cardID)
}

// FactionHref returns the resource href.
func FactionURL(factionID interface{}) string {
	return configuration.Conf.Server.BaseURL + app.FactionHref(factionID)
}

func VariationURL(cardID interface{}, variationID interface{}) string {
	return fmt.Sprintf("%v/v0/cards/%v/variations/%v", configuration.Conf.Server.BaseURL, cardID, variationID)
}

func CategoryURL(categoryUUID interface{}) string {
	return configuration.Conf.Server.BaseURL + app.CategoryHref(categoryUUID)
}

// GroupHref returns the resource href.
func GroupURL(groupID interface{}) string {
	return configuration.Conf.Server.BaseURL + app.GroupHref(groupID)
}
func SwaggerURL() string {
	return configuration.Conf.Server.BaseURL + "/swagger/index.html"
}

// RarityHref returns the resource href.
func RarityURL(rarityID interface{}) string {
	return configuration.Conf.Server.BaseURL + app.RarityHref(rarityID)
}

// TypeHref returns the resource href.
func MediaURL(filename string) string {
	return configuration.Conf.Server.BaseURL + "/media/" + filename
}
