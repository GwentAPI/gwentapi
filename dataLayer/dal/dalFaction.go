package dal

import (
	"github.com/GwentAPI/gwentapi/dataLayer/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	DalFaction struct {
		collection *mgo.Collection
	}
)

func NewDalFaction(ds *DataStore) *DalFaction {
	return &DalFaction{ds.Collection("factions")}
}

func (dc DalFaction) Fetch(uuid []byte) (*models.Faction, error) {
	result := models.Faction{}
	err := dc.collection.Find(bson.M{"uuid": uuid}).One(&result)
	return &result, err
}

func (dc DalFaction) FetchWithName(name string) (*models.Faction, error) {
	result := models.Faction{}
	err := dc.collection.Find(bson.M{"name": name}).One(&result)
	return &result, err
}

func (dc DalFaction) FetchAll() (*[]models.Faction, error) {
	results := []models.Faction{}
	err := dc.collection.Find(nil).Sort("name").All(&results)
	return &results, err
}

func (dc DalFaction) FetchAllPaging(limit int, offset int) (*[]models.Faction, error) {
	results := []models.Faction{}
	err := dc.collection.Find(nil).Limit(limit).Sort("name").Skip(offset).All(&results)
	return &results, err
}

func (dc DalFaction) Count() (int, error) {
	result, err := dc.collection.Count()
	return result, err
}
