package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1

	//var (
	//	a int = 1
	//	b     = 1
	//)

	// 类型推断
	a := 1
	b := 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}

}

func TestChange(t *testing.T) {
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp
	t.Log(a, b)
}
func TestChange2(t *testing.T) {
	a := 1
	b := 2
	// 交换语法糖
	a, b = b, a
	t.Log(a, b)
}
