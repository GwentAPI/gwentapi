package dal

import (
	"github.com/GwentAPI/gwentapi/dataLayer/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	DalVariation struct {
		collection *mgo.Collection
	}
)

func NewDalVariation(ds *DataStore) *DalVariation {
	return &DalVariation{ds.Collection("variations")}
}

func (dc DalVariation) Fetch(uuid []byte) (*models.Variation, error) {
	result := models.Variation{}
	err := dc.collection.Find(bson.M{"uuid": uuid}).One(&result)
	return &result, err
}

func (dc DalVariation) FetchAll() (*[]models.Variation, error) {
	results := []models.Variation{}
	err := dc.collection.Find(nil).Sort("card_id").All(&results)
	return &results, err
}

func (dc DalVariation) FetchAllPaging(limit int, offset int) (*[]models.Variation, error) {
	results := []models.Variation{}
	err := dc.collection.Find(nil).Limit(limit).Sort("card_id").Skip(offset).All(&results)
	return &results, err
}

func (dc DalVariation) Count() (int, error) {
	result, err := dc.collection.Count()
	return result, err
}

func (dc DalVariation) FetchFromCardID(cardID bson.ObjectId) (*[]models.Variation, error) {
	results := []models.Variation{}
	err := dc.collection.Find(bson.M{"card_id": cardID}).All(&results)
	return &results, err
}

func (dc DalVariation) FetchCardIDFromRarityPaging(rarityID bson.ObjectId, limit int, offset int) (*[]bson.ObjectId, error) {
	var results []struct {
		Card_id bson.ObjectId "card_id,omitempty"
	}
	err := dc.collection.Find(bson.M{"rarity_id": rarityID}).Select(bson.M{"card_id": 1}).Limit(limit).Skip(offset).All(&results)

	var tmp []bson.ObjectId
	for _, v := range results {
		tmp = append(tmp, v.Card_id)
	}

	return &tmp, err
}

func (dc DalVariation) CountFromRarity(rarityID bson.ObjectId) (int, error) {
	result, err := dc.collection.Find(bson.M{"rarity_id": rarityID}).Count()
	return result, err
}
