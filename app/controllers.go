//************************************************************************//
// API "gwentapi": Application Controllers
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

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
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
	service.Mux.Handle("OPTIONS", "/v0/cards/factions/:factionID", ctrl.MuxHandler("preflight", handleCardOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/cards/leaders", ctrl.MuxHandler("preflight", handleCardOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/cards/rarities/:rarityID", ctrl.MuxHandler("preflight", handleCardOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/cards", ctrl.MuxHandler("preflight", handleCardOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/cards/:cardID", ctrl.MuxHandler("preflight", handleCardOrigin(cors.HandlePreflight()), nil))

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
	h = handleCardOrigin(h)
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
	h = handleCardOrigin(h)
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
	h = handleCardOrigin(h)
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
	h = handleCardOrigin(h)
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
	h = handleCardOrigin(h)
	service.Mux.Handle("GET", "/v0/cards/:cardID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "Show", "route", "GET /v0/cards/:cardID")
}

// handleCardOrigin applies the CORS response headers corresponding to the origin.
func handleCardOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
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
	service.Mux.Handle("OPTIONS", "/v0/factions", ctrl.MuxHandler("preflight", handleFactionOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/factions/:factionID", ctrl.MuxHandler("preflight", handleFactionOrigin(cors.HandlePreflight()), nil))

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
	h = handleFactionOrigin(h)
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
	h = handleFactionOrigin(h)
	service.Mux.Handle("GET", "/v0/factions/:factionID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Faction", "action", "Show", "route", "GET /v0/factions/:factionID")
}

// handleFactionOrigin applies the CORS response headers corresponding to the origin.
func handleFactionOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
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
	service.Mux.Handle("OPTIONS", "/v0/glyphs", ctrl.MuxHandler("preflight", handleGlyphOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/glyphs/:glyphID", ctrl.MuxHandler("preflight", handleGlyphOrigin(cors.HandlePreflight()), nil))

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
	h = handleGlyphOrigin(h)
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
	h = handleGlyphOrigin(h)
	service.Mux.Handle("GET", "/v0/glyphs/:glyphID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Glyph", "action", "Show", "route", "GET /v0/glyphs/:glyphID")
}

// handleGlyphOrigin applies the CORS response headers corresponding to the origin.
func handleGlyphOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
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
	service.Mux.Handle("OPTIONS", "/v0/patches/latest", ctrl.MuxHandler("preflight", handlePatchOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/patches", ctrl.MuxHandler("preflight", handlePatchOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/patches/:patchID", ctrl.MuxHandler("preflight", handlePatchOrigin(cors.HandlePreflight()), nil))

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
	h = handlePatchOrigin(h)
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
	h = handlePatchOrigin(h)
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
	h = handlePatchOrigin(h)
	service.Mux.Handle("GET", "/v0/patches/:patchID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Patch", "action", "Show", "route", "GET /v0/patches/:patchID")
}

// handlePatchOrigin applies the CORS response headers corresponding to the origin.
func handlePatchOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
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
	service.Mux.Handle("OPTIONS", "/v0", ctrl.MuxHandler("preflight", handlePhonebookOrigin(cors.HandlePreflight()), nil))

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
	h = handlePhonebookOrigin(h)
	service.Mux.Handle("GET", "/v0", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Phonebook", "action", "Show", "route", "GET /v0")
}

// handlePhonebookOrigin applies the CORS response headers corresponding to the origin.
func handlePhonebookOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
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
	service.Mux.Handle("OPTIONS", "/v0/rarities", ctrl.MuxHandler("preflight", handleRarityOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/rarities/:rarityID", ctrl.MuxHandler("preflight", handleRarityOrigin(cors.HandlePreflight()), nil))

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
	h = handleRarityOrigin(h)
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
	h = handleRarityOrigin(h)
	service.Mux.Handle("GET", "/v0/rarities/:rarityID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Rarity", "action", "Show", "route", "GET /v0/rarities/:rarityID")
}

// handleRarityOrigin applies the CORS response headers corresponding to the origin.
func handleRarityOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// RowController is the controller interface for the Row actions.
type RowController interface {
	goa.Muxer
	List(*ListRowContext) error
	Show(*ShowRowContext) error
}

// MountRowController "mounts" a Row resource controller on the given service.
func MountRowController(service *goa.Service, ctrl RowController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v0/rows", ctrl.MuxHandler("preflight", handleRowOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/rows/:rowID", ctrl.MuxHandler("preflight", handleRowOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListRowContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleRowOrigin(h)
	service.Mux.Handle("GET", "/v0/rows", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Row", "action", "List", "route", "GET /v0/rows")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRowContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleRowOrigin(h)
	service.Mux.Handle("GET", "/v0/rows/:rowID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Row", "action", "Show", "route", "GET /v0/rows/:rowID")
}

// handleRowOrigin applies the CORS response headers corresponding to the origin.
func handleRowOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
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
	service.Mux.Handle("OPTIONS", "/v0/types", ctrl.MuxHandler("preflight", handleTypeOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v0/types/:typeID", ctrl.MuxHandler("preflight", handleTypeOrigin(cors.HandlePreflight()), nil))

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
	h = handleTypeOrigin(h)
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
	h = handleTypeOrigin(h)
	service.Mux.Handle("GET", "/v0/types/:typeID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Type", "action", "Show", "route", "GET /v0/types/:typeID")
}

// handleTypeOrigin applies the CORS response headers corresponding to the origin.
func handleTypeOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
