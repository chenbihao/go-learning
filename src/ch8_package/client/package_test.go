package client

import (
	// 这里加到go-path的话 直接用 ch8_package/series
	"go-demo/src/ch8_package/series"
	"testing"
)

func TestPackage(t *testing.T) {

	// getRandArray 小写的方法名不可被访问到
	t.Log(series.GetRandArray(3))
	t.Log(series.GetRandArray(4))

}
