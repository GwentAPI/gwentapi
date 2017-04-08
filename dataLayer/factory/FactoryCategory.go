package factory

import (
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/models"
	"github.com/tri125/gwentapi/helpers"
)

func CreateCategory(c *models.Category) (*app.GwentapiCategory, error) {
	uuid := helpers.EncodeUUID(c.UUID)

	result := &app.GwentapiCategory{
		Name: c.Name,
		Href: helpers.CategoryURL(uuid),
		UUID: uuid,
	}

	return result, nil
}
