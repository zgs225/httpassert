package httpassert

import (
	"bytes"
	"testing"
)

func TestEqualJSON(t *testing.T) {
	r1 := response{
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
	}

	r2 := `{"code":0,"message":"ok","data":{"total":1234,"offset":30,"limit":15,"previous":"http://localhost/previous","next":"http://localhost/next","items":[{"name":"John","age":18,"height":182.3},{"name":"Lily","age":22,"height":170.8}]}}`

	EqualJSON(t, r1, bytes.NewBufferString(r2))
}
