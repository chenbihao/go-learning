package client

import (
	"ch8_package/series"

	"testing"
)

func TestPackage(t *testing.T) {

	// getRandArray 小写的方法名不可被访问到
	t.Log(series.GetRandArray(3))
	t.Log(series.GetRandArray(4))

}
