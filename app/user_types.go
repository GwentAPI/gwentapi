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
	FullSize *string `form:"full_size,omitempty" json:"full_size,omitempty" xml:"full_size,omitempty"`
	// Href to normal size artwork
	NormalSize *string `form:"normal_size,omitempty" json:"normal_size,omitempty" xml:"normal_size,omitempty"`
}

// Validate validates the artworkType type instance.
func (ut *artworkType) Validate() (err error) {
	if ut.NormalSize == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "normal_size"))
	}

	return
}

// Publicize creates ArtworkType from artworkType
func (ut *artworkType) Publicize() *ArtworkType {
	var pub ArtworkType
	if ut.Artist != nil {
		pub.Artist = ut.Artist
	}
	if ut.FullSize != nil {
		pub.FullSize = ut.FullSize
	}
	if ut.NormalSize != nil {
		pub.NormalSize = *ut.NormalSize
	}
	return &pub
}

// Type of card artwork
type ArtworkType struct {
	// Name of the artist
	Artist *string `form:"artist,omitempty" json:"artist,omitempty" xml:"artist,omitempty"`
	// Href to full size artwork
	FullSize *string `form:"full_size,omitempty" json:"full_size,omitempty" xml:"full_size,omitempty"`
	// Href to normal size artwork
	NormalSize string `form:"normal_size" json:"normal_size" xml:"normal_size"`
}

// Validate validates the ArtworkType type instance.
func (ut *ArtworkType) Validate() (err error) {
	if ut.NormalSize == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "normal_size"))
	}

	return
}
