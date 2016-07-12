package main

import (
	"github.com/goadesign/goa"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/controllers"
	"log"
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

	// Put your logic here

	// CardController_CardFaction: end_implement
	res := app.GwentapiCardCollection{}
	return ctx.OK(res)
}

// CardLeader runs the cardLeader action.
func (c *CardController) CardLeader(ctx *app.CardLeaderCardContext) error {
	// CardController_CardLeader: start_implement

	// Put your logic here

	// CardController_CardLeader: end_implement
	res := app.GwentapiCardCollection{}
	return ctx.OK(res)
}

// CardRarity runs the cardRarity action.
func (c *CardController) CardRarity(ctx *app.CardRarityCardContext) error {
	// CardController_CardRarity: start_implement

	// Put your logic here

	// CardController_CardRarity: end_implement
	res := app.GwentapiCardCollection{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *CardController) List(ctx *app.ListCardContext) error {
	// CardController_List: start_implement

	cards, err := controllers.FetchAllCards()
	if err != nil {
		log.Println(err)
		return ctx.NotFound()
	}
	return ctx.NotFound()
	res := make(app.GwentapiCardCollection, len(cards))
	/*
		for i, card := range cards {
			c := &app.GwentapiCard{
				ID:       card.ID,
				Href:     app.CardHref(card.ID),
				Name:     card.Name,
				Faction:  card.Faction,
				Type:     card.TypeCard,
				Rarity:   card.Rarity,
				Subtypes: card.Subtypes,
				Rows:     card.Rows,
				Text:     card.Text,
				Flavor:   card.Flavor,
			}
			res[i] = c
		}
	*/
	// CardController_List: end_implement
	return ctx.OK(res)
}

// Show runs the show action.
func (c *CardController) Show(ctx *app.ShowCardContext) error {
	// CardController_Show: start_implement

	// Put your logic here

	// CardController_Show: end_implement
	res := &app.GwentapiCard{}
	return ctx.OK(res)
}
