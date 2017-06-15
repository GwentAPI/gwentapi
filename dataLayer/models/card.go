package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Card struct {
	ID            bson.ObjectId   "_id,omitempty"
	Categories    []string        "categories"
	Faction       string          "faction"
	Flavor        *string         "flavor"
	Info          *string         "info"
	Strength      *int            "strength"
	Name          string          "name"
	Positions     []string        "positions"
	Faction_id    bson.ObjectId   "faction_id,omitempty"
	Group         string          "group"
	Group_id      bson.ObjectId   "group_id,omitempty"
	Categories_id []bson.ObjectId "categories_id,omitempty"
	UUID          []byte          "uuid"
	Last_Modified time.Time       "last_modified"
}
