package dal

import (
	"github.com/tri125/gwentapi/dataLayer/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	DalRarity struct {
		collection *mgo.Collection
	}
)

func NewDalRarity(ds *DataStore) *DalRarity {
	return &DalRarity{ds.Collection("rarities")}
}

func (dc DalRarity) Fetch(uuid []byte) (*models.Rarity, error) {
	result := models.Rarity{}
	err := dc.collection.Find(bson.M{"uuid": uuid}).One(&result)
	return &result, err
}

func (dc DalRarity) FetchWithName(name string) (*models.Rarity, error) {
	result := models.Rarity{}
	err := dc.collection.Find(bson.M{"name": name}).One(&result)
	return &result, err
}

func (dc DalRarity) FetchAll() (*[]models.Rarity, error) {
	results := []models.Rarity{}
	err := dc.collection.Find(nil).Sort("name").All(&results)
	return &results, err
}

func (dc DalRarity) FetchAllPaging(limit int, offset int) (*[]models.Rarity, error) {
	results := []models.Rarity{}
	err := dc.collection.Find(nil).Limit(limit).Sort("name").Skip(offset).All(&results)
	return &results, err
}

func (dc DalRarity) Count() (int, error) {
	result, err := dc.collection.Count()
	return result, err
}
