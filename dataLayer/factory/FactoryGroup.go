package factory

import (
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/models"
	"github.com/tri125/gwentapi/helpers"
)

func CreateGroup(g *models.Group) (*app.GwentapiGroup, error) {
	uuid := helpers.UUIDToURLBase64(g.UUID)

	result := &app.GwentapiGroup{
		Name: g.Name,
		Href: helpers.GroupURL(uuid),
		UUID: uuid,
	}

	return result, nil
}
