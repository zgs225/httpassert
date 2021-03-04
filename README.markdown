HTTP ASSERT
===

A HTTP response body assert tool for golang.

## Install

`go get github.com/zgs225/httpassert`

## Example

``` go
func TestSomeHTTPRequest(t *testing.T) {
  req := httptest.NewRequest("GET", "http://localhost/users?page=2")
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
