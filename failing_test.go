package should_test

import (
	"testing"

	"github.com/maprost/should"
)

type mockTest struct {
	*testing.T
	FailedTest bool
}

func (t *mockTest) Fatal(args ...interface{}) {
	t.FailedTest = true
}

func testToFail(t *testing.T, testFunc func(t testing.TB)) {
	mt := mockTest{t, false}
	testFunc(&mt)
	should.BeTrue(t, mt.FailedTest)
}

func TestToFail_equalIntDouble(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.BeEqual(t, 1.0, 1)
	})
}

func TestToFail_equalIntString(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.BeEqual(t, "1", 1, "Why not?")
	})
}

func TestToFail_notEqualIntDouble(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.NotBeEqual(t, 1, 1)
	})
}

func TestToFail_true(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.BeTrue(t, false)
	})
}

func TestToFail_false(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.BeFalse(t, true)
	})
}

func TestToFail_nil(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.BeNil(t, 1)
	})
}

func TestToFail_notNil(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.NotBeNil(t, nil)
	})
}

func TestToFail_fail(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.Fail(t, "a good reason!")
	})
}

func TestToFail_len(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.HaveLength(t, []int{1, 2, 3}, 2)
	})
}

func TestToFail_len_wrongType(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.HaveLength(t, 42, 0)
	})
}

func TestToFail_contains_wrongType(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.Contain(t, 42, 0)
	})
}

func TestToFail_containsSlice_notIn(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.Contain(t, []int{1, 2, 3}, 4)
	})
}

func TestToFail_containsMap_notIn(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.Contain(t, map[int]int{1: 1, 2: 3, 3: 5}, 4)
	})
}

func TestToFail_containsNot_wrongType(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.NotContain(t, 42, 0)
	})
}

func TestToFail_containsNotSlice_isIn(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.NotContain(t, []int{1, 2, 3}, 2)
	})
}

func TestToFail_containsNotMap_isIn(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.NotContain(t, map[int]int{1: 2, 2: 4, 3: 6}, 6)
	})
}

func TestToFail_similar_not(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.BeSimilar(t, []int{1, 2, 3}, []int{2, 3, 4})
	})
}

func TestToFail_similar_wrongType_map(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.BeSimilar(t, []int{1, 2, 3}, map[int]int{2: 1, 3: 1, 4: 1})
	})
}

func TestToFail_similar_wrongType_int(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.BeSimilar(t, 4, []int{1, 2, 3})
	})
}

func TestToFail_similar_wrongLen(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.BeSimilar(t, []int{1, 2, 3, 4}, []int{2, 3, 4})
	})
}

func TestToFail_notSimilar(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.NotBeSimilar(t, []int{1, 2, 3, 4}, []int{2, 3, 4, 1})
	})
}

func TestToFail_notSimilar_wrongType_map(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.NotBeSimilar(t, []int{1, 2, 3}, map[int]int{2: 1, 3: 1, 4: 1})
	})
}

func TestToFail_notSimilar_wrongType_int(t *testing.T) {
	testToFail(t, func(t testing.TB) {
		should.NotBeSimilar(t, 4, []int{1, 2, 3})
	})
}
