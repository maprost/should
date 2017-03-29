package assertion

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"
	"testing"
)

type Assert struct {
	t *testing.T
}

func New(t *testing.T) Assert {
	return Assert{t: t}
}

func (t Assert) Equal(a interface{}, b interface{}, msg ...interface{}) {
	t.equal(a, b, defaultMsg(msg, "Not equal:"))
}

func (t Assert) NotEqual(a interface{}, b interface{}, msg ...interface{}) {
	if reflect.DeepEqual(a, b) == true {
		t.fail(1, defaultMsg(msg, "Should not equal"), actualExpectedValues(a, b))
	}
}

func (t Assert) True(a bool, msg ...interface{}) {
	t.equal(a, true, defaultMsg(msg, "Should be true:"))
}

func (t Assert) False(a bool, msg ...interface{}) {
	t.equal(a, false, defaultMsg(msg, "Should be false:"))
}

func (t Assert) Len(list interface{}, size int, msg ...interface{}) {
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		t.equal(reflect.ValueOf(list).Len(), size, defaultMsg(msg, "Wrong length:"))
	default:
		t.fail(1, "Wrong type, should be a slice, array or map.")
	}
}

func (t Assert) Nil(a interface{}, msg ...interface{}) {
	if (a == nil) == false {
		t.fail(1, defaultMsg(msg, fmt.Sprint("Is not nil: ", a)))
	}
}

func (t Assert) NotNil(a interface{}, msg ...interface{}) {
	if a == nil {
		t.fail(1, defaultMsg(msg, "Is nil"))
	}
}

// check if the collection 'c' contains the given element 'elem'.
// if the 'c' is a map, it will check if the map have a value that is equal with the 'elem'
func (t Assert) Contains(c interface{}, elem interface{}) {
	switch reflect.TypeOf(c).Kind() {
	case reflect.Slice, reflect.Array:
		if isInList(c, elem) {
			return
		}
		t.fail(1, "Element '", elem, "' is not in array/slice: \n", c)
	case reflect.Map:
		if isInMap(c, elem) {
			return
		}
		t.fail(1, "Element '", elem, "' is not in map: \n", c)
	default:
		t.fail(1, "Wrong type, should be a slice, array or map.")
	}
}

// check if the collection 'c' contains not the given element 'elem'.
// if the 'c' is a map, it will check if the map have not a value that is equal with the 'elem'
func (t Assert) ContainsNot(c interface{}, elem interface{}) {
	switch reflect.TypeOf(c).Kind() {
	case reflect.Slice, reflect.Array:
		if isInList(c, elem) {
			t.fail(1, "Element '", elem, "' is in array/slice: \n", c)
		}
	case reflect.Map:
		if isInMap(c, elem) {
			t.fail(1, "Element '", elem, "' is in map: \n", c)
		}
	default:
		t.fail(1, "Wrong type, should be a slice, array or map.")
	}
}

func (t Assert) Fail(msg ...interface{}) {
	t.fail(1, msg...)
}

func (t Assert) equal(a interface{}, b interface{}, msg string) {
	if reflect.DeepEqual(a, b) == false {
		t.fail(2, msg, actualExpectedValues(a, b))
	}
}

func (t Assert) fail(offset int, msg ...interface{}) {
	stack := getStack(offset)
	t.t.Fatal(fmt.Sprint(msg...), "\n", stack)
}

func defaultMsg(msg []interface{}, defaultMsg string) string {
	if msg == nil || len(msg) == 0 {
		return defaultMsg
	}
	return fmt.Sprint(msg...)
}

func actualExpectedValues(a interface{}, b interface{}) string {
	return fmt.Sprint(
		"\n   actual: ", a, " (", reflect.ValueOf(a).Type(), ")",
		"\n expected: ", b, " (", reflect.ValueOf(b).Type(), ")")
}

// check if the given array/slice contains the element
func isInList(c interface{}, elem interface{}) bool {
	list := reflect.ValueOf(c)
	for i := 0; i < list.Len(); i++ {
		e := list.Index(i).Interface()
		if reflect.DeepEqual(e, elem) {
			return true
		}
	}
	return false
}

// check if the given map contains the element as value
func isInMap(c interface{}, elem interface{}) bool {
	mp := reflect.ValueOf(c)
	for _, key := range mp.MapKeys() {
		e := mp.MapIndex(key).Interface()
		if reflect.DeepEqual(e, elem) {
			return true
		}
	}
	return false
}

func getStack(offset int) string {
	var stackString string = string(debug.Stack())
	var stack []string = strings.Split(stackString, "\n")

	var result string
	var goCounter int = 0
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
