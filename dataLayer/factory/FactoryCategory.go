package factory

import (
	"github.com/GwentAPI/gwentapi/app"
	"github.com/GwentAPI/gwentapi/dataLayer/models"
	"github.com/GwentAPI/gwentapi/helpers"
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

func CreateCategoryLink(c *models.Category) (*app.GwentapiCategoryLink, error) {
	uuid := helpers.EncodeUUID(c.UUID)
	categoryLink := &app.GwentapiCategoryLink{
		Href: helpers.CategoryURL(uuid),
		Name: c.Name,
	}
	return categoryLink, nil
}
