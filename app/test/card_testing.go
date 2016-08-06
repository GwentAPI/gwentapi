package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/tri125/gwentapi/app"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

// CardFactionCardInternalServerError runs the method CardFaction of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardFactionCardInternalServerError(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, factionID string, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/factions/%v", factionID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["factionID"] = []string{fmt.Sprintf("%v", factionID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardFactionCtx, err := app.NewCardFactionCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardFaction(cardFactionCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// CardFactionCardNotFound runs the method CardFaction of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardFactionCardNotFound(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, factionID string, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/factions/%v", factionID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["factionID"] = []string{fmt.Sprintf("%v", factionID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardFactionCtx, err := app.NewCardFactionCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardFaction(cardFactionCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// CardFactionCardOK runs the method CardFaction of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardFactionCardOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, factionID string, limit *int, offset *int) (http.ResponseWriter, app.GwentapiCardCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/factions/%v", factionID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["factionID"] = []string{fmt.Sprintf("%v", factionID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardFactionCtx, err := app.NewCardFactionCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardFaction(cardFactionCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.GwentapiCardCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.GwentapiCardCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.GwentapiCardCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// CardFactionCardOKLink runs the method CardFaction of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardFactionCardOKLink(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, factionID string, limit *int, offset *int) (http.ResponseWriter, app.GwentapiCardLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/factions/%v", factionID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["factionID"] = []string{fmt.Sprintf("%v", factionID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardFactionCtx, err := app.NewCardFactionCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardFaction(cardFactionCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.GwentapiCardLinkCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.GwentapiCardLinkCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.GwentapiCardLinkCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// CardLeaderCardInternalServerError runs the method CardLeader of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardLeaderCardInternalServerError(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/leaders"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardLeaderCtx, err := app.NewCardLeaderCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardLeader(cardLeaderCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// CardLeaderCardNotFound runs the method CardLeader of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardLeaderCardNotFound(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/leaders"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardLeaderCtx, err := app.NewCardLeaderCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardLeader(cardLeaderCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// CardLeaderCardOK runs the method CardLeader of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardLeaderCardOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, limit *int, offset *int) (http.ResponseWriter, app.GwentapiCardCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/leaders"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardLeaderCtx, err := app.NewCardLeaderCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardLeader(cardLeaderCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.GwentapiCardCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.GwentapiCardCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.GwentapiCardCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// CardLeaderCardOKLink runs the method CardLeader of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardLeaderCardOKLink(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, limit *int, offset *int) (http.ResponseWriter, app.GwentapiCardLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/leaders"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardLeaderCtx, err := app.NewCardLeaderCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardLeader(cardLeaderCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.GwentapiCardLinkCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.GwentapiCardLinkCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.GwentapiCardLinkCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// CardRarityCardInternalServerError runs the method CardRarity of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardRarityCardInternalServerError(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, rarityID string, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/rarities/%v", rarityID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["rarityID"] = []string{fmt.Sprintf("%v", rarityID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardRarityCtx, err := app.NewCardRarityCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardRarity(cardRarityCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// CardRarityCardNotFound runs the method CardRarity of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardRarityCardNotFound(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, rarityID string, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/rarities/%v", rarityID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["rarityID"] = []string{fmt.Sprintf("%v", rarityID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardRarityCtx, err := app.NewCardRarityCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardRarity(cardRarityCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// CardRarityCardOK runs the method CardRarity of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardRarityCardOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, rarityID string, limit *int, offset *int) (http.ResponseWriter, app.GwentapiCardCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/rarities/%v", rarityID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["rarityID"] = []string{fmt.Sprintf("%v", rarityID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardRarityCtx, err := app.NewCardRarityCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardRarity(cardRarityCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.GwentapiCardCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.GwentapiCardCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.GwentapiCardCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// CardRarityCardOKLink runs the method CardRarity of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CardRarityCardOKLink(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, rarityID string, limit *int, offset *int) (http.ResponseWriter, app.GwentapiCardLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/rarities/%v", rarityID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["rarityID"] = []string{fmt.Sprintf("%v", rarityID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	cardRarityCtx, err := app.NewCardRarityCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.CardRarity(cardRarityCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.GwentapiCardLinkCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.GwentapiCardLinkCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.GwentapiCardLinkCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ListCardInternalServerError runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListCardInternalServerError(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	listCtx, err := app.NewListCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.List(listCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// ListCardNotFound runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListCardNotFound(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	listCtx, err := app.NewListCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.List(listCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ListCardOK runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListCardOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, limit *int, offset *int) (http.ResponseWriter, *app.GwentapiPagecard) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	listCtx, err := app.NewListCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.List(listCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.GwentapiPagecard
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.GwentapiPagecard)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.GwentapiPagecard", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ShowCardInternalServerError runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowCardInternalServerError(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, cardID string, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/%v", cardID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["cardID"] = []string{fmt.Sprintf("%v", cardID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	showCtx, err := app.NewShowCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// ShowCardNotFound runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowCardNotFound(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, cardID string, limit *int, offset *int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/%v", cardID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["cardID"] = []string{fmt.Sprintf("%v", cardID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	showCtx, err := app.NewShowCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ShowCardOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowCardOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, cardID string, limit *int, offset *int) (http.ResponseWriter, *app.GwentapiCard) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/%v", cardID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["cardID"] = []string{fmt.Sprintf("%v", cardID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	showCtx, err := app.NewShowCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.GwentapiCard
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.GwentapiCard)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.GwentapiCard", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ShowCardOKLink runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowCardOKLink(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CardController, cardID string, limit *int, offset *int) (http.ResponseWriter, *app.GwentapiCardLink) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v0/cards/%v", cardID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["cardID"] = []string{fmt.Sprintf("%v", cardID)}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if offset != nil {
		sliceVal := []string{strconv.Itoa(*offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CardTest"), rw, req, prms)
	showCtx, err := app.NewShowCardContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.GwentapiCardLink
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.GwentapiCardLink)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.GwentapiCardLink", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}
