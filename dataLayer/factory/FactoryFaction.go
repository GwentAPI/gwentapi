package factory

import (
	"github.com/GwentAPI/gwentapi/app"
	"github.com/GwentAPI/gwentapi/dataLayer/models"
	"github.com/GwentAPI/gwentapi/helpers"
)

func CreateFaction(f *models.Faction) (*app.GwentapiFaction, error) {
	uuid := helpers.EncodeUUID(f.UUID)

	result := &app.GwentapiFaction{
		Name: f.Name,
		Href: helpers.FactionURL(uuid),
		UUID: uuid,
	}

	return result, nil
}

func CreateFactionLink(f *models.Faction) (*app.GwentapiFactionLink, error) {
	uuid := helpers.EncodeUUID(f.UUID)
	factionLink := &app.GwentapiFactionLink{
		Href: helpers.FactionURL(uuid),
		Name: f.Name,
	}
	return factionLink, nil
}
