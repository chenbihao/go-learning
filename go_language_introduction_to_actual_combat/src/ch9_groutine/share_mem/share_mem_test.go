package share_mem

import (
	"sync"
	"testing"
	"time"
)

// 共享内存导致的并发问题
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 10000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Millisecond * 100)
	t.Logf("counert = %d", counter)
}

func TestCounterSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 10000; i++ {
		go func() {
			// 相当于 final 关锁
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	// 防止外面的携程比里面的先执行完
	time.Sleep(time.Millisecond * 100)
	t.Logf("counert = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			// 相当于 final 关锁
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	// 防止外面的携程比里面的先执行完
	wg.Wait()
	t.Logf("counert = %d", counter)
}
