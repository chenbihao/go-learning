package series

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var NilError = errors.New("不能为0！")
var RangeError = errors.New("超出范围！")

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func GetRandArray(n int) ([]int, error) {
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
