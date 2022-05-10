package obj_cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create obj")
			return 404
		},
	}

	t.Log(pool.Get().(int))

	pool.Put(10)
	t.Log(pool.Get().(int))

	t.Log(pool.Get().(int))

}
func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create obj")
			return 404
		},
	}

	pool.Put(10)
	pool.Put(10)
	pool.Put(10)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			t.Log(pool.Get().(int))
			wg.Done()
		}()
	}
	wg.Wait()
}
