package helpers

import (
	"strconv"
)

func ValidateLimitOffset(count int, limit int, offset int) (int, int) {
	if offset < 0 {
		offset = 0
	}

	if offset > count {
		offset = count
	}

	if limit < 0 {
		limit = 1
	}

	if limit > count {
		limit = count
	}

	return limit, offset
}

func GeneratePrevNextPageHref(count int, limit int, offset int, href string) (*string, *string) {
	var nextHref, prevHref string
	var next, prev *string

	if count <= limit {
		return nil, nil
	}

	nextHref = href + "?limit=" + strconv.Itoa(limit) + "&offset="
	prevHref = nextHref

	prevOffset := offset - limit
	nextOffset := offset + limit

	if prevOffset < 0 {
		prevOffset = 0
	}

	if offset > 0 {
		prevHref += strconv.Itoa(prevOffset)
		prev = &prevHref
	}

	if nextOffset <= count && limit != count {
		nextHref += strconv.Itoa(nextOffset)
		next = &nextHref
	}

	return prev, next
}
