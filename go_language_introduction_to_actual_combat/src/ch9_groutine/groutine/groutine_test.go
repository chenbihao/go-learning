package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}

// 携程值共享导致并发问题
func TestGroutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

// 值复制，携程正常运行
func TestGroutine3(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 100)
}
