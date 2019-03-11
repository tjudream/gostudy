package operator_test

import "testing"

// 07 运算符
const (
	Readable = 1 << iota
	Wirtable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1,2,3,4}
	b := [...]int{1,3,4,2}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1,2,3,4}
	t.Log(a == b)
	// 长度不同的数组无法比较
	//t.Log(a == c)
	t.Log(a == d)
}

/**
	右侧是1，则左侧清零
	右侧是0，则左侧不变
	&^ 按位置零
	1 &^ 0 -- 1
	1 &^ 1 -- 0
	0 &^ 1 -- 0
	0 &^ 0 -- 0
 */

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a & Readable, a & Wirtable, a & Executable)
	t.Log(a & Readable == Readable, a & Wirtable == Wirtable, a & Executable == Executable)
}