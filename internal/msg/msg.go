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

func WrongType(a interface{}) string {
	return fmt.Sprint(
		"\ntype: ", reflect.ValueOf(a).Type())
}

func Collection(a interface{}, b interface{}) string {
	return fmt.Sprint(
		"\n collection: ", a, " (", reflect.ValueOf(a).Type(), ")",
		"\n    element: ", b, " (", reflect.ValueOf(b).Type(), ")")
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
