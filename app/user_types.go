//************************************************************************//
// API "gwentapi": Application User Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/tri125/gwentapi/design
// --out=$(GOPATH)\src\github.com\tri125\gwentapi
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// Type of card artwork
type artworkType struct {
	// Name of the artist
	Artist *string `form:"artist,omitempty" json:"artist,omitempty" xml:"artist,omitempty"`
	// Href to full size artwork
	Full *string `form:"full,omitempty" json:"full,omitempty" xml:"full,omitempty"`
	// Href to normal size artwork
	Normal *string `form:"normal,omitempty" json:"normal,omitempty" xml:"normal,omitempty"`
}

// Validate validates the artworkType type instance.
func (ut *artworkType) Validate() (err error) {
	if ut.Normal == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "normal"))
	}

	return
}

// Publicize creates ArtworkType from artworkType
func (ut *artworkType) Publicize() *ArtworkType {
	var pub ArtworkType
	if ut.Artist != nil {
		pub.Artist = ut.Artist
	}
	if ut.Full != nil {
		pub.Full = ut.Full
	}
	if ut.Normal != nil {
		pub.Normal = *ut.Normal
	}
	return &pub
}

// Type of card artwork
type ArtworkType struct {
	// Name of the artist
	Artist *string `form:"artist,omitempty" json:"artist,omitempty" xml:"artist,omitempty"`
	// Href to full size artwork
	Full *string `form:"full,omitempty" json:"full,omitempty" xml:"full,omitempty"`
	// Href to normal size artwork
	Normal string `form:"normal" json:"normal" xml:"normal"`
}

// Validate validates the ArtworkType type instance.
func (ut *ArtworkType) Validate() (err error) {
	if ut.Normal == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "normal"))
	}

	return
}
