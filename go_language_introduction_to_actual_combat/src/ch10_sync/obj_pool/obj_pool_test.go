package obj_pool

import (
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)

	// 放置一个进去 channel，会报 overflow
	//if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	//	t.Error(err)
	//}

	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			t.Logf("%T\n", v)
			// 如果没有快速失败，则会报获取超时
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	t.Log("Done")
}
