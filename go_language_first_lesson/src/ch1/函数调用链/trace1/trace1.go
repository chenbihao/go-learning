package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// 参考 $GOROOT/src/net/http/h2_bundle.go ，非导出函数，直接拷出来使用
var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	// bp := http2littleBuf.Get().(*[]byte)  // 替代原 http2curGoroutineID 函数中从一个 pool 池获取 byte 切片的方式
	// defer http2littleBuf.Put(bp)
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64) // 使用 strconv.ParseUint 替代了原先的 http2parseUintBytes
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

// 作者：在我的初衷里，生产环境是不应该开启该Trace的，该Trace更多是在日常dev/debug/read source code时使用。可以通过go build -tag方式开启和关闭Trace

// 并发写
var mu sync.Mutex
var m = make(map[uint64]int)

func Trace() func() {
	// 跟踪函数名（Trace1）
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	// 跟踪 GoroutineID（Trace2）
	// Tips1：可以通过 Goroutine ID 进行过滤查看（例如使用 grep 工具）
	// Tips2：Go核心团队建议：不要依赖 Goroutine ID
	gid := curGoroutineID()

	// 跟踪缩进层次（Trace3）
	mu.Lock()
	indents := m[gid]    // 获取当前gid对应的缩进层次
	m[gid] = indents + 1 // 缩进层次+1后存入map
	mu.Unlock()
	printTrace(gid, name, "->", indents+1)
	return func() {
		mu.Lock()
		indents := m[gid]    // 获取当前gid对应的缩进层次
		m[gid] = indents - 1 // 缩进层次-1后存入map
		mu.Unlock()
		printTrace(gid, name, "<-", indents)
	}
}
func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "    "
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}

func A1() {
	defer Trace()()
	B1()
}
func B1() {
	defer Trace()()
	C1()
}
func C1() {
	defer Trace()()
	D()
}
func D() { defer Trace()() }
func A2() {
	defer Trace()()
	B2()
}
func B2() {
	defer Trace()()
	C2()
}
func C2() {
	defer Trace()()
	D()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()
	A1()
	wg.Wait()
}
