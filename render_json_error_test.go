package httpassert

import "testing"

func TestRenderInterface(t *testing.T) {
	i := response{
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

	actual := renderInterface(i, 0)

	expected := `(struct response) {
    Code: (int) 0,
    Message: (string) "ok",
    Data: (*struct paginationData) {
        Total: (uint64) 1234,
        Offset: (int64) 30,
        Limit: (int64) 15,
        Previous: (string) "http://localhost/previous",
        Next: (string) "http://localhost/next",
        Items: ([]*struct user) [
            {
                Name: (string) "John",
                Age: (int) 18,
                Height: (float64) 182.3,
            },
            {
                Name: (string) "Lily",
                Age: (int) 22,
                Height: (float64) 170.8,
            },
        ],
    },
}`

	if actual != expected {
		t.Errorf("renderInterface error: \nwant: \n%s\n got: \n%s\n", expected, actual)
	}
}
