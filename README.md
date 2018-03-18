[![Build Status](https://travis-ci.org/maprost/should.svg)](https://travis-ci.org/maprost/should)
[![Coverage Status](https://coveralls.io/repos/github/maprost/should/badge.svg)](https://coveralls.io/github/maprost/should)
[![GoDoc](https://godoc.org/github.com/maprost/should?status.svg)](https://godoc.org/github.com/maprost/should)
[![Go Report Card](https://goreportcard.com/badge/github.com/maprost/should)](https://goreportcard.com/report/github.com/maprost/should)
[![Version](https://img.shields.io/github/release/maprost/should.svg)](https://github.com/maprost/should/releases)

# Should
lightweight test environment

## Install
```
go get github.com/maprost/should
```

## Supported methods
- `should.BeEqual`(`t`, `element`, `element`) -> check if two elements are equal
- `should.NotBeEqual`(`t`,`element`, `element`) -> check if two elements are not equal
- `should.BeNil`(`t`,`element`)  -> check if an element it nil
- `should.NotBeNil`t`,`(`element`)  -> check if an element is not nil
- `should.BeTrue`(`t`,`element`)  -> check if an element is `true`
- `should.BeFalse`(`t`,`element`) -> check if an element is `false`
- `should.HaveLength`(`t`,`collection`, `length`) -> (only for `array`/`slice`/`map`) check the length of the collection
- `should.Contain`(`t`,`collection`, `element`...) -> (only for `array`/`slice`/`map`/`string`) check if the collection contains the elements
- `should.NotContain`(`t`,`collection`, `element`...) -> (only for `array`/`slice`/`map`) check if the collection contains not the elements
- `should.BeSimilar`(`t`,`collection`, `collection`) -> (only for `array`/`slice`) check if the two collections contains the same elements
- `should.NotBeSimilar`(`t`,`collection`, `collection`) -> (only for `array`/`slice`) check if the two collections contains at least one different element
- `should.Fail` -> stops the tests with the given error message

## Usage
Please look into the test files to see the possibilities. For the first look
here some examples:

```go
func TestSimple(t *testing.T) {
	should.BeEqual(t, 1, 1, "should be equal")
	should.BeNil(t, nil)
	should.NotBeNil(t, 1)
	should.BeTrue(t, true)
	should.BeFalse(t, false)
	should.NotBeEqual(t, 1, 2)
	should.NotBeEqual(t, 1.0, 1)
	should.NotBeEqual(t, int64(1), int32(1))
}

func TestDataStructures(t *testing.T) {
	should.HaveLength(t, []int{1, 2, 3}, 3)
	should.HaveLength(t, [3]int{1, 2, 3}, 3)
	should.HaveLength(t, map[int]int{1: 1, 2: 3, 3: 5}, 3)
	should.BeEqual(t, []int{1, 2, 3}, []int{1, 2, 3})
	should.NotBeEqual(t, []int{1, 2, 3}, []int{3, 2, 1})
	should.BeSimilar(t, []int{1, 2, 3}, []int{3, 2, 1})
	should.BeSimilar(t, []int{1, 2, 3}, [3]int{3, 2, 1})
	should.NotBeSimilar(t, []int{1, 2, 3}, [3]int{3, 2, 4})
	should.NotBeSimilar(t, []int{1, 2, 3}, []int{2, 3, 4, 1})
	should.NotBeEqual(t, []int{1, 2, 3}, [3]int{1, 2, 3})
	should.BeEqual(t, [3]int{1, 2, 3}, [3]int{1, 2, 3})
	should.Contain(t, []int{1, 2, 3}, 2)
	should.NotContain(t, []int{1, 2, 3}, 4)
	should.Contain(t, [3]int{1, 2, 3}, 3)
	should.NotContain(t, [3]int{1, 2, 3}, 4)
	should.BeEqual(t, map[int]string{1: "1", 2: "2", 3: "3"}, map[int]string{1: "1", 2: "2", 3: "3"})
	should.NotBeEqual(t, map[int]string{1: "1", 3: "3"}, map[int]string{1: "1", 4: "4"})
	should.BeEqual(t, map[int]string{1: "1", 3: "3"}, map[int]string{3: "3", 1: "1"})
	should.Contain(t, map[int]string{1: "11", 3: "33"}, "33")
	should.NotContain(t, map[int]string{1: "11", 3: "33"}, "55")
}

func TestStructs(t *testing.T) {
	type Post struct {
		Id  int64
		Msg string
	}

	p1 := Post{Id: 12, Msg: "New"}
	p2 := Post{Id: 12, Msg: "New"}
	p3 := Post{Id: 12, Msg: "Old"}

	should.BeEqual(t, p1, p1)
	should.BeEqual(t, p1, p2)
	should.BeEqual(t, &p1, &p2)
	should.NotBeEqual(t, p1, p3)
	should.NotBeEqual(t, p1, &p1)

	should.BeEqual(t, []Post{p1}, []Post{p2})
	should.NotBeEqual(t, []Post{p1}, []Post{p3})
	should.Contain(t, []Post{p1, p2, p3}, p3)
	should.Contain(t, []*Post{&p1, &p2, &p3}, &p3)
	should.Contain(t, []*Post{&p1, &p3}, &p2)
	should.NotContain(t, []Post{p1, p2}, p3)
	should.NotContain(t, []Post{p1}, 22)
	should.NotContain(t, []Post{}, p1)
	should.NotContain(t, []*Post{&p1, &p2}, &p3)
	should.NotContain(t, []*Post{&p1, &p2}, p2)

	should.BeEqual(t, [1]Post{p1}, [1]Post{p2})
	should.NotBeEqual(t, [1]Post{p1}, [1]Post{p3})
	should.Contain(t, [2]Post{p1, p2}, p2)
	should.Contain(t, [3]*Post{&p1, &p2, &p3}, &p3)
	should.Contain(t, [2]*Post{&p1, &p3}, &p2)
	should.NotContain(t, [2]Post{p1, p2}, p3)
	should.NotContain(t, [1]Post{p1}, "blob")
	should.NotContain(t, [0]Post{}, p1)
	should.NotContain(t, [2]*Post{&p1, &p2}, &p3)
	should.NotContain(t, [2]*Post{&p1, &p2}, p1)

	should.BeEqual(t, map[int]Post{1: p1}, map[int]Post{1: p2})
	should.NotBeEqual(t, map[int]Post{1: p1}, map[int]Post{1: p3})
	should.NotBeEqual(t, map[int]Post{1: p1}, map[int]Post{2: p1})
	should.Contain(t, map[int]Post{1: p1, 2: p2}, p2)
	should.Contain(t, map[int]*Post{1: &p1, 2: &p2}, &p2)
	should.NotContain(t, map[int]Post{1: p1, 2: p2}, p3)
	should.NotContain(t, map[int]Post{1: p1}, "blob")
	should.NotContain(t, map[int]Post{}, p1)
	should.NotContain(t, map[int]*Post{1: &p1}, p1)
}
```

## Output
The output of a failed test shows you the actual and expected value and a stacktrace.
```
should.go:75: Not equal:
	  actual: 1(float64)
	expected: 1(int)
	/.../src/github.com/maprost/should/failing_test.go:12 +0xd1
	/usr/local/go/src/testing/testing.go:657 +0x96
	/usr/local/go/src/testing/testing.go:697 +0x2ca
```