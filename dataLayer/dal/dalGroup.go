package dal

import (
	"github.com/tri125/gwentapi/dataLayer/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	DalGroup struct {
		collection *mgo.Collection
	}
)

func NewDalGroup(ds *DataStore) *DalGroup {
	return &DalGroup{ds.Collection("groups")}
}

func (dc DalGroup) Fetch(uuid []byte) (*models.Group, error) {
	result := models.Group{}
	err := dc.collection.Find(bson.M{"uuid": uuid}).One(&result)
	return &result, err
}

func (dc DalGroup) FetchWithName(name string) (*models.Group, error) {
	result := models.Group{}
	err := dc.collection.Find(bson.M{"name": name}).One(&result)
	return &result, err
}

func (dc DalGroup) FetchAll() (*[]models.Group, error) {
	results := []models.Group{}
	err := dc.collection.Find(nil).Sort("name").All(&results)
	return &results, err
}

func (dc DalGroup) FetchAllPaging(limit int, offset int) (*[]models.Group, error) {
	results := []models.Group{}
	err := dc.collection.Find(nil).Limit(limit).Sort("name").Skip(offset).All(&results)
	return &results, err
}

func (dc DalGroup) Count() (int, error) {
	result, err := dc.collection.Count()
	return result, err
}
