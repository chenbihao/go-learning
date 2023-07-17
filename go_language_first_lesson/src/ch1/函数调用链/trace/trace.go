package main

func Trace(name string) func() {
	println("enter:", name)
	return func() { // 闭包函数
		println("exit:", name)
	}
}
func foo() {
	defer Trace("foo")()
	bar()
}
func bar() {
	defer Trace("bar")()
}

func main() {
	defer Trace("main")()
	foo()
}
