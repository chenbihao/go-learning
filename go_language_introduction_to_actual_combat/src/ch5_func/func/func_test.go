package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 多值返回
func returnMultiValues() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10), rand.Intn(10)
}
func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

// 传递函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("花费", time.Since(start).Seconds(), "秒")
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn2(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestSum(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("清理资源")
}

// 延迟执行
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("开始")
	panic("err")
}
