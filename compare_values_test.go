package httpassert

import (
	"reflect"
	"testing"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type paginationData struct {
	Total    int64       `json:"total"`
	Offset   int64       `json:"offset"`
	Limit    int64       `json:"limit"`
	Previous string      `json:"previous"`
	Next     string      `json:"next"`
	Items    interface{} `json:"items"`
}

type user struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

func TestCompareValues(t *testing.T) {
	r1 := response{
		Code:    0,
		Message: "ok",
		Data: &paginationData{
			Total:    1234,
			Offset:   30,
			Limit:    15,
			Previous: "http://localhost/previous",
			Next:     "http://localhost/next",
			Items: []*user{
				{
					Name:   "John",
					Age:    18,
					Height: 182.3,
				},
				{
					Name:   "Lily",
					Age:    22,
					Height: 170.8,
				},
			},
		},
	}

	r2 := response{
		Code:    0,
		Message: "ok",
		Data: map[string]interface{}{
			"total":    1234,
			"offset":   30,
			"limit":    15,
			"previous": "http://localhost/previous",
			"next":     "http://localhost/next",
			"items": []map[string]interface{}{
				{
					"name":   "John",
					"age":    18,
					"height": 182.3,
				},
				{
					"name":   "Lily",
					"age":    22,
					"height": 170.8,
				},
			},
		},
	}

	if ok := compareValues(reflect.ValueOf(r1), reflect.ValueOf(r2)); !ok {
		t.Error("r1 should equal r2, but got false")
	}
}
