package _map

import "testing"

/**
10 Map声明、元素访问及遍历
 */
func TestInitMap(t *testing.T)  {
	m1 := map[int]int {1:1, 2:4, 3:9}
	t.Log(m1[2])
	t.Log("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Log("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Log("len m3=%d", len(m3))
	//t.Log("cap m3=%d", cap(m3)) // invalid argument m3 (type map[int]int) for cap
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v,ok := m1[3]; ok {
		t.Logf("key 3's value is %d",v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T)  {
	m1 := map[int]int {1:1, 2:4, 3:9}
	for k,v := range m1 {
		t.Log(k, v)
	}
}