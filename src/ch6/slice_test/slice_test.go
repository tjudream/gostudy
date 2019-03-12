package slice_test

import "testing"

func TestSliceInit(t *testing.T)  {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2], s2[3], s2[4]) //panic: runtime error: index out of range
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
	//s2[4] = 5 //panic: runtime error: index out of range [recovered]
	//t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T)  {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
}

func TestSliceCompare(t *testing.T) {
	a := []int{1,2,3,4}
	b := []int{1,2,3,4}
	//invalid operation: a == b (slice can only be compared to nil)
	//if a == b {
	//	t.Log("equal")
	//}
	if a != nil {
		t.Log(a)
	}
	if b != nil {
		t.Log(b)
	}

	c := [2]int{1,2}
	d := [2]int{1,2}
	if c == d {
		t.Log("c == b")
	}
}