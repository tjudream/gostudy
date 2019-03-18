package _interface

import "testing"
/**
16 Go语言的相关接口
Go的接口与其他语言的区别：
	1. 接口为非侵入性的，实现不依赖于接口定义
	2. 所以接口的定义可以包含在接口使用者包内
 */
type Programmer interface{
	WriteHelloWorld() string
}

type GoProgramer struct{

}

func (g *GoProgramer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgramer)
	t.Log(p.WriteHelloWorld())
}