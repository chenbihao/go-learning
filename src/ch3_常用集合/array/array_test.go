package array

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[1], arr[2])

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 5, 7}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
	for _, e := range arr3 {
		t.Log(e)
	}
}

// 截取  （不支持负数倒数截取
func TestArratSection(t *testing.T) {
	arr4 := [...]int{1, 3, 5, 7}
	arr4_sec := arr4[2:]
	t.Log(arr4_sec)
}
