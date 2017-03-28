package assertion_test

import (
	"github.com/mleuth/assertion"
	"testing"
)

func TestSimple(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(1, 1)
	assert.Len([]int{1, 2, 3}, 3)
	assert.Nil(nil)
	assert.NotNil(1)
	assert.True(true)
	assert.False(false)
	assert.NotEqual(1, 2)
	assert.NotEqual(1.0, 1)
	assert.NotEqual(int64(1), int32(1))
}

func TestStructs(t *testing.T) {
	assert := assertion.New(t)

	type structToTest struct {
		Id  int64
		Msg string
	}

	s1 := structToTest{Id: 12, Msg: "New"}
	s2 := structToTest{Id: 12, Msg: "New"}
	s3 := structToTest{Id: 12, Msg: "Old"}

	assert.Equal(s1, s1)
	assert.Equal(s1, s2)
	assert.NotEqual(s1, s3)
}
