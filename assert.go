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
		t.fail(1, defaultMsg(msg,
			fmt.Sprint("Should not equal:\n  actual: ",
				valueToString(a),
				"\nexpected: ", valueToString(b))))
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

func (t Assert) Fail(msg ...interface{}) {
	t.fail(1, fmt.Sprint(msg...))
}

func (t Assert) equal(a interface{}, b interface{}, msg string) {
	if reflect.DeepEqual(a, b) == false {
		t.fail(2, fmt.Sprint(msg,
			"\n   actual: ", valueToString(a),
			"\n expected: ", valueToString(b)))
	}
}

func (t Assert) fail(offset int, msg string) {
	stack := getStack(offset)
	t.t.Fatal(msg, "\n", stack)
}

func defaultMsg(msg []interface{}, defaultMsg string) string {
	if msg == nil || len(msg) == 0 {
		return defaultMsg
	}
	return fmt.Sprint(msg...)
}

func valueToString(a interface{}) string {
	return fmt.Sprint(a, "(", reflect.ValueOf(a).Type(), ")")
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
