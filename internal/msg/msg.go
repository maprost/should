package msg

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"
)

func Default(msg []interface{}, defaultMsg string) string {
	if msg == nil || len(msg) == 0 {
		return Error(defaultMsg)
	}
	return Error(msg...)
}

func Error(msg ...interface{}) string {
	return fmt.Sprint(msg...)
}

func Values(a interface{}, b interface{}) string {
	return fmt.Sprint(
		"\n   actual: ", a,
		"\n expected: ", b)
}

func TypeValues(a interface{}, b interface{}) string {
	return fmt.Sprint(
		"\n   actual: ", a, " (", reflect.ValueOf(a).Type(), ")",
		"\n expected: ", b, " (", reflect.ValueOf(b).Type(), ")")
}

func Type(a interface{}) string {
	return fmt.Sprint(
		"\ntype: ", reflect.ValueOf(a).Type())
}

func Collection(col interface{}, elem interface{}) string {
	//return fmt.Sprint(
	//	"\n collection: ", col, " (", reflect.ValueOf(col).Type(), ")",
	//	"\n    element: ", elem, " (", reflect.ValueOf(elem).Type(), ")")

	return fmt.Sprintf(
		"\n collection:\n%s (%T)\n\n    element: %v (%T)", collectionToString(col), col, elem, elem)
}

func StackTrace(offset int) string {
	var stackString string = string(debug.Stack())
	var stack []string = strings.Split(stackString, "\n")

	var result string
	goCounter := 0
	for _, value := range stack {
		if strings.Contains(value, ".go") {
			if goCounter > 2+offset {
				result += value + "\n"
			}
			goCounter++
		}
	}

	return result
}

func collectionToString(col interface{}) string {
	switch reflect.TypeOf(col).Kind() {
	case reflect.Slice, reflect.Array:
		return arrayToString(col)
	case reflect.Map:
		return mapToString(col)
	}
	return ""
}

func arrayToString(col interface{}) (res string) {
	res = "["
	l := reflect.ValueOf(col)
	for i := 0; i < l.Len(); i++ {
		e := l.Index(i)
		res += elemToString(e)
		if i+1 != l.Len() && e.Kind() == reflect.Struct {
			res += ",\n"
		}
	}
	res += "]"
	return res
}

func mapToString(col interface{}) (res string) {
	res = "{"
	mp := reflect.ValueOf(col)
	for _, key := range mp.MapKeys() {
		e := mp.MapIndex(key)
		res += fmt.Sprint(key) + ":" + elemToString(e) + ",\n"
	}
	res += "}"
	return res
}

func elemToString(elem reflect.Value) (res string) {
	for {
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
			res = "&"
		} else {
			break
		}
	}
	res += fmt.Sprintf("%+v", elem)
	return
}
