HTTP ASSERT[![Go Reference](https://pkg.go.dev/badge/github.com/zgs225/httpassert.svg)](https://pkg.go.dev/github.com/zgs225/httpassert)[![Go](https://github.com/zgs225/httpassert/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/zgs225/httpassert/actions/workflows/go.yml)[![Coverage Status](https://coveralls.io/repos/github/zgs225/httpassert/badge.svg?branch=main)](https://coveralls.io/github/zgs225/httpassert?branch=main)
===

A HTTP response body assert tool for golang.

## Install

`go get github.com/zgs225/httpassert`

## Example

``` go
func TestSomeHTTPEndpoint(t *testing.T) {
  req := httptest.NewRequest("GET", "http://localhost/users?page=2", nil)
  res := httptest.NewRecord()

  someHttpServer.ServeHTTP(res, req)

  expected := Response{
    Total: 10,
    Current: 2,
    Items: []User{
      {
        Name: "John",
        Age: 20,
        Height: 183.0,
      },
      {
        Name: "Zhang San",
        Age: 18,
        Height: 179.3,
      },
    },
  }

  httpassert.EqualJSON(t, expected, res.Body)
}
```

If test fail, the print info may like:

```
  render_json_error.go:13: error: unexpected value.                                                                                                                                [205/2264]
      want response:
      (struct response) {
          Code: (int) 0,
          Message: (string) "ok",
          Data: (*struct paginationData) {
              Total: (int64) 1234,
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
      }
      got json:
      {
          "code": 1,
          "message": "ok",
          "data": {
              "total": 1234,
              "offset": 30,
              "limit": 15,
              "previous": "http://localhost/previous",
              "next": "http://localhost/next",
              "items": [
                  {
                      "name": "John",
                      "age": 18,
                      "height": 182.3
                  },
                  {
                      "name": "Lily",
                      "age": 22,
                      "height": 170.8
                  }
              ]
          }
      }
```
