package assertion

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"
)

// TestEnvironment interface represent *testing.T or *testing.B.
type TestEnvironment interface {
	Fatal(args ...interface{})
}

// Assert struct holds the method and the given test environment.
type Assert struct {
	t TestEnvironment
}

// New Assert struct
func New(t TestEnvironment) Assert {
	return Assert{t: t}
}

// Equal checks if 'a' == 'e'
func (t Assert) Equal(a interface{}, e interface{}, msg ...interface{}) {
	if reflect.DeepEqual(a, e) == false {
		t.fail(1, defaultMsg(msg, "Not equal:"), actualExpectedTypeValues(a, e))
	}
}

// NotEqual checks if 'a' != 'e'
func (t Assert) NotEqual(a interface{}, e interface{}, msg ...interface{}) {
	if reflect.DeepEqual(a, e) == true {
		t.fail(1, defaultMsg(msg, "Should not equal:"), actualExpectedTypeValues(a, e))
	}
}

// True checks if 'a' == true
func (t Assert) True(a bool, msg ...interface{}) {
	if !a {
		t.fail(1, defaultMsg(msg, "Should be true: "), a)
	}
}

// False checks if 'a' == false
func (t Assert) False(a bool, msg ...interface{}) {
	if a {
		t.fail(1, defaultMsg(msg, "Should be false: "), a)
	}
}

// Len checks if len(c) == 'e'
func (t Assert) Len(c interface{}, e int, msg ...interface{}) {
	switch reflect.TypeOf(c).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		a := reflect.ValueOf(c).Len()
		if a != e {
			t.fail(1, defaultMsg(msg, "Wrong length:"), actualExpectedValues(a, e))
		}

	default:
		t.fail(1, errorMsg("Wrong type, should be a slice, array or map."), collectionElementValues(c, e))
	}
}

// Nil checks if 'a' == nil
func (t Assert) Nil(a interface{}, msg ...interface{}) {
	if (a == nil) == false {
		v := reflect.ValueOf(a)
		if v.Type().Kind() == reflect.Ptr && v.IsNil() == false {
			t.fail(1, defaultMsg(msg, "Is not nil: "), a)
		}
	}
}

// NotNil checks if 'a' != nil
func (t Assert) NotNil(a interface{}, msg ...interface{}) {
	if a == nil {
		t.fail(1, defaultMsg(msg, "Is nil!"))
	}
}

// Contains checks if the collection 'c' contains the given elements 'e'.
// if the 'c' is a map, it will check if the map have a value that is equal with the 'e'
func (t Assert) Contains(c interface{}, e ...interface{}) {
	switch reflect.TypeOf(c).Kind() {
	case reflect.Slice, reflect.Array:
		for _, eValue := range e {
			if isInList(c, eValue) == false {
				t.fail(1, errorMsg("Element is not in array/slice:"), collectionElementValues(c, e))
			}
		}
	case reflect.Map:
		for _, eValue := range e {
			if isInMap(c, eValue) == false {
				t.fail(1, errorMsg("Element is not in map:"), collectionElementValues(c, e))
			}
		}
	default:
		t.fail(1, errorMsg("Wrong type, should be a slice, array or map."), collectionElementValues(c, e))
	}
}

// ContainsNot checks if the collection 'c' contains not the given element 'e'.
// if the 'c' is a map, it will check if the map have not a value that is equal with the 'e'
func (t Assert) ContainsNot(c interface{}, e interface{}, msg ...interface{}) {
	switch reflect.TypeOf(c).Kind() {
	case reflect.Slice, reflect.Array:
		if isInList(c, e) {
			t.fail(1, defaultMsg(msg, "Element is in array/slice:"), collectionElementValues(c, e))
		}
	case reflect.Map:
		if isInMap(c, e) {
			t.fail(1, defaultMsg(msg, "Element is in map:"), collectionElementValues(c, e))
		}
	default:
		t.fail(1, errorMsg("Wrong type, should be a slice, array or map.", collectionElementValues(c, e)))
	}
}

// Similar checks if two arrays/slices contains the same items.
func (t Assert) Similar(a interface{}, e interface{}, msg ...interface{}) {
	aKind := reflect.TypeOf(a).Kind()
	eKind := reflect.TypeOf(e).Kind()

	if (aKind == reflect.Array || aKind == reflect.Slice) &&
		(eKind == reflect.Array || eKind == reflect.Slice) {
		aList := reflect.ValueOf(a)
		eList := reflect.ValueOf(e)

		// check first the length
		if aList.Len() != eList.Len() {
			t.fail(1, defaultMsg(msg, "Not similar, collections doesn't have the same length:"),
				actualExpectedValues(aList.Len(), eList.Len()))
		}

		// check if every element of 'e' is in 'a'
		for i := 0; i < eList.Len(); i++ {
			eValue := eList.Index(i).Interface()
			if isInList(a, eValue) == false {
				t.fail(1, defaultMsg(msg, "Not similar:"), actualExpectedTypeValues(a, e))
			}
		}
	} else {
		t.fail(1, errorMsg("Wrong type, should be a slice or array."), actualExpectedTypeValues(a, e))
	}
}

// NotSimilar checks if two arrays/slices contains at least one different item.
func (t Assert) NotSimilar(a interface{}, e interface{}, msg ...interface{}) {
	aKind := reflect.TypeOf(a).Kind()
	eKind := reflect.TypeOf(e).Kind()

	if (aKind == reflect.Array || aKind == reflect.Slice) &&
		(eKind == reflect.Array || eKind == reflect.Slice) {
		aList := reflect.ValueOf(a)
		eList := reflect.ValueOf(e)

		// check first the length
		if aList.Len() != eList.Len() {
			// not similar
			return
		}

		// check if every element of 'e' is in 'a'
		for i := 0; i < eList.Len(); i++ {
			eValue := eList.Index(i).Interface()
			if isInList(a, eValue) == false {
				// no similar
				return
			}
		}
		// all element are in -> similar -> fail!
		t.fail(1, defaultMsg(msg, "Similar:"), actualExpectedTypeValues(a, e))
	} else {
		t.fail(1, errorMsg("Wrong type, should be a slice or array."), actualExpectedTypeValues(a, e))
	}
}

// Fail with message
func (t Assert) Fail(msg ...interface{}) {
	t.fail(1, errorMsg(msg...))
}

// ====================== Helper ===============================

func (t Assert) fail(offset int, msg ...interface{}) {
	stack := getStack(offset)
	t.t.Fatal(fmt.Sprint(msg...), "\n", stack)
}

func defaultMsg(msg []interface{}, defaultMsg string) string {
	if msg == nil || len(msg) == 0 {
		return errorMsg(defaultMsg)
	}
	return errorMsg(msg...)
}

func errorMsg(msg ...interface{}) string {
	return fmt.Sprint(msg...)
}

func actualExpectedValues(a interface{}, b interface{}) string {
	return fmt.Sprint(
		"\n   actual: ", a,
		"\n expected: ", b)
}

func actualExpectedTypeValues(a interface{}, b interface{}) string {
	return fmt.Sprint(
		"\n   actual: ", a, " (", reflect.ValueOf(a).Type(), ")",
		"\n expected: ", b, " (", reflect.ValueOf(b).Type(), ")")
}

func collectionElementValues(a interface{}, b interface{}) string {
	return fmt.Sprint(
		"\n collection: ", a, " (", reflect.ValueOf(a).Type(), ")",
		"\n    element: ", b, " (", reflect.ValueOf(b).Type(), ")")
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
