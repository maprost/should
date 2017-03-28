package assertion_test

import (
	"github.com/mleuth/assertion"
	"testing"
)

// all tests have to fail

func TestToFail_equalIntDouble(t *testing.T) {
	assert := assertion.New(t)
	assert.Equal(1.0, 1, "equal!")
}

func TestToFail_notEqualIntDouble(t *testing.T) {
	assert := assertion.New(t)
	assert.NotEqual(1, 1, "not equal!")
}

func TestToFail_true(t *testing.T) {
	assert := assertion.New(t)
	assert.True(false, "true!")
}

func TestToFail_false(t *testing.T) {
	assert := assertion.New(t)
	assert.False(true, "false!")
}

func TestToFail_nil(t *testing.T) {
	assert := assertion.New(t)
	assert.Nil(1, "nil!")
}

func TestToFail_notNil(t *testing.T) {
	assert := assertion.New(t)
	assert.NotNil(nil, "not nil!")
}

func TestToFail_fail(t *testing.T) {
	assert := assertion.New(t)
	assert.Fail("a good reason!")
}

func TestToFail_len(t *testing.T) {
	assert := assertion.New(t)
	assert.Len([]int{1, 2, 3}, 2)
}

func TestToFail_len_wrongType(t *testing.T) {
	assert := assertion.New(t)
	assert.Len(42, 0)
}
