package assertion_test

import (
	"github.com/mleuth/assertion"
	"testing"
)

type TestEnv struct {
	Failed bool
}

func (t *TestEnv) Fatal(args ...interface{}) {
	t.Failed = true
}

func testToFail(t *testing.T, testFunc func(assertion.Assert)) {
	testEnv := TestEnv{Failed: false}
	fakeAssert := assertion.New(&testEnv)
	assert := assertion.New(t)

	testFunc(fakeAssert)
	assert.True(testEnv.Failed)
}

func TestToFail_equalIntDouble(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Equal(1.0, 1)
	})
}

func TestToFail_equalIntString(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Equal("1", 1, "Why not?")
	})
}

func TestToFail_notEqualIntDouble(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.NotEqual(1, 1)
	})
}

func TestToFail_true(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.True(false)
	})
}

func TestToFail_false(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.False(true)
	})
}

func TestToFail_nil(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Nil(1)
	})
}

func TestToFail_notNil(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.NotNil(nil)
	})
}

func TestToFail_fail(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Fail("a good reason!")
	})
}

func TestToFail_len(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Len([]int{1, 2, 3}, 2)
	})
}

func TestToFail_len_wrongType(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Len(42, 0)
	})
}

func TestToFail_contains_wrongType(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Contains(42, 0)
	})
}

func TestToFail_containsSlice_notIn(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Contains([]int{1, 2, 3}, 4)
	})
}

func TestToFail_containsMap_notIn(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Contains(map[int]int{1: 1, 2: 3, 3: 5}, 4)
	})
}

func TestToFail_containsNot_wrongType(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.ContainsNot(42, 0)
	})
}

func TestToFail_containsNotSlice_isIn(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.ContainsNot([]int{1, 2, 3}, 2)
	})
}

func TestToFail_containsNotMap_isIn(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.ContainsNot(map[int]int{1: 2, 2: 4, 3: 6}, 6)
	})
}

func TestToFail_similar_not(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Similar([]int{1, 2, 3}, []int{2, 3, 4})
	})
}

func TestToFail_similar_wrongType_map(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Similar([]int{1, 2, 3}, map[int]int{2: 1, 3: 1, 4: 1})
	})
}

func TestToFail_similar_wrongType_int(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Similar(4, []int{1, 2, 3})
	})
}

func TestToFail_similar_wrongLen(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.Similar([]int{1, 2, 3, 4}, []int{2, 3, 4})
	})
}

func TestToFail_notSimilar(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.NotSimilar([]int{1, 2, 3, 4}, []int{2, 3, 4, 1})
	})
}

func TestToFail_notSimilar_wrongType_map(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.NotSimilar([]int{1, 2, 3}, map[int]int{2: 1, 3: 1, 4: 1})
	})
}

func TestToFail_notSimilar_wrongType_int(t *testing.T) {
	testToFail(t, func(assert assertion.Assert) {
		assert.NotSimilar(4, []int{1, 2, 3})
	})
}
