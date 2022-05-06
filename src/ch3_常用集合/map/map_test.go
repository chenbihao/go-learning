package map_ext

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 6}
	t.Log(m1[2])
	t.Logf("len = %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 11
	t.Logf("len = %d", len(m2))

	m3 := make(map[int]int, 11)
	t.Logf("len = %d", len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[1] = 0
	t.Log(m1[1])

	// 默认返回0 需要进行判定

	m1[3] = 11
	if v, ok := m1[3]; ok {
		t.Log("Key 存在 ：", v)
	} else {
		t.Log("Key 不存在")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 6}
	for k, v := range m1 {
		t.Log(k, v)
	}

}
