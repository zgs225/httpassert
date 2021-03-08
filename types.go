package httpassert

import (
	"reflect"
	"time"
)

var (
	boolType = reflect.TypeOf(true)
	timeType = reflect.TypeOf(time.Now())
)
