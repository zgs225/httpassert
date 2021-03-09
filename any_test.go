package httpassert

import (
	"reflect"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	any := Any{}

	tests := []interface{}{
		1, "hello", time.Now(), user{}, nil, 1.3, &user{}, Any{}, []int{},
		[]interface{}{},
	}
	for _, tt := range tests {
		actual := compareValues(reflect.ValueOf(any), reflect.ValueOf(tt))

		if !actual {
			t.Errorf("any compare to (%T) %v got unexpected result", tt, tt)
		}

		actual = any.Equal(tt)

		if !actual {
			t.Errorf("any compare to (%T) %v got unexpected result", tt, tt)
		}
	}
}
