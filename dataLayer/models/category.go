package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	ID   bson.ObjectId "_id,omitempty"
	Name string
	UUID string
}
