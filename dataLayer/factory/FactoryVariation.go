package factory

import (
	"github.com/GwentAPI/gwentapi/app"
	"github.com/GwentAPI/gwentapi/dataLayer/dal"
	"github.com/GwentAPI/gwentapi/dataLayer/models"
	"github.com/GwentAPI/gwentapi/helpers"
	"time"
)

func CreateVariation(v *models.Variation, cardID []byte, ds *dal.DataStore) (*app.GwentapiVariation, error) {
	variationUuid := helpers.EncodeUUID(v.UUID)
	cardUUID := helpers.EncodeUUID(cardID)
	dalR := dal.NewDalRarity(ds)

	rarity, err := dalR.FetchWithName(v.Rarity)
	if err != nil {
		return nil, err
	}
	rarityMedia, _ := CreateRarityLink(rarity)

	craft := &app.CostType{
		Normal:  v.Craft.Normal,
		Premium: v.Craft.Premium,
	}

	mill := &app.CostType{
		Normal:  v.Mill.Normal,
		Premium: v.Mill.Premium,
	}

	//fullSizeUrl := helpers.MediaURL(v.Art.FullsizeImage)
	thumbnailSizeUrl := helpers.MediaURL(v.Art.ThumbnailImage)
	art := &app.ArtType{
		//FullsizeImage:  &fullSizeUrl,
		ThumbnailImage: thumbnailSizeUrl,
	}
	result := &app.GwentapiVariation{
		Availability: v.Availability,
		Rarity:       rarityMedia,
		Craft:        craft,
		Mill:         mill,
		Art:          art,
		Href:         helpers.VariationURL(cardUUID, variationUuid),
		UUID:         variationUuid,
	}

	return result, nil
}

func CreateVariationLink(v *models.Variation, cardID []byte, ds *dal.DataStore) (*app.GwentapiVariationLink, error) {
	variationUuid := helpers.EncodeUUID(v.UUID)
	cardUUID := helpers.EncodeUUID(cardID)
	dalR := dal.NewDalRarity(ds)

	rarity, err := dalR.FetchWithName(v.Rarity)
	if err != nil {
		return nil, err
	}
	rarityMedia, _ := CreateRarityLink(rarity)

	result := &app.GwentapiVariationLink{
		Availability: v.Availability,
		Rarity:       rarityMedia,
		Href:         helpers.VariationURL(cardUUID, variationUuid),
	}

	return result, nil
}

func CreateVariationCollection(v *[]models.Variation, cardID []byte, ds *dal.DataStore) (app.GwentapiVariationCollection, time.Time, error) {
	variations := make(app.GwentapiVariationCollection, len(*v))
	lastModified := time.Time{}

	for i, variation := range *v {
		v, err := CreateVariation(&variation, cardID, ds)
		if err != nil {
			return nil, time.Time{}, err
		}
		if lastModified.Before(variation.Last_Modified) {
			lastModified = variation.Last_Modified
		}
		variations[i] = v
	}

	return variations, lastModified, nil
}
