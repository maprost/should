package should

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/maprost/should/internal/msg"
)

// BeEqual checks if 'act' == 'exp'
func BeEqual(t testing.TB, act interface{}, exp interface{}, m ...interface{}) {
	if reflect.DeepEqual(act, exp) == false {
		fail(t, msg.Default(m, "Not equal:"), msg.TypeValues(act, exp))
	}
}

// NotBeEqual checks if 'act' != 'exp'
func NotBeEqual(t testing.TB, act interface{}, exp interface{}, m ...interface{}) {
	if reflect.DeepEqual(act, exp) == true {
		fail(t, msg.Default(m, "Should not equal:"), msg.TypeValues(act, exp))
	}
}

// BeTrue checks if 'act' == true
func BeTrue(t testing.TB, act bool, m ...interface{}) {
	if !act {
		fail(t, msg.Default(m, "Should be true: "), act)
	}
}

// BeFalse checks if 'act' == false
func BeFalse(t testing.TB, act bool, m ...interface{}) {
	if act {
		fail(t, msg.Default(m, "Should be false: "), act)
	}
}

// HaveLength checks if len(col) == 'exp'
func HaveLength(t testing.TB, col interface{}, len int, m ...interface{}) {
	switch reflect.TypeOf(col).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		a := reflect.ValueOf(col).Len()
		if a != len {
			fail(t, msg.Default(m, "Wrong length:"), msg.Values(a, len))
		}

	default:
		fail(t, msg.Error("Wrong type, should be a slice, array or map."), msg.WrongType(col))
	}
}

// BeNil checks if 'act' == nil
func BeNil(t testing.TB, act interface{}, m ...interface{}) {
	if (act == nil) == false {
		fail(t, msg.Default(m, "Is not nil: "), act)
	}
}

// NotBeNil checks if 'act' != nil
func NotBeNil(t testing.TB, act interface{}, m ...interface{}) {
	if act == nil {
		fail(t, msg.Default(m, "Is nil!"))
	}
}

// Contain checks if the collection 'col' contains the given elements 'exp'.
// if 'col' is a map, it will check if the map have a value that is equal with 'exp'
func Contain(t testing.TB, col interface{}, exp interface{}, m ...interface{}) {
	switch reflect.TypeOf(col).Kind() {
	case reflect.Slice, reflect.Array:
		if isInList(col, exp) == false {
			fail(t, msg.Default(m, "Element is not in array/slice:"), msg.Collection(col, exp))
		}
	case reflect.Map:
		if isInMap(col, exp) == false {
			fail(t, msg.Default(m, "Element is not in map:"), msg.Collection(col, exp))
		}
	case reflect.String:
		colString := col.(string)
		if in := strings.Contains(colString, exp.(string)); !in {
			fail(t, msg.Default(m, "Element is not in string:"), msg.Collection(col, exp))
		}
	default:
		fail(t, msg.Error("Wrong type, should be a slice, array or map."), msg.WrongType(col))
	}
}

// NotContain checks if the collection 'col' contains not the given element 'exp'.
// if 'col' is a map, it will check if the map have not a value that is equal with 'exp'
func NotContain(t testing.TB, col interface{}, exp interface{}, m ...interface{}) {
	switch reflect.TypeOf(col).Kind() {
	case reflect.Slice, reflect.Array:
		if isInList(col, exp) {
			fail(t, msg.Default(m, "Element is in array/slice:"), msg.Collection(col, exp))
		}
	case reflect.Map:
		if isInMap(col, exp) {
			fail(t, msg.Default(m, "Element is in map:"), msg.Collection(col, exp))
		}
	default:
		fail(t, msg.Error("Wrong type, should be a slice, array or map.", msg.WrongType(col)))
	}
}

// BeSimilar checks if two arrays/slices contains the same items.
func BeSimilar(t testing.TB, act interface{}, exp interface{}, m ...interface{}) {
	aKind := reflect.TypeOf(act).Kind()
	eKind := reflect.TypeOf(exp).Kind()

	if (aKind == reflect.Array || aKind == reflect.Slice) &&
		(eKind == reflect.Array || eKind == reflect.Slice) {
		actList := reflect.ValueOf(act)
		expList := reflect.ValueOf(exp)

		// check first the length
		if actList.Len() != expList.Len() {
			fail(t, msg.Default(m, "Not similar, collections doesn't have the same length:"),
				msg.Values(actList.Len(), expList.Len()))
		}

		// check if every element of 'exp' is in 'act'
		for i := 0; i < expList.Len(); i++ {
			eValue := expList.Index(i).Interface()
			if isInList(act, eValue) == false {
				fail(t, msg.Default(m, "Not similar:"), msg.TypeValues(act, exp))
			}
		}
	} else {
		fail(t, msg.Error("Wrong type, should be a slice or array."), msg.WrongType(act))
	}
}

// NotBeSimilar checks if two arrays/slices contains at least one different item.
func NotBeSimilar(t testing.TB, act interface{}, exp interface{}, m ...interface{}) {
	aKind := reflect.TypeOf(act).Kind()
	eKind := reflect.TypeOf(exp).Kind()

	if (aKind == reflect.Array || aKind == reflect.Slice) &&
		(eKind == reflect.Array || eKind == reflect.Slice) {
		aList := reflect.ValueOf(act)
		eList := reflect.ValueOf(exp)

		// check first the length
		if aList.Len() != eList.Len() {
			// not similar
			return
		}

		// check if every element of 'exp' is in 'act'
		for i := 0; i < eList.Len(); i++ {
			eValue := eList.Index(i).Interface()
			if isInList(act, eValue) == false {
				// no similar
				return
			}
		}
		// all element are in -> similar -> fail!
		fail(t, msg.Default(m, "Similar:"), msg.TypeValues(act, exp))
	} else {
		fail(t, msg.Error("Wrong type, should be a slice or array."), msg.WrongType(act))
	}
}

// Fail with message
func Fail(t testing.TB, m ...interface{}) {
	fail(t, msg.Error(m...))
}

// ====================== Helper ===============================

func fail(t testing.TB, m ...interface{}) {
	stack := msg.StackTrace(1)
	t.Fatal(fmt.Sprint(m...), "\n", stack)
}

// check if the given array/slice contains the element
func isInList(col interface{}, elem interface{}) bool {
	list := reflect.ValueOf(col)
	for i := 0; i < list.Len(); i++ {
		e := list.Index(i).Interface()
		if reflect.DeepEqual(e, elem) {
			return true
		}
	}
	return false
}

// check if the given map contains the element as value
func isInMap(col interface{}, elem interface{}) bool {
	mp := reflect.ValueOf(col)
	for _, key := range mp.MapKeys() {
		e := mp.MapIndex(key).Interface()
		if reflect.DeepEqual(e, elem) {
			return true
		}
	}
	return false
}
