package factory

import (
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/models"
	"github.com/tri125/gwentapi/helpers"
)

func CreateVariation(v *models.Variation, cardID []byte) (*app.GwentapiVariation, error) {
	variationUuid := helpers.UUIDToURLBase64(v.UUID)
	cardUUID := helpers.UUIDToURLBase64(cardID)

	craft := &app.CostType{
		Normal:  v.Craft.Normal,
		Premium: v.Craft.Premium,
	}

	mill := &app.CostType{
		Normal:  v.Mill.Normal,
		Premium: v.Mill.Premium,
	}
	rarity := &app.GwentapiRarityLink{
		Href: "",
		Name: v.Rarity,
	}

	//fullSizeUrl := helpers.MediaURL(v.Art.FullsizeImage)
	thumbnailSizeUrl := helpers.MediaURL(v.Art.ThumbnailImage)
	art := &app.ArtType{
		//FullsizeImage:  &fullSizeUrl,
		ThumbnailImage: thumbnailSizeUrl,
	}
	result := &app.GwentapiVariation{
		Availability: v.Availability,
		Rarity:       rarity,
		Craft:        craft,
		Mill:         mill,
		Art:          art,
		Href:         helpers.VariationURL(cardUUID, variationUuid),
		UUID:         variationUuid,
	}

	return result, nil
}

func CreateLinkVariation(v *models.Variation, cardID []byte) (*app.GwentapiVariationLink, error) {
	variationUuid := helpers.UUIDToURLBase64(v.UUID)
	cardUUID := helpers.UUIDToURLBase64(cardID)

	rarity := &app.GwentapiRarityLink{
		Href: "",
		Name: v.Rarity,
	}
	result := &app.GwentapiVariationLink{
		Availability: v.Availability,
		Rarity:       rarity,
		Href:         helpers.VariationURL(variationUuid, cardUUID),
	}

	return result, nil
}

func CreateLinkVariationCollection(v *[]models.Variation, cardID []byte) (app.GwentapiVariationCollection, error) {
	variations := make(app.GwentapiVariationCollection, len(*v))
	for i, variation := range *v {
		v, _ := CreateVariation(&variation, cardID)
		variations[i] = v
	}

	return variations, nil
}
