package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Variation struct {
	ID           bson.ObjectId "_id,omitempty"
	Card_id      bson.ObjectId "card_id,omitempty"
	Rarity_id    bson.ObjectId "rarity_id,omitempty"
	UUID         string
	Availability string
	Rarity       string
	Craft        Cost "craft,omitempty"
	Mill         Cost "mill,omitempty"
	Art          Art  "art,omitempty"
}

type Cost struct {
	Normal  int
	Premium int
}

type Art struct {
	FullsizeImage  *string "fullsizeImage"
	ThumbnailImage string  "thumbnailImage"
}
