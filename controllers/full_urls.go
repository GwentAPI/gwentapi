package controllers

import "github.com/tri125/gwentapi/app"

const host string = "http://localhost:8080"

// ArtworkHref returns the resource href.
func ArtworkURL(cardID interface{}) string {
	return host + app.ArtworkHref(cardID)
}

// CardHref returns the resource href.
func CardURL(cardID interface{}) string {
	return host + app.CardHref(cardID)
}

// FactionHref returns the resource href.
func FactionURL(factionID interface{}) string {
	return host + app.FactionHref(factionID)
}

// GlyphHref returns the resource href.
func GlyphURL(glyphID interface{}) string {
	return host + app.GlyphHref(glyphID)
}

// PatchHref returns the resource href.
func PatchURL(patchID interface{}) string {
	return host + app.PatchHref(patchID)
}

// PhonebookHref returns the resource href.
func PhonebookURL() string {
	return host
}

// RarityHref returns the resource href.
func RarityURL(rarityID interface{}) string {
	return host + app.RarityHref(rarityID)
}

// TypeHref returns the resource href.
func TypeURL(typeID interface{}) string {
	return host + app.TypeHref(typeID)
}

// TypeHref returns the resource href.
func MediaURL(cardID string) string {
	return host + "/media/" + cardID
}
