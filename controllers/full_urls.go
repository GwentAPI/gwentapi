package controllers

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

// GlyphHref returns the resource href.
func GlyphURL(glyphID interface{}) string {
	return configuration.Conf.Server.BaseURL + app.GlyphHref(glyphID)
}

// PatchHref returns the resource href.
func PatchURL(patchID interface{}) string {
	return configuration.Conf.Server.BaseURL + app.PatchHref(patchID)
}

// PhonebookHref returns the resource href.
func PhonebookURL() string {
	return configuration.Conf.Server.BaseURL
}

// RarityHref returns the resource href.
func RarityURL(rarityID interface{}) string {
	return configuration.Conf.Server.BaseURL + app.RarityHref(rarityID)
}

// TypeHref returns the resource href.
func TypeURL(typeID interface{}) string {
	return configuration.Conf.Server.BaseURL + app.TypeHref(typeID)
}

// TypeHref returns the resource href.
func MediaURL(filename string) string {
	return configuration.Conf.Server.BaseURL + "/media/" + filename
}
