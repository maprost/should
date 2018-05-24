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

func TestType(t *testing.T) {
	var a map[*int]*string
	should.BeEqual(t, msg.Type(a), "\ntype: map[*int]*string")
}

func TestStackTrace(t *testing.T) {
	stackTrace := msg.StackTrace(1)

	should.HaveLength(t, strings.Split(stackTrace, "\n"), 2)
	should.Contain(t, stackTrace, "/go/src/testing/testing.go")
}

func TestCollection(t *testing.T) {
	//type Drop struct {
	//	Hidden string
	//}
	//
	//drop := &Drop{"secret"}
	//should.BeEqual(t, msg.Collection([]*Drop{drop, drop}, drop), `
	//collection: [0xc420052470] ([]*msg_test.Drop)
	//   element: &{secret} (*msg_test.Drop)`)

	// first item
	// second item
	// third item
}
