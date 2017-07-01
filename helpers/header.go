package helpers

import (
	"github.com/goadesign/goa"
	"gopkg.in/mgo.v2"
	"net/http"
	"time"
)

func LastModified(data *goa.ResponseData, date time.Time) {
	setHeader(data, "Last-Modified", date)
}

func setHeader(data *goa.ResponseData, header string, value time.Time) {
	data.Header().Set(header, value.UTC().Format(http.TimeFormat))
}

func IsModified(notModifiedSince string, lastModified time.Time) bool {
	// Truncate to the seconds.
	lastModified = lastModified.UTC().Truncate(1 * time.Second)
	// Try to parse the string to the 3 formats allowed by HTTP/1.1: TimeFormat.
	timeIfNotModifiedSince, err := http.ParseTime(notModifiedSince)
	return err != nil || lastModified.After(timeIfNotModifiedSince)
}

func IsNotFoundError(err error) bool {
	if err == mgo.ErrNotFound {
		return true
	} else {
		return false
	}
}
