package constant

import "testing"

const (
	// 连续常量
	Monday = iota + 1
	Tuesday
	Wednesday
)
const (
	// 状态位
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstant2(t *testing.T) {
	t.Log(Readable, Writable, Executable)
	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	a = 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
