package httpassert

import (
	"bytes"
	"testing"
)

func TestEqualJSON(t *testing.T) {
	tests := []struct {
		resp interface{}
		json string
	}{
		{
			resp: response{
				Code:    0,
				Message: "ok",
				Data: &paginationData{
					Total:  1234,
					Offset: 30,
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
					Limit:    15,
					Previous: "http://localhost/previous",
					Next:     "http://localhost/next",
				},
			},
			json: `{"code":0,"message":"ok","data":{"total":1234,"offset":30,"limit":15,"previous":"http://localhost/previous","next":"http://localhost/next","items":[{"name":"John","age":18,"height":182.3},{"name":"Lily","age":22,"height":170.8}]}}`,
		},
		{
			resp: 1,
			json: "1",
		},
		{
			resp: true,
			json: "true",
		},
		{
			resp: 3.14,
			json: "3.14",
		},
	}

	for _, tt := range tests {
		EqualJSON(t, tt.resp, bytes.NewBufferString(tt.json))
	}
}
