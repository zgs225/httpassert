package httpassert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func renderJSONError(t *testing.T, expected interface{}, b []byte) {
	t.Errorf(
		"error: unexpected value.\nwant response:\n%s\ngot json:\n%s\n",
		renderInterface(expected, 0),
		renderJSON(b),
	)
}

func renderInterface(i interface{}, n int) string {
	v := reflect.ValueOf(i)

	return fmt.Sprintf(
		"(%s) %s",
		renderType(reflect.TypeOf(i)),
		renderValue(v, n),
	)
}

func renderType(typ reflect.Type) string {
	var (
		name string
		kid  = typ.Kind()
	)

	switch kid {
	case reflect.Ptr:
		return fmt.Sprintf("*%s", renderType(typ.Elem()))
	case reflect.Struct:
		name := typ.Name()
		if len(name) == 0 {
			name = "anonymous"
		}
		return fmt.Sprintf("struct %s", name)
	case reflect.Array, reflect.Slice:
		return fmt.Sprintf("[]%s", renderType(typ.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", renderType(typ.Key()), renderType(typ.Elem()))
	case reflect.Chan:
		return fmt.Sprintf("chan %s", renderType(typ.Elem()))
	default:
		name = kid.String()
	}

	return name
}

func renderValue(v reflect.Value, n int) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Complex64:
		return strconv.FormatComplex(v.Complex(), 'f', -1, 64)
	case reflect.Complex128:
		return strconv.FormatComplex(v.Complex(), 'f', -1, 128)
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Ptr, reflect.Interface:
		return renderValue(v.Elem(), n)
	case reflect.Array, reflect.Slice:
		return renderArrayLikeValue(v, n)
	case reflect.Struct:
		return renderStruct(v, n)
	case reflect.Map:
		return renderMap(v, n)
	}

	return ""
}

func renderMap(v reflect.Value, n int) string {
	buf := new(bytes.Buffer)

	buf.WriteByte('{')

	keyIter := v.MapRange()

	for keyIter.Next() {
		key := keyIter.Key()
		val := keyIter.Value()

		buf.WriteByte('\n')
		buf.WriteString(nindent(renderValue(key, n), indentStr, n+1))
		buf.WriteByte(':')
		buf.WriteByte(' ')
		buf.WriteString(renderValue(val, n+1))
		buf.WriteByte(',')
	}

	buf.WriteByte('\n')
	buf.WriteString(nindent("}", indentStr, n))

	return buf.String()
}

func renderStruct(v reflect.Value, n int) string {
	buf := new(bytes.Buffer)
	typ := v.Type()

	buf.WriteByte('{')

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		if len(f.PkgPath) > 0 {
			continue
		}

		buf.WriteByte('\n')
		buf.WriteString(nindent(f.Name, indentStr, n+1))
		buf.WriteByte(':')
		buf.WriteByte(' ')
		buf.WriteString(renderInterface(v.Field(i).Interface(), n+1))
		buf.WriteByte(',')
	}

	buf.WriteByte('\n')
	buf.WriteString(nindent("}", indentStr, n))

	return buf.String()
}

func renderArrayLikeValue(v reflect.Value, n int) string {
	buf := new(bytes.Buffer)
	truncateN := 0

	buf.WriteByte('[')

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)

		for item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		if isCompositeStructKind(item.Kind()) {
			buf.WriteByte('\n')
			buf.WriteString(nindent(renderValue(item, 0), indentStr, n+1))
			buf.WriteByte(',')

			truncateN = 0

			// The last item
			if i == v.Len()-1 {
				buf.WriteByte('\n')
			}
		} else {
			buf.WriteString(renderValue(item, 0))
			buf.WriteString(", ")
			// Remove last ",\n" string
			truncateN = 2
		}
	}

	buf.Truncate(buf.Len() - truncateN)

	buf.WriteString(nindent("]", indentStr, n))

	return buf.String()
}

func renderJSON(b []byte) string {
	buf := new(bytes.Buffer)

	if err := json.Indent(buf, b, "", indentStr); err != nil {
		panic(err)
	}

	return buf.String()
}

func isCompositeStructKind(k reflect.Kind) bool {
	return k == reflect.Map ||
		k == reflect.Struct ||
		k == reflect.Array ||
		k == reflect.Slice
}
