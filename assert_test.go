package assertion_test

import (
	"github.com/mleuth/assertion"
	"testing"
)

func TestSimple(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(1, 1, "should be equal")
	assert.Nil(nil)
	assert.NotNil(1)
	assert.True(true)
	assert.False(false)
	assert.NotEqual(1, 2)
	assert.NotEqual(1.0, 1)
	assert.NotEqual(int64(1), int32(1))
}

func TestDataStructures(t *testing.T) {
	assert := assertion.New(t)

	assert.Len([]int{1, 2, 3}, 3)
	assert.Len([3]int{1, 2, 3}, 3)
	assert.Len(map[int]int{1: 1, 2: 3, 3: 5}, 3)
	assert.Equal([]int{1, 2, 3}, []int{1, 2, 3})
	assert.NotEqual([]int{1, 2, 3}, [3]int{1, 2, 3})
	assert.Equal([3]int{1, 2, 3}, [3]int{1, 2, 3})
	assert.Contains([]int{1, 2, 3}, 2)
	assert.ContainsNot([]int{1, 2, 3}, 4)
	assert.Contains([3]int{1, 2, 3}, 3)
	assert.ContainsNot([3]int{1, 2, 3}, 4)
	assert.Equal(map[int]string{1: "1", 2: "2", 3: "3"}, map[int]string{1: "1", 2: "2", 3: "3"})
	assert.NotEqual(map[int]string{1: "1", 3: "3"}, map[int]string{1: "1", 4: "4"})
	assert.Contains(map[int]string{1: "11", 3: "33"}, "33")
	assert.ContainsNot(map[int]string{1: "11", 3: "33"}, "55")
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

	assert.Equal([]structToTest{s1}, []structToTest{s2})
	assert.NotEqual([]structToTest{s1}, []structToTest{s3})
	assert.Contains([]structToTest{s1, s2, s3}, s3)
	assert.ContainsNot([]structToTest{s1, s2}, s3)
	assert.ContainsNot([]structToTest{s1}, 22)
	assert.ContainsNot([]structToTest{}, s1)

	assert.Equal([1]structToTest{s1}, [1]structToTest{s2})
	assert.NotEqual([1]structToTest{s1}, [1]structToTest{s3})
	assert.Contains([2]structToTest{s1, s2}, s2)
	assert.ContainsNot([2]structToTest{s1, s2}, s3)
	assert.ContainsNot([1]structToTest{s1}, "blob")
	assert.ContainsNot([0]structToTest{}, s1)

	assert.Equal(map[int]structToTest{1: s1}, map[int]structToTest{1: s2})
	assert.NotEqual(map[int]structToTest{1: s1}, map[int]structToTest{1: s3})
	assert.NotEqual(map[int]structToTest{1: s1}, map[int]structToTest{2: s1})
	assert.Contains(map[int]structToTest{1: s1, 2: s2}, s2)
	assert.ContainsNot(map[int]structToTest{1: s1, 2: s2}, s3)
	assert.ContainsNot(map[int]structToTest{1: s1}, "blob")
	assert.ContainsNot(map[int]structToTest{}, s1)
}
