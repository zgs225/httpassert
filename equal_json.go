package httpassert

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

// EqualJSON assert http response body equal the exptected object. The body's
// decoder is json
func EqualJSON(t *testing.T, expected interface{}, r io.Reader) {
	ev := reflect.ValueOf(expected)

	if ev.Kind() == reflect.Ptr {
		EqualJSON(t, ev.Elem().Interface(), r)
		return
	}

	av := reflect.New(ev.Type())

	b, err := ioutil.ReadAll(r)

	if err != nil {
		t.Errorf("unexpected error when decoding: %T(%v)", err, err)
		return
	}

	if err := json.Unmarshal(b, av.Interface()); err != nil {
		t.Errorf("unexpected error when decoding: %T(%v)", err, err)
		return
	}

	if ok := compareValues(reflect.ValueOf(expected), av.Elem()); ok {
		return
	}

	renderJSONError(t, expected, b)
}
