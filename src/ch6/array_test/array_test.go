package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int{1,3,4,5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T)  {
	arr3 := [...]int{1,3,4,5}
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
/*
数组截取：
a[开始索引（包含），结束索引（不包含）]
a := [...]int{1,2,3,4,5}
a[1:2] //2
a[1:3] //2,3
a[1:len(a)] //2,3,4,5
a[1:] //2,3,4,5
a[:3] //1,2,3
*/

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1,3,4,5}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)
	arr3_sec2 := arr3[:]
	t.Log(arr3_sec2)
}