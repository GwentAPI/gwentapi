package helpers

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

func LastModified(data *goa.ResponseData, date time.Time) {
	setHeader(data, "Last-Modified", date)
}

func setHeader(data *goa.ResponseData, header string, value time.Time) {
	data.Header().Set(header, value.UTC().Format(http.TimeFormat))
}

func isModified(ifNotModifiedSince string, lastModified time.Time) bool {
	lastModified = lastModified.UTC()
	timeIfNotModifiedSince, err := http.ParseTime(ifNotModifiedSince)

	if err == nil && timeIfNotModifiedSince.Before(lastModified) {
		return true
	} else if err != nil {
		return true
	} else {
		return false
	}
}
