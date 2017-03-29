# assertion
lightweight test environment

## Install
```
go get github.com/mleuth/assertion
```

## Supported methods
- `Equal` -> check if two element are equal
- `NotEqual` -> check if two element are not equal
- `Nil`  -> check if an element it nil
- `NotNil`  -> check if an element is not nil
- `True`  -> check if an element is `true`
- `False` -> check if an element is `false`
- `Len` -> (only for `array`/`slice`/`map`) check the length of the collection
- `Contains` -> (only for `array`/`slice`/`map`) check if the collection contains the element
- `ContainsNot` -> (only for `array`/`slice`/`map`) check if the collection contains not the element
- `Fail` -> stops the tests with the given error message

## Usage
Please look into the test files to see the possibilities. For the first look
here some examples:

```
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
}
```

## Output
The output of a failed test shows you the actual and expected value and a stacktrace.
```
assert.go:75: Not equal:
	  actual: 1(float64)
	expected: 1(int)
	/.../src/github.com/mleuth/assertion/failing_test.go:12 +0xd1
	/usr/local/go/src/testing/testing.go:657 +0x96
	/usr/local/go/src/testing/testing.go:697 +0x2ca
```