package helpers

import (
	"fmt"
	"github.com/GwentAPI/gwentapi/app"
	"github.com/GwentAPI/gwentapi/configuration"
)

var baseURL string

// CardHref returns the resource href.
func CardURL(cardID interface{}) string {
	return baseURL + app.CardHref(cardID)
}

// FactionHref returns the resource href.
func FactionURL(factionID interface{}) string {
	return baseURL + app.FactionHref(factionID)
}

func VariationURL(cardID interface{}, variationID interface{}) string {
	return fmt.Sprintf("%v/v0/cards/%v/variations/%v", baseURL, cardID, variationID)
}

func CategoryURL(categoryUUID interface{}) string {
	return baseURL + app.CategoryHref(categoryUUID)
}

// GroupHref returns the resource href.
func GroupURL(groupID interface{}) string {
	return baseURL + app.GroupHref(groupID)
}
func SwaggerURL() string {
	return "https://gwentapi.com/swagger/"
}

// RarityHref returns the resource href.
func RarityURL(rarityID interface{}) string {
	return baseURL + app.RarityHref(rarityID)
}

// TypeHref returns the resource href.
func MediaURL(filename interface{}) string {
	return fmt.Sprintf("%v/media/%v", baseURL, filename)
}

func init() {
	baseURL = configuration.GetConfig().App.BaseURL
}
