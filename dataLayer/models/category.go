package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Category struct {
	ID            bson.ObjectId "_id,omitempty"
	Name          string
	UUID          []byte
	Last_Modified time.Time "last_modified"
}
