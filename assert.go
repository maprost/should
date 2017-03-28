package assertion

import (
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
	t.equal(a, b, msg...)
}

func (t Assert) NotEqual(a interface{}, b interface{}, msg ...interface{}) {
	if reflect.DeepEqual(a, b) == true {
		t.fail(1, append(msg, "\nShould not equal: ", a, "(", reflect.ValueOf(a).Type(), ")",
			" - ", b, "(", reflect.ValueOf(b).Type(), ")")...)
	}
}

func (t Assert) True(a bool, msg ...interface{}) {
	t.equal(a, true, msg...)
}

func (t Assert) False(a bool, msg ...interface{}) {
	t.equal(a, false, msg...)
}

func (t Assert) Len(list interface{}, size int, msg ...interface{}) {
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		t.equal(reflect.ValueOf(list).Len(), size, msg...)
	default:
		t.fail(1, "No Slice, Array or Map.")
	}
}

func (t Assert) Nil(a interface{}, msg ...interface{}) {
	if (a == nil) == false {
		t.fail(1, append(msg, "\nIs not nil: ", a)...)
	}
}

func (t Assert) NotNil(a interface{}, msg ...interface{}) {
	if a == nil {
		t.fail(1, append(msg, "\nIs nil")...)
	}
}

func (t Assert) Fail(msg ...interface{}) {
	t.fail(0, msg...)
}

func (t Assert) equal(a interface{}, b interface{}, msg ...interface{}) {
	if reflect.DeepEqual(a, b) == false {
		t.fail(2, append(msg, "\nNot equal: \n", a, "(", reflect.ValueOf(a).Type(), ")",
			"\n", b, "(", reflect.ValueOf(b).Type(), ")")...)
	}
}

func (t Assert) fail(offset int, msg ...interface{}) {
	stack := getStack(offset)
	t.t.Fatal(append(msg, "\n", stack)...)
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
