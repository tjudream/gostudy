package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)
/**
15 行为定义和实现
 */
type Employee struct {
	Id string
	Name string
	Age int
}

//第一种定义方式在实例对应方法被调用时，实例的成员会进行复制
//func (e Employee) String() string {
//	fmt.Printf("in func String() e's address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

//通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String() string {
	fmt.Printf("in func String() e's address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("1. e's address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
	e1 := &Employee{"0", "Bob", 20}
	fmt.Printf("2. e's address is %x\n", unsafe.Pointer(&e1.Name))
	t.Log(e1.String())
}
func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name:"Mike", Age:30}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("&e is %T", &e)
	t.Logf("e2 is %T", e2)
}