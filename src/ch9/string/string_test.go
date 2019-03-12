package string

import (
	"testing"
	"unsafe"
)

/**
12 字符串
 */
func TestString(t *testing.T)  {
	var s string
	t.Log(s) //初始化为默认值""
	s = "hello"
	t.Log(len(s))
	// s[1] = '3' //string是不可变的byte slice
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据,"严"
	s = "\xE4\xBA\xB5\xFF" //不知道是什么
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T)  {
	s := "中华人民共和国"
	for _,c := range s {
		t.Logf("%[1]c %[1]d %[1]x", c)
	}
}