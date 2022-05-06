package operator

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{2, 2, 4, 4}
	c := [...]int{1, 2, 3, 4}
	//d:=[...]int{1,2,3,4,5}

	// 只要值一样则true
	t.Log(a == b)
	t.Log(a == c)

	// 长度不同不可比较
	// Invalid operation: a==d (mismatched types [4]int and [5]int)
	// t.Log(a==d)
}

const (
	// 状态位
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant2(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	// 按位清零
	a = a &^ Readable
	t.Log(a)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
