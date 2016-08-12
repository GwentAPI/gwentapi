package controllers

import (
	"fmt"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/configuration"
)

// ArtworkHref returns the resource href.
func ArtworkURL(cardID interface{}) string {
	return fmt.Sprintf("%v/v0/cards/%v/artworks", configuration.Hostname, cardID)
}

// CardHref returns the resource href.
func CardURL(cardID interface{}) string {
	return configuration.Hostname + app.CardHref(cardID)
}

// FactionHref returns the resource href.
func FactionURL(factionID interface{}) string {
	return configuration.Hostname + app.FactionHref(factionID)
}

// GlyphHref returns the resource href.
func GlyphURL(glyphID interface{}) string {
	return configuration.Hostname + app.GlyphHref(glyphID)
}

// PatchHref returns the resource href.
func PatchURL(patchID interface{}) string {
	return configuration.Hostname + app.PatchHref(patchID)
}

// PhonebookHref returns the resource href.
func PhonebookURL() string {
	return configuration.Hostname
}

// RarityHref returns the resource href.
func RarityURL(rarityID interface{}) string {
	return configuration.Hostname + app.RarityHref(rarityID)
}

// TypeHref returns the resource href.
func TypeURL(typeID interface{}) string {
	return configuration.Hostname + app.TypeHref(typeID)
}

// TypeHref returns the resource href.
func MediaURL(filename string) string {
	return configuration.Hostname + "/media/" + filename
}
