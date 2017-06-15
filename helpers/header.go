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
