package dal

import (
	"github.com/tri125/gwentapi/dataLayer/models"
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

func (dc DalVariation) Fetch(uuid string) (*models.Variation, error) {
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
