package msg_test

import (
	"strings"
	"testing"

	"github.com/maprost/should"
	"github.com/maprost/should/internal/msg"
)

func TestDefault(t *testing.T) {
	should.BeEqual(t, msg.Default([]interface{}{"Blob"}, "Drop"), "Blob")
	should.BeEqual(t, msg.Default([]interface{}{}, "Drop"), "Drop")
	should.BeEqual(t, msg.Default(nil, "Drop"), "Drop")
}

func TestError(t *testing.T) {
	should.BeEqual(t, msg.Error("jo", "no", "so"), "jonoso")
	should.BeEqual(t, msg.Error(), "")
}

func TestWrongType(t *testing.T) {
	var a map[*int]*string
	should.BeEqual(t, msg.WrongType(a), "\ntype: map[*int]*string")
}

func TestStackTrace(t *testing.T) {
	stackTrace := msg.StackTrace(1)

	should.HaveLength(t, strings.Split(stackTrace, "\n"), 2)
	should.Contain(t, stackTrace, "/go/src/testing/testing.go")
}
