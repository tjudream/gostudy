package empty_interface

import (
	"fmt"
	"testing"
)
/**
18 不一样的接口类型，一样的多态
Go接口最佳实践：
	1. 倾向于使用小的接口定义，很多接口只包含一个方法
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
		type Write interface {
			Write(p []byte)(n int, err error)
		}
	2. 较大的接口定义，可以由多个小接口定义组合而成
		type ReadWriter interface {
			Reader
			Writer
		}
	3. 只依赖于必要功能的最小接口
		func StoreData(reader Reader) error {
			...
		}
 */
func DoSomething(p interface{}) {
	//if i,ok := p.(int); ok{
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if s,ok := p.(string); ok {
	//	fmt.Println("string", s)
	//	return
	//}
	//fmt.Println("Unknow Type")
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")

	}
}

func TestEmptyInterfaceAssertion(t *testing.T)  {
	DoSomething(10)
	DoSomething("10")
	DoSomething(10.5)
}
