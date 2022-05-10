package try_type

import "testing"

type MyInt int64

// 隐式类型转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)

	t.Log(a, b, c)
}

// 指针
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// 不支持指针计算
	//aPtr = aPtr+1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("【" + s + "】")
	t.Log(len(s))
	if s == "" {
		t.Log("string 为空")
	}

}
