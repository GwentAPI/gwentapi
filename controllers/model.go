package controllers

import (
	"github.com/tri125/gwentapi/app"
	"time"
)

type TypeModel struct {
	Name string
	ID   string
}

type RarityModel struct {
	Name string
	ID   string
}

type PatchModel struct {
	Version     string
	ID          string
	ReleaseDate time.Time
	Changelog   *string
}

type GlyphModel struct {
	Name           string
	ID             string
	IsWeatherGlyph bool
	Text           string
}

type FactionModel struct {
	Name string
	ID   string
}

type IllustratorModel struct {
	Name string
	ID   string
}

type CardModel struct {
	Name     string
	ID       string
	Rarity   RarityModel
	Faction  FactionModel
	TypeCard TypeModel //as to not conflict with the reserved type keyword
	Subtypes []*TypeModel
	Rows     []string
	Strength *int
	Text     *string
	Flavor   *string
}

type ArtworkMediaModel struct {
	ID           string
	Alternatives []*app.ArtworkType
	Artwork      *app.ArtworkType
}
