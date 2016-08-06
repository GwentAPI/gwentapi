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
	// Href to thumbnail
	Small *string `form:"small,omitempty" json:"small,omitempty" xml:"small,omitempty"`
}

// Validate validates the artworkType type instance.
func (ut *artworkType) Validate() (err error) {
	if ut.Full == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "full"))
	}
	if ut.Small == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "small"))
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
		pub.Full = *ut.Full
	}
	if ut.Small != nil {
		pub.Small = *ut.Small
	}
	return &pub
}

// Type of card artwork
type ArtworkType struct {
	// Name of the artist
	Artist *string `form:"artist,omitempty" json:"artist,omitempty" xml:"artist,omitempty"`
	// Href to full size artwork
	Full string `form:"full" json:"full" xml:"full"`
	// Href to thumbnail
	Small string `form:"small" json:"small" xml:"small"`
}

// Validate validates the ArtworkType type instance.
func (ut *ArtworkType) Validate() (err error) {
	if ut.Full == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "full"))
	}
	if ut.Small == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "small"))
	}

	return
}
