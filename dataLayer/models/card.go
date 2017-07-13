package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Card struct {
	ID            bson.ObjectId     "_id,omitempty"
	Categories    []string          "categories"
	Faction       string            "faction"
	Flavor        map[string]string "flavor,omitempty"
	Info          map[string]string "info,omitempty"
	Strength      *int              "strength"
	Name          map[string]string "name"
	Positions     []string          "positions"
	Faction_id    bson.ObjectId     "faction_id,omitempty"
	Group         string            "group"
	Group_id      bson.ObjectId     "group_id,omitempty"
	Categories_id []bson.ObjectId   "categories_id,omitempty"
	UUID          []byte            "uuid"
	Last_Modified time.Time         "last_modified"
}
