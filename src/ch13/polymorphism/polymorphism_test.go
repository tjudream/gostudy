package polymorphism

import (
	"fmt"
	"testing"
)
/**
18 不一样的接口类型，一样的多态
 */
type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct{

}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	goProg1 := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(goProg1)
	writeFirstProgram(javaProg)
}

/**
空接口与断言
1. 空接口可以表示任何类型
2. 通过断言来将空接口转换为制定类型
v,ok := p.(int) // ok=true 时为转换成功
 */