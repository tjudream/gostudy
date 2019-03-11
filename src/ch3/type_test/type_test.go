package type_test

import "testing"

type MyInt int64
/**
 06 数据类型
 */
func TestImplicit(t *testing.T) {
	//go 不支持隐士类型转换
	//var a int = 1
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//go 不支持指针运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("s is \"\"")
	}
	//if s == nil {
	//	t.Log("s is nil")
	//}
}