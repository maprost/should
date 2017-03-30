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
	assert.NotEqual([]int{1, 2, 3}, []int{3, 2, 1})
	assert.Similar([]int{1, 2, 3}, []int{3, 2, 1})
	assert.Similar([]int{1, 2, 3}, [3]int{3, 2, 1})
	assert.NotSimilar([]int{1, 2, 3}, [3]int{3, 2, 4})
	assert.NotSimilar([]int{1, 2, 3}, []int{2, 3, 4, 1})
	assert.NotEqual([]int{1, 2, 3}, [3]int{1, 2, 3})
	assert.Equal([3]int{1, 2, 3}, [3]int{1, 2, 3})
	assert.Contains([]int{1, 2, 3}, 2)
	assert.ContainsNot([]int{1, 2, 3}, 4)
	assert.Contains([3]int{1, 2, 3}, 3)
	assert.ContainsNot([3]int{1, 2, 3}, 4)
	assert.Equal(map[int]string{1: "1", 2: "2", 3: "3"}, map[int]string{1: "1", 2: "2", 3: "3"})
	assert.NotEqual(map[int]string{1: "1", 3: "3"}, map[int]string{1: "1", 4: "4"})
	assert.Equal(map[int]string{1: "1", 3: "3"}, map[int]string{3: "3", 1: "1"})
	assert.Contains(map[int]string{1: "11", 3: "33"}, "33")
	assert.ContainsNot(map[int]string{1: "11", 3: "33"}, "55")
}

func TestStructs(t *testing.T) {
	assert := assertion.New(t)

	type Post struct {
		Id  int64
		Msg string
	}

	p1 := Post{Id: 12, Msg: "New"}
	p2 := Post{Id: 12, Msg: "New"}
	p3 := Post{Id: 12, Msg: "Old"}

	assert.Equal(p1, p1)
	assert.Equal(p1, p2)
	assert.Equal(&p1, &p2)
	assert.NotEqual(p1, p3)
	assert.NotEqual(p1, &p1)

	assert.Equal([]Post{p1}, []Post{p2})
	assert.NotEqual([]Post{p1}, []Post{p3})
	assert.Contains([]Post{p1, p2, p3}, p3)
	assert.Contains([]*Post{&p1, &p2, &p3}, &p3)
	assert.Contains([]*Post{&p1, &p3}, &p2)
	assert.ContainsNot([]Post{p1, p2}, p3)
	assert.ContainsNot([]Post{p1}, 22)
	assert.ContainsNot([]Post{}, p1)
	assert.ContainsNot([]*Post{&p1, &p2}, &p3)
	assert.ContainsNot([]*Post{&p1, &p2}, p2)

	assert.Equal([1]Post{p1}, [1]Post{p2})
	assert.NotEqual([1]Post{p1}, [1]Post{p3})
	assert.Contains([2]Post{p1, p2}, p2)
	assert.Contains([3]*Post{&p1, &p2, &p3}, &p3)
	assert.Contains([2]*Post{&p1, &p3}, &p2)
	assert.ContainsNot([2]Post{p1, p2}, p3)
	assert.ContainsNot([1]Post{p1}, "blob")
	assert.ContainsNot([0]Post{}, p1)
	assert.ContainsNot([2]*Post{&p1, &p2}, &p3)
	assert.ContainsNot([2]*Post{&p1, &p2}, p1)

	assert.Equal(map[int]Post{1: p1}, map[int]Post{1: p2})
	assert.NotEqual(map[int]Post{1: p1}, map[int]Post{1: p3})
	assert.NotEqual(map[int]Post{1: p1}, map[int]Post{2: p1})
	assert.Contains(map[int]Post{1: p1, 2: p2}, p2)
	assert.Contains(map[int]*Post{1: &p1, 2: &p2}, &p2)
	assert.ContainsNot(map[int]Post{1: p1, 2: p2}, p3)
	assert.ContainsNot(map[int]Post{1: p1}, "blob")
	assert.ContainsNot(map[int]Post{}, p1)
	assert.ContainsNot(map[int]*Post{1: &p1}, p1)
}
