package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
	"strconv"
)

// CardController implements the card resource.
type CardController struct {
	*goa.Controller
}

// NewCardController creates a card controller.
func NewCardController(service *goa.Service) *CardController {
	return &CardController{Controller: service.NewController("CardController")}
}

// CardFaction runs the cardFaction action.
func (c *CardController) CardFaction(ctx *app.CardFactionCardContext) error {
	// CardController_CardFaction: start_implement
	var cards []*controllers.CardModel
	var err error

	count, err := controllers.CountCards(controllers.FactionFiltered, ctx.FactionID)

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardFaction", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError (count)", err.Error())
		return ctx.InternalServerError()
	}

	limit, offset := validateLimitOffset(count, ctx.Limit, ctx.Offset)

	cards, err = controllers.FetchPageCards(controllers.FactionFiltered, limit, offset, ctx.FactionID)

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardFaction", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := generatePage(cards, count, limit, offset, controllers.CardURL("factions/"+ctx.FactionID))

	// CardController_CardFaction: end_implement
	return ctx.OK(res)
}

// CardLeader runs the cardLeader action.
func (c *CardController) CardLeader(ctx *app.CardLeaderCardContext) error {
	// CardController_CardLeader: start_implement
	var cards []*controllers.CardModel
	var err error

	count, err := controllers.CountCards(controllers.LeaderFiltered, "")

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardLeader", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError (count)", err.Error())
		return ctx.InternalServerError()
	}

	limit, offset := validateLimitOffset(count, ctx.Limit, ctx.Offset)

	cards, err = controllers.FetchPageCards(controllers.LeaderFiltered, limit, offset, "")

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardLeader", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := generatePage(cards, count, limit, offset, controllers.CardURL("leaders"))

	// CardController_CardLeader: end_implement
	return ctx.OK(res)
}

// CardRarity runs the cardRarity action.
func (c *CardController) CardRarity(ctx *app.CardRarityCardContext) error {
	// CardController_CardRarity: start_implement
	var cards []*controllers.CardModel
	var err error

	count, err := controllers.CountCards(controllers.RarityFiltered, ctx.RarityID)

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardRarity", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError (count)", err.Error())
		return ctx.InternalServerError()
	}

	limit, offset := validateLimitOffset(count, ctx.Limit, ctx.Offset)

	cards, err = controllers.FetchPageCards(controllers.RarityFiltered, limit, offset, ctx.RarityID)

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardRarity", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := generatePage(cards, count, limit, offset, controllers.CardURL("rarities/"+ctx.RarityID))
	// CardController_CardRarity: end_implement
	return ctx.OK(res)
}

// CardArtworks runs the cardArtworks action.
func (c *CardController) CardArtworks(ctx *app.CardArtworksCardContext) error {
	// CardController_CardArtworks: start_implement

	// CardController_CardArtworks: end_implement
	artwork, err := controllers.FetchArtwork(ctx.CardID)
	if controllers.NotFound(err) {
		return ctx.NotFound()
	}
	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardArtworks", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}
	// ArtworkController_Show: end_implement
	artwork.Artwork.NormalSize = controllers.MediaURL(artwork.Artwork.NormalSize)

	for index, _ := range artwork.Alternatives {
		artwork.Alternatives[index].NormalSize = controllers.MediaURL(artwork.Alternatives[index].NormalSize)
	}

	res := &app.GwentapiArtwork{
		ID:           artwork.ID,
		Href:         controllers.ArtworkURL(artwork.ID),
		Alternatives: artwork.Alternatives,
		Artwork:      artwork.Artwork,
	}
	return ctx.OK(res)
}

// List runs the list action.
func (c *CardController) List(ctx *app.ListCardContext) error {
	// CardController_List: start_implement
	var cards []*controllers.CardModel
	var err error

	count, err := controllers.CountCards(controllers.AllCards, "")

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "List", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError (count)", err.Error())
		return ctx.InternalServerError()
	}

	limit, offset := validateLimitOffset(count, ctx.Limit, ctx.Offset)

	cards, err = controllers.FetchPageCards(controllers.AllCards, limit, offset, "")

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "List", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	res := generatePage(cards, count, limit, offset, controllers.CardURL(""))

	// CardController_List: end_implement
	return ctx.OK(res)
}

