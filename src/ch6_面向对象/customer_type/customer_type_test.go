package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

// func timeSpent(inner func(op int) int) func(op int) int {
// 改成自定义类型别名
func timeSpent(inner IntConv) IntConv {
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

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
