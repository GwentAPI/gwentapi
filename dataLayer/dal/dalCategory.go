package dal

import (
	"github.com/tri125/gwentapi/dataLayer/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PageQueryType int

const (
	AllCards PageQueryType = 1 << iota
	RarityFiltered
	LeaderFiltered
	FactionFiltered
)

type (
	DalCategory struct {
		collection *mgo.Collection
	}
)

func NewDalCategory(ds *DataStore) *DalCategory {
	return &DalCategory{ds.Collection("categories")}
}

func (dc DalCategory) Fetch(uuid string) (*models.Category, error) {
	result := models.Category{}
	err := dc.collection.Find(bson.M{"uuid": uuid}).One(&result)
	return &result, err
}

func (dc DalCategory) FetchAll() (*[]models.Category, error) {
	results := []models.Category{}
	err := dc.collection.Find(nil).Sort("name").All(&results)
	return &results, err
}

func (dc DalCategory) FetchAllPaging(limit int, offset int) (*[]models.Category, error) {
	results := []models.Category{}
	err := dc.collection.Find(nil).Limit(limit).Sort("name").Skip(offset).All(&results)
	return &results, err
}

func (dc DalCategory) Count() (int, error) {
	result, err := dc.collection.Count()
	return result, err
}
