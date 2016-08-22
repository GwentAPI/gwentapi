//************************************************************************//
// API "gwentapi": Application Contexts
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
	"strconv"
)

// CardArtworksCardContext provides the card cardArtworks action context.
type CardArtworksCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CardID string
	Limit  *int
	Offset *int
}

// NewCardArtworksCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardArtworks action.
func NewCardArtworksCardContext(ctx context.Context, service *goa.Service) (*CardArtworksCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CardArtworksCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCardID := req.Params["cardID"]
	if len(paramCardID) > 0 {
		rawCardID := paramCardID[0]
		rctx.CardID = rawCardID
	}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp2 := limit
			tmp1 := &tmp2
			rctx.Limit = tmp1
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			tmp4 := offset
			tmp3 := &tmp4
			rctx.Offset = tmp3
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardArtworksCardContext) OK(r *GwentapiArtwork) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.artwork+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *CardArtworksCardContext) OKLink(r *GwentapiArtworkLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.artwork+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardArtworksCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardArtworksCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CardFactionCardContext provides the card cardFaction action context.
type CardFactionCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FactionID string
	Limit     *int
	Offset    *int
}

// NewCardFactionCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardFaction action.
func NewCardFactionCardContext(ctx context.Context, service *goa.Service) (*CardFactionCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CardFactionCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFactionID := req.Params["factionID"]
	if len(paramFactionID) > 0 {
		rawFactionID := paramFactionID[0]
		rctx.FactionID = rawFactionID
	}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp6 := limit
			tmp5 := &tmp6
			rctx.Limit = tmp5
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			tmp8 := offset
			tmp7 := &tmp8
			rctx.Offset = tmp7
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardFactionCardContext) OK(r *GwentapiPagecard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.pagecard+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardFactionCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardFactionCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CardLeaderCardContext provides the card cardLeader action context.
type CardLeaderCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit  *int
	Offset *int
}

// NewCardLeaderCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardLeader action.
func NewCardLeaderCardContext(ctx context.Context, service *goa.Service) (*CardLeaderCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CardLeaderCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp10 := limit
			tmp9 := &tmp10
			rctx.Limit = tmp9
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			tmp12 := offset
			tmp11 := &tmp12
			rctx.Offset = tmp11
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardLeaderCardContext) OK(r *GwentapiPagecard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.pagecard+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardLeaderCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardLeaderCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CardRarityCardContext provides the card cardRarity action context.
type CardRarityCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit    *int
	Offset   *int
	RarityID string
}

// NewCardRarityCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardRarity action.
func NewCardRarityCardContext(ctx context.Context, service *goa.Service) (*CardRarityCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CardRarityCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp14 := limit
			tmp13 := &tmp14
			rctx.Limit = tmp13
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			tmp16 := offset
			tmp15 := &tmp16
			rctx.Offset = tmp15
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	paramRarityID := req.Params["rarityID"]
	if len(paramRarityID) > 0 {
		rawRarityID := paramRarityID[0]
		rctx.RarityID = rawRarityID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardRarityCardContext) OK(r *GwentapiPagecard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.pagecard+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardRarityCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardRarityCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListCardContext provides the card list action context.
type ListCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit  *int
	Offset *int
}

// NewListCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller list action.
func NewListCardContext(ctx context.Context, service *goa.Service) (*ListCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp18 := limit
			tmp17 := &tmp18
			rctx.Limit = tmp17
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			tmp20 := offset
			tmp19 := &tmp20
			rctx.Offset = tmp19
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCardContext) OK(r *GwentapiPagecard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.pagecard+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowCardContext provides the card show action context.
type ShowCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CardID string
	Limit  *int
	Offset *int
}

// NewShowCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller show action.
func NewShowCardContext(ctx context.Context, service *goa.Service) (*ShowCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCardID := req.Params["cardID"]
	if len(paramCardID) > 0 {
		rawCardID := paramCardID[0]
		rctx.CardID = rawCardID
	}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp22 := limit
			tmp21 := &tmp22
			rctx.Limit = tmp21
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) > 0 {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			tmp24 := offset
			tmp23 := &tmp24
			rctx.Offset = tmp23
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCardContext) OK(r *GwentapiCard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.card+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowCardContext) OKLink(r *GwentapiCardLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.card+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListFactionContext provides the faction list action context.
type ListFactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListFactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the faction controller list action.
func NewListFactionContext(ctx context.Context, service *goa.Service) (*ListFactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListFactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListFactionContext) OK(r GwentapiFactionCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListFactionContext) OKLink(r GwentapiFactionLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListFactionContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListFactionContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowFactionContext provides the faction show action context.
type ShowFactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FactionID string
}

// NewShowFactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the faction controller show action.
func NewShowFactionContext(ctx context.Context, service *goa.Service) (*ShowFactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowFactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFactionID := req.Params["factionID"]
	if len(paramFactionID) > 0 {
		rawFactionID := paramFactionID[0]
		rctx.FactionID = rawFactionID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowFactionContext) OK(r *GwentapiFaction) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowFactionContext) OKLink(r *GwentapiFactionLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowFactionContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowFactionContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListGlyphContext provides the glyph list action context.
type ListGlyphContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListGlyphContext parses the incoming request URL and body, performs validations and creates the
// context used by the glyph controller list action.
func NewListGlyphContext(ctx context.Context, service *goa.Service) (*ListGlyphContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListGlyphContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListGlyphContext) OK(r GwentapiGlyphCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.glyph+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListGlyphContext) OKLink(r GwentapiGlyphLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.glyph+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListGlyphContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListGlyphContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowGlyphContext provides the glyph show action context.
type ShowGlyphContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	GlyphID string
}

// NewShowGlyphContext parses the incoming request URL and body, performs validations and creates the
// context used by the glyph controller show action.
func NewShowGlyphContext(ctx context.Context, service *goa.Service) (*ShowGlyphContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowGlyphContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramGlyphID := req.Params["glyphID"]
	if len(paramGlyphID) > 0 {
		rawGlyphID := paramGlyphID[0]
		rctx.GlyphID = rawGlyphID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowGlyphContext) OK(r *GwentapiGlyph) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.glyph+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowGlyphContext) OKLink(r *GwentapiGlyphLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.glyph+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowGlyphContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowGlyphContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// LatestPatchContext provides the patch latest action context.
type LatestPatchContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewLatestPatchContext parses the incoming request URL and body, performs validations and creates the
// context used by the patch controller latest action.
func NewLatestPatchContext(ctx context.Context, service *goa.Service) (*LatestPatchContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := LatestPatchContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LatestPatchContext) OK(r *GwentapiPatch) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *LatestPatchContext) OKFull(r *GwentapiPatchFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *LatestPatchContext) OKLink(r *GwentapiPatchLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *LatestPatchContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *LatestPatchContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListPatchContext provides the patch list action context.
type ListPatchContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListPatchContext parses the incoming request URL and body, performs validations and creates the
// context used by the patch controller list action.
func NewListPatchContext(ctx context.Context, service *goa.Service) (*ListPatchContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListPatchContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListPatchContext) OK(r GwentapiPatchCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ListPatchContext) OKFull(r GwentapiPatchFullCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListPatchContext) OKLink(r GwentapiPatchLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListPatchContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListPatchContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowPatchContext provides the patch show action context.
type ShowPatchContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PatchID string
}

// NewShowPatchContext parses the incoming request URL and body, performs validations and creates the
// context used by the patch controller show action.
func NewShowPatchContext(ctx context.Context, service *goa.Service) (*ShowPatchContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowPatchContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPatchID := req.Params["patchID"]
	if len(paramPatchID) > 0 {
		rawPatchID := paramPatchID[0]
		rctx.PatchID = rawPatchID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPatchContext) OK(r *GwentapiPatch) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ShowPatchContext) OKFull(r *GwentapiPatchFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowPatchContext) OKLink(r *GwentapiPatchLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPatchContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowPatchContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowPhonebookContext provides the phonebook show action context.
type ShowPhonebookContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowPhonebookContext parses the incoming request URL and body, performs validations and creates the
// context used by the phonebook controller show action.
func NewShowPhonebookContext(ctx context.Context, service *goa.Service) (*ShowPhonebookContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowPhonebookContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPhonebookContext) OK(r *GwentapiResource) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.resource+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPhonebookContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListRarityContext provides the rarity list action context.
type ListRarityContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListRarityContext parses the incoming request URL and body, performs validations and creates the
// context used by the rarity controller list action.
func NewListRarityContext(ctx context.Context, service *goa.Service) (*ListRarityContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListRarityContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListRarityContext) OK(r GwentapiRarityCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListRarityContext) OKLink(r GwentapiRarityLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListRarityContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListRarityContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowRarityContext provides the rarity show action context.
type ShowRarityContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RarityID string
}

// NewShowRarityContext parses the incoming request URL and body, performs validations and creates the
// context used by the rarity controller show action.
func NewShowRarityContext(ctx context.Context, service *goa.Service) (*ShowRarityContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowRarityContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRarityID := req.Params["rarityID"]
	if len(paramRarityID) > 0 {
		rawRarityID := paramRarityID[0]
		rctx.RarityID = rawRarityID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowRarityContext) OK(r *GwentapiRarity) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowRarityContext) OKLink(r *GwentapiRarityLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowRarityContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowRarityContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListTypeContext provides the type list action context.
type ListTypeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListTypeContext parses the incoming request URL and body, performs validations and creates the
// context used by the type controller list action.
func NewListTypeContext(ctx context.Context, service *goa.Service) (*ListTypeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListTypeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListTypeContext) OK(r GwentapiTypeCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.type+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListTypeContext) OKLink(r GwentapiTypeLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.type+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListTypeContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListTypeContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowTypeContext provides the type show action context.
type ShowTypeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TypeID string
}

// NewShowTypeContext parses the incoming request URL and body, performs validations and creates the
// context used by the type controller show action.
func NewShowTypeContext(ctx context.Context, service *goa.Service) (*ShowTypeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowTypeContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTypeID := req.Params["typeID"]
	if len(paramTypeID) > 0 {
		rawTypeID := paramTypeID[0]
		rctx.TypeID = rawTypeID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowTypeContext) OK(r *GwentapiType) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.type+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowTypeContext) OKLink(r *GwentapiTypeLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.type+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowTypeContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowTypeContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}
