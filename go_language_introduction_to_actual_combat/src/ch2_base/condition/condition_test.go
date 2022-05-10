package condition

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log(a)
	}

	// 方法支持多返回，if支持两段式写法：
	//if v, err := someFun(); err == nil {
	//	t.Log("")
	//} else {
	//	t.Log("")
	//}
}

// 自带 break
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("超出范围")
		}
	}
}

// 可以当做多条件if
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}
