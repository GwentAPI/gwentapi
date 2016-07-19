//************************************************************************//
// API "gwentapi": Application Resource Href Factories
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/tri125/gwentapi/design
// --out=$(GOPATH)\src\github.com\tri125\gwentapi
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// CardHref returns the resource href.
func CardHref(cardID interface{}) string {
	return fmt.Sprintf("/v0/cards/%v", cardID)
}

// FactionHref returns the resource href.
func FactionHref(factionID interface{}) string {
	return fmt.Sprintf("/v0/factions/%v", factionID)
}

// GlyphHref returns the resource href.
func GlyphHref(glyphID interface{}) string {
	return fmt.Sprintf("/v0/glyphs/%v", glyphID)
}

// PatchHref returns the resource href.
func PatchHref(patchID interface{}) string {
	return fmt.Sprintf("/v0/patches/%v", patchID)
}

// PhonebookHref returns the resource href.
func PhonebookHref() string {
	return fmt.Sprintf("/v0")
}

// RarityHref returns the resource href.
func RarityHref(rarityID interface{}) string {
	return fmt.Sprintf("/v0/rarities/%v", rarityID)
}

// RowHref returns the resource href.
func RowHref(rowID interface{}) string {
	return fmt.Sprintf("/v0/rows/%v", rowID)
}

// TypeHref returns the resource href.
func TypeHref(typeID interface{}) string {
	return fmt.Sprintf("/v0/types/%v", typeID)
}
