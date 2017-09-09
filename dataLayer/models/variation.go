package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Variation struct {
	ID            bson.ObjectId "_id,omitempty"
	Card_id       bson.ObjectId "card_id,omitempty"
	Rarity_id     bson.ObjectId "rarity_id,omitempty"
	UUID          []byte
	Availability  string
	Rarity        string
	Craft         Cost      "craft,omitempty"
	Mill          Cost      "mill,omitempty"
	Art           Art       "art,omitempty"
	Last_Modified time.Time "last_modified"
}

type Cost struct {
	Normal  int
	Premium int
}

type Art struct {
	Artist          string  "artist,omitempty"
	FullsizeImage   *string "fullsizeImage"
	MediumsizeImage string  "mediumsizeImage"
	ThumbnailImage  string  "thumbnailImage"
}