// Show runs the show action.
func (c *CardController) Show(ctx *app.ShowCardContext) error {
	// CardController_Show: start_implement

	// Put your logic here
	card, err := controllers.FetchCard(ctx.CardID)
	if controllers.NotFound(err) {
		return ctx.NotFound()
	}
	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "Show", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}
	// CardController_Show: end_implement
	res := createCard(&card)

	return ctx.OK(res)
}

func createCard(card *controllers.CardModel) *app.GwentapiCard {
	c := &app.GwentapiCard{
		ID:     card.ID,
		Href:   controllers.CardURL(card.ID),
		Name:   card.Name,
		Text:   card.Text,
		Flavor: card.Flavor,
		Rows:   card.Rows,
	}

	f := &app.GwentapiFactionLink{
		Href: controllers.FactionURL(card.Faction.ID),
		Name: card.Faction.Name,
	}
	c.Faction = f

	t := &app.GwentapiTypeLink{
		Href: controllers.TypeURL(card.TypeCard.ID),
		Name: card.TypeCard.Name,
	}
	c.Type = t

	rar := &app.GwentapiRarityLink{
		Href: controllers.RarityURL(card.Rarity.ID),
		Name: card.Rarity.Name,
	}
	c.Rarity = rar

	art := &app.GwentapiArtworkLink{
		Href: controllers.ArtworkURL(card.ID),
	}

	c.Artwork = art

	typeCollection := make(app.GwentapiTypeLinkCollection, len(card.Subtypes))
	for i, cardType := range card.Subtypes {
		subt := &app.GwentapiTypeLink{
			Href: controllers.TypeURL(cardType.ID),
			Name: cardType.Name,
		}
		typeCollection[i] = subt
	}
	c.Subtypes = typeCollection

	return c
}

func createLinkCard(card *controllers.CardModel) *app.GwentapiCardLink {
	c := &app.GwentapiCardLink{
		Href: controllers.CardURL(card.ID),
		Name: card.Name,
	}
	return c
}

func generatePrevNextPageHref(count int, limit int, offset int, href string) (*string, *string) {
	var nextHref, prevHref string
	var next, prev *string

	nextHref = href + "?limit=" + strconv.Itoa(limit) + "&offset="
	prevHref = nextHref

	prevOffset := offset - limit
	nextOffset := offset + limit

	if prevOffset < 0 {
		prevOffset = 0
	}

	if offset > 0 {
		prevHref += strconv.Itoa(prevOffset)
		prev = &prevHref
	}

	if nextOffset <= count && limit != count {
		nextHref += strconv.Itoa(nextOffset)
		next = &nextHref
	}

	return prev, next
}

func validateLimitOffset(count int, ctxLimit *int, ctxOffset *int) (int, int) {
	var limit, offset int

	if ctxOffset == nil {
		offset = 0
	} else {
		offset = *ctxOffset
	}
	if ctxLimit == nil {
		limit = 20
	} else {
		limit = *ctxLimit
	}

	if offset < 0 {
		offset = 0
	}

	if offset > count {
		offset = count
	}

	if limit < 0 {
		limit = 1
	}

	if limit > count {
		limit = count
	}

	return limit, offset
}

func generatePage(cards []*controllers.CardModel, count int, limit int, offset int, href string) *app.GwentapiPagecard {

	cardResult := make(app.GwentapiCardLinkCollection, len(cards))

	for i, card := range cards {
		cardResult[i] = createLinkCard(card)
	}

	prev, next := generatePrevNextPageHref(count, limit, offset, href)

	res := &app.GwentapiPagecard{
		Count:    count,
		Next:     next,
		Previous: prev,
		Results:  cardResult,
	}

	return res
}
