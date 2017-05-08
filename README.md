[![Build Status](https://travis-ci.org/maprost/assertion.svg)](https://travis-ci.org/maprost/assertion)
[![Coverage Status](https://coveralls.io/repos/github/maprost/assertion/badge.svg)](https://coveralls.io/github/maprost/assertion)
[![GoDoc](https://godoc.org/github.com/maprost/assertion?status.svg)](https://godoc.org/github.com/maprost/assertion)
[![Go Report Card](https://goreportcard.com/badge/github.com/maprost/assertion)](https://goreportcard.com/report/github.com/maprost/assertion)
[![Version](https://img.shields.io/github/release/maprost/assertion.svg)](https://github.com/maprost/assertion/releases)

# assertion
lightweight test environment

## Install
```
go get github.com/maprost/assertion
```

## Supported methods
- `Equal`(`element`, `element`) -> check if two element are equal
- `NotEqual`(`element`, `element`) -> check if two element are not equal
- `Nil`(`element`)  -> check if an element it nil
- `NotNil`(`element`)  -> check if an element is not nil
- `True`(`element`)  -> check if an element is `true`
- `False`(`element`) -> check if an element is `false`
- `Len`(`collection`, `length`) -> (only for `array`/`slice`/`map`) check the length of the collection
- `Contains`(`collection`, `element`...) -> (only for `array`/`slice`/`map`) check if the collection contains the elements
- `ContainsNot`(`collection`, `element`...) -> (only for `array`/`slice`/`map`) check if the collection contains not the elements
- `Similar`(`collection`, `collection`) -> (only for `array`/`slice`) check if the two collections contains the same elements
- `NotSimilar`(`collection`, `collection`) -> (only for `array`/`slice`) check if the two collections contains at least one different element
- `Fail` -> stops the tests with the given error message

## Usage
Please look into the test files to see the possibilities. For the first look
here some examples:

```go
func TestSimple(t *testing.T) {
    assert := assertion.New(t)

    assert.Equal(1, 1)
    assert.Nil(nil)
    assert.NotNil(1)
    assert.True(true)
    assert.False(false)
    assert.NotEqual(1, 2)
}

func TestDataStructures(t *testing.T) {
    assert := assertion.New(t)

    assert.Len([]int{1, 2, 3}, 3)
    assert.Equal([]int{1, 2, 3}, []int{1, 2, 3})
    assert.NotEqual([]int{1, 2, 3}, [3]int{1, 2, 3})
    assert.Equal(map[int]string{1:"1", 2:"2", 3:"3"}, map[int]string{1:"1", 2:"2", 3:"3"})
    assert.NotEqual(map[int]string{1:"1", 3:"3"}, map[int]string{1:"1", 4:"4"})
    assert.Contains([]int{1, 2, 3}, 2)
    assert.ContainsNot([]int{1, 2, 3}, 4)
    assert.Similar([]int{1, 2, 3}, [3]int{3, 2, 1})
    assert.NotSimilar([]int{1, 2, 3}, [3]int{3, 2, 4})
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
    assert.NotEqual(p1, p3)

    assert.Equal([]Post{p1}, []Post{p2})
    assert.NotEqual([]Post{p1}, []Post{p3})
    assert.Contains([]Post{p1, p2, p3}, p3)
    assert.ContainsNot([]Post{p1, p2}, p3)
}
```

## Output
The output of a failed test shows you the actual and expected value and a stacktrace.
```
assert.go:75: Not equal:
	  actual: 1(float64)
	expected: 1(int)
	/.../src/github.com/maprost/assertion/failing_test.go:12 +0xd1
	/usr/local/go/src/testing/testing.go:657 +0x96
	/usr/local/go/src/testing/testing.go:697 +0x2ca
```