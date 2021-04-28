package httpassert

import (
	"encoding/json"
	"math"
	"reflect"
)

func compareValues(v1, v2 reflect.Value) bool {
	if v1.Kind() == reflect.Ptr || v1.Kind() == reflect.Interface {
		return compareValues(v1.Elem(), v2)
	}

	if v2.Kind() == reflect.Ptr || v2.Kind() == reflect.Interface {
		return compareValues(v1, v2.Elem())
	}

	switch v1.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return compareInt(v1, v2, v1.Kind())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return compareUint(v1, v2, v1.Kind())
	case reflect.Float32, reflect.Float64:
		return compareFloat(v1, v2, v1.Kind())
	case reflect.Complex64, reflect.Complex128:
		return compareComplex(v1, v2, v1.Kind())
	case reflect.String:
		return compareString(v1, v2)
	case reflect.Array, reflect.Slice:
		return compareArrayLike(v1, v2, v1.Kind())
	case reflect.Bool:
		return compareBool(v1, v2)
	case reflect.Map:
		return compareMap(v1, v2)
	case reflect.Struct:
		return compareStruct(v1, v2)
	case reflect.Invalid:
		return compareInvalid(v1, v2)
	}

	return false
}

func compareInvalid(v1, v2 reflect.Value) bool {
	return v2.Kind() == reflect.Invalid
}

func compareBool(v1, v2 reflect.Value) bool {
	if v2.Kind() != reflect.Bool {
		return false
	}

	return v1.Bool() == v2.Bool()
}

func compareInt(v1, v2 reflect.Value, kind reflect.Kind) bool {
	if v2.Kind() != kind {
		return false
	}

	return v1.Int() == v2.Int()
}

func compareUint(v1, v2 reflect.Value, kind reflect.Kind) bool {
	if v2.Kind() != kind {
		return false
	}

	return v1.Uint() == v2.Uint()
}

func compareString(v1, v2 reflect.Value) bool {
	if v2.Kind() != reflect.String {
		return false
	}

	if v1.Len() != v2.Len() {
		return false
	}

	return v1.String() == v2.String()
}

func compareFloat(v1, v2 reflect.Value, kind reflect.Kind) bool {
	if v2.Kind() != kind {
		return false
	}

	f1 := v1.Float()
	f2 := v2.Float()

	return math.Float64bits(f1) == math.Float64bits(f2)
}

func compareComplex(v1, v2 reflect.Value, kind reflect.Kind) bool {
	if v2.Kind() != kind {
		return false
	}

	c1 := v1.Complex()
	c2 := v2.Complex()

	r1 := float64(real(c1))
	r2 := float64(real(c2))

	i1 := float64(imag(c1))
	i2 := float64(imag(c2))

	return math.Float64bits(r1) == math.Float64bits(r2) &&
		math.Float64bits(i1) == math.Float64bits(i2)
}

func compareArrayLike(v1, v2 reflect.Value, kind reflect.Kind) bool {
	if v2.Kind() != kind {
		return false
	}

	if v1.Len() != v2.Len() {
		return false
	}

	for i := 0; i < v1.Len(); i++ {
		if ok := compareValues(v1.Index(i), v2.Index(i)); !ok {
			return false
		}
	}

	return true
}

func compareMap(v1, v2 reflect.Value) bool {
	if v2.Kind() != reflect.Map {
		if v2.Kind() == reflect.Struct {
			return compareStructAndMap(v2, v1)
		}

		return false
	}

	if v1.Len() != v2.Len() {
		return false
	}

	iter := v1.MapRange()

	for iter.Next() {
		key := iter.Key()

		if ok := compareValues(iter.Value(), v2.MapIndex(key)); !ok {
			return false
		}
	}

	return true
}

func compareStruct(v1, v2 reflect.Value) bool {
	if v1.Type() == anyType {
		return true
	}

	if v2.Kind() != reflect.Struct {
		if v2.Kind() != reflect.Map {
			return false
		}

		return compareStructAndMap(v1, v2)
	}

	type1 := v1.Type()
	type2 := v2.Type()

	if type1 != type2 {
		return false
	}

	// If the struct have method Equal, it only receive one argument that same
	// type with v1, and only return one bool value. Then returns v1.Equal(v2)
	method := v1.MethodByName("Equal")

	if method.IsValid() {
		methodType := method.Type()

		if methodType.NumIn() == 1 && methodType.In(0) == v2.Type() && methodType.NumOut() == 1 && methodType.Out(0) == boolType {
			rets := method.Call([]reflect.Value{v2})
			return rets[0].Bool()
		}
	}

	// Otherwise compare each field of the struct
	for i := 0; i < type1.NumField(); i++ {
		field1 := type1.Field(i)

		if len(field1.PkgPath) > 0 {
			continue
		}

		name := field1.Name
		if ok := compareValues(v1.FieldByName(name), v2.FieldByName(name)); !ok {
			return false
		}
	}

	return true
}

func compareStructAndMap(v1, v2 reflect.Value) bool {
	// TODO: use encoding instead of hard code encode

	type1 := v1.Type()

	for type1.Kind() == reflect.Ptr || type1.Kind() == reflect.Interface {
		type1 = type1.Elem()
	}

	o := reflect.New(type1)

	b, err := json.Marshal(v2.Interface())

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(b, o.Interface()); err != nil {
		panic(err)
	}

	return compareStruct(v1, o.Elem())
}
