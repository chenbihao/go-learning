package error

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

func getRandArray(n int) ([]int, error) {
	if n == 0 {
		return nil, NilError
	}
	if n < 0 || n > 10 {
		return nil, RangeError
	}
	var ret []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		ret = append(ret, rand.Intn(10))
	}
	return ret, nil
}

var NilError = errors.New("不能为0！")
var RangeError = errors.New("超出范围！")

func TestError(t *testing.T) {
	if v, err := getRandArray(0); err != nil {
		if err == NilError {
			t.Log("捕获:", err)
		}
	} else {
		t.Log(v)
	}
	if v, err := getRandArray(11); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
	if v, err := getRandArray(8); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

// 反过来写：
//	if v, err := getRandArray(8); err != nil {
//	t.Error(err)
//	return
//	}
