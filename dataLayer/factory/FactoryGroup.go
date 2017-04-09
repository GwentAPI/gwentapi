package factory

import (
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/models"
	"github.com/tri125/gwentapi/helpers"
)

func CreateGroup(g *models.Group) (*app.GwentapiGroup, error) {
	uuid := helpers.EncodeUUID(g.UUID)

	result := &app.GwentapiGroup{
		Name: g.Name,
		Href: helpers.GroupURL(uuid),
		UUID: uuid,
	}

	return result, nil
}

func CreateGroupLink(g *models.Group) (*app.GwentapiGroupLink, error) {
	uuid := helpers.EncodeUUID(g.UUID)
	groupLink := &app.GwentapiGroupLink{
		Href: helpers.GroupURL(uuid),
		Name: g.Name,
	}
	return groupLink, nil
}
