//************************************************************************//
// API "gwentapi": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/tri125/gwentapi/design
// --out=$(GOPATH)\src\github.com\tri125\gwentapi
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ArtworkController is the controller interface for the Artwork actions.
type ArtworkController interface {
	goa.Muxer
	Show(*ShowArtworkContext) error
}

// MountArtworkController "mounts" a Artwork resource controller on the given service.
func MountArtworkController(service *goa.Service, ctrl ArtworkController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowArtworkContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/artworks/:cardID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Artwork", "action", "Show", "route", "GET /v0/artworks/:cardID")
}

// CardController is the controller interface for the Card actions.
type CardController interface {
	goa.Muxer
	CardFaction(*CardFactionCardContext) error
	CardLeader(*CardLeaderCardContext) error
	CardRarity(*CardRarityCardContext) error
	List(*ListCardContext) error
	Show(*ShowCardContext) error
}

// MountCardController "mounts" a Card resource controller on the given service.
func MountCardController(service *goa.Service, ctrl CardController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCardFactionCardContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.CardFaction(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards/factions/:factionID", ctrl.MuxHandler("CardFaction", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "CardFaction", "route", "GET /v0/cards/factions/:factionID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCardLeaderCardContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.CardLeader(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards/leaders", ctrl.MuxHandler("CardLeader", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "CardLeader", "route", "GET /v0/cards/leaders")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCardRarityCardContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.CardRarity(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards/rarities/:rarityID", ctrl.MuxHandler("CardRarity", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "CardRarity", "route", "GET /v0/cards/rarities/:rarityID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCardContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "List", "route", "GET /v0/cards")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCardContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards/:cardID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "Show", "route", "GET /v0/cards/:cardID")
}

// FactionController is the controller interface for the Faction actions.
type FactionController interface {
	goa.Muxer
	List(*ListFactionContext) error
	Show(*ShowFactionContext) error
}

// MountFactionController "mounts" a Faction resource controller on the given service.
func MountFactionController(service *goa.Service, ctrl FactionController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListFactionContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/factions", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Faction", "action", "List", "route", "GET /v0/factions")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowFactionContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/factions/:factionID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Faction", "action", "Show", "route", "GET /v0/factions/:factionID")
}

// GlyphController is the controller interface for the Glyph actions.
type GlyphController interface {
	goa.Muxer
	List(*ListGlyphContext) error
	Show(*ShowGlyphContext) error
}

// MountGlyphController "mounts" a Glyph resource controller on the given service.
func MountGlyphController(service *goa.Service, ctrl GlyphController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListGlyphContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/glyphs", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Glyph", "action", "List", "route", "GET /v0/glyphs")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowGlyphContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/glyphs/:glyphID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Glyph", "action", "Show", "route", "GET /v0/glyphs/:glyphID")
}

// PatchController is the controller interface for the Patch actions.
type PatchController interface {
	goa.Muxer
	Latest(*LatestPatchContext) error
	List(*ListPatchContext) error
	Show(*ShowPatchContext) error
}

// MountPatchController "mounts" a Patch resource controller on the given service.
func MountPatchController(service *goa.Service, ctrl PatchController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLatestPatchContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Latest(rctx)
	}
	service.Mux.Handle("GET", "/v0/patches/latest", ctrl.MuxHandler("Latest", h, nil))
	service.LogInfo("mount", "ctrl", "Patch", "action", "Latest", "route", "GET /v0/patches/latest")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListPatchContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/patches", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Patch", "action", "List", "route", "GET /v0/patches")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPatchContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/patches/:patchID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Patch", "action", "Show", "route", "GET /v0/patches/:patchID")
}

// PhonebookController is the controller interface for the Phonebook actions.
type PhonebookController interface {
	goa.Muxer
	Show(*ShowPhonebookContext) error
}

// MountPhonebookController "mounts" a Phonebook resource controller on the given service.
func MountPhonebookController(service *goa.Service, ctrl PhonebookController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPhonebookContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Phonebook", "action", "Show", "route", "GET /v0")
}

// RarityController is the controller interface for the Rarity actions.
type RarityController interface {
	goa.Muxer
	List(*ListRarityContext) error
	Show(*ShowRarityContext) error
}

// MountRarityController "mounts" a Rarity resource controller on the given service.
func MountRarityController(service *goa.Service, ctrl RarityController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListRarityContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/rarities", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Rarity", "action", "List", "route", "GET /v0/rarities")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRarityContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/rarities/:rarityID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Rarity", "action", "Show", "route", "GET /v0/rarities/:rarityID")
}

// TypeController is the controller interface for the Type actions.
type TypeController interface {
	goa.Muxer
	List(*ListTypeContext) error
	Show(*ShowTypeContext) error
}

// MountTypeController "mounts" a Type resource controller on the given service.
func MountTypeController(service *goa.Service, ctrl TypeController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListTypeContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/types", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Type", "action", "List", "route", "GET /v0/types")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowTypeContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/types/:typeID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Type", "action", "Show", "route", "GET /v0/types/:typeID")
}
