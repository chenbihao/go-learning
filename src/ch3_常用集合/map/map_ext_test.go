package map_ext

import "testing"

// map 工厂实现
func TestMapWhitFunValue(t *testing.T) {
	m := map[int]func(op int) int{}

	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))

}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d 存在", n)
	} else {
		t.Logf("%d 不存在", n)
	}
	mySet[3] = true
	t.Logf("长度：%d", len(mySet))
	delete(mySet, 1)
	t.Logf("长度：%d", len(mySet))

	n = 1
	if mySet[n] {
		t.Logf("%d 存在", n)
	} else {
		t.Logf("%d 不存在", n)
	}
}
