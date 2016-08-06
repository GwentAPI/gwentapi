package controllers

import (
	"github.com/tri125/gwentapi/app"
)

const host string = "http://localhost:8080"

// ArtworkHref returns the resource href.
func ArtworkURL(cardID interface{}) string {
	return host + ArtworkHref(cardID)
}

// CardHref returns the resource href.
func CardURL(cardID interface{}) string {
	return host + CardHref(cardID)
}

// FactionHref returns the resource href.
func FactionURL(factionID interface{}) string {
	return host + FactionHref(factionID)
}

// GlyphHref returns the resource href.
func GlyphURL(glyphID interface{}) string {
	return host + GlyphHref(glyphID)
}

// PatchHref returns the resource href.
func PatchURL(patchID interface{}) string {
	return host + PatchHref(patchID)
}

// PhonebookHref returns the resource href.
func PhonebookURL() string {
	return host
}

// RarityHref returns the resource href.
func RarityURL(rarityID interface{}) string {
	return host + RarityHref(rarityID)
}

// TypeHref returns the resource href.
func TypeURL(typeID interface{}) string {
	return host + TypeHref(typeID)
}
