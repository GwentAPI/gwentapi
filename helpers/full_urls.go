package helpers

import (
	"fmt"
	"github.com/GwentAPI/gwentapi/app"
	"github.com/GwentAPI/gwentapi/configuration"
)

var publicURL string

// CardHref returns the resource href.
func CardURL(cardID interface{}) string {
	return publicURL + app.CardHref(cardID)
}

// FactionHref returns the resource href.
func FactionURL(factionID interface{}) string {
	return publicURL + app.FactionHref(factionID)
}

func VariationURL(cardID interface{}, variationID interface{}) string {
	return fmt.Sprintf("%v/v0/cards/%v/variations/%v", publicURL, cardID, variationID)
}

func CategoryURL(categoryUUID interface{}) string {
	return publicURL + app.CategoryHref(categoryUUID)
}

// GroupHref returns the resource href.
func GroupURL(groupID interface{}) string {
	return publicURL + app.GroupHref(groupID)
}
func SwaggerURL() string {
	return "https://gwentapi.com/swagger/"
}

// RarityHref returns the resource href.
func RarityURL(rarityID interface{}) string {
	return publicURL + app.RarityHref(rarityID)
}

// TypeHref returns the resource href.
func MediaURL(filename interface{}) string {
	return fmt.Sprintf("%v/media/%v", publicURL, filename)
}

func init() {
	publicURL = configuration.GetConfig().App.PublicURL
}
