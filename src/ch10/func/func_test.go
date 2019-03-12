package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
/**
13 Go 语言的函数
14 可变参数和defer
defer延迟执行
 */
func returnMultiValues() (int,int){
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int)  func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op;
}

func TestFn(t *testing.T)  {
	a,b := returnMultiValues()
	t.Log(a, b)
	c,_ := returnMultiValues()
	t.Log(c)
	tsSF := timeSpent(slowFun)
	t.Log("return", tsSF(10))
}

func Sum(ops ...int) int{
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clean resources")
	}()
	t.Log("Started")
	panic("Fatal error")
	t.Log("final")
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer2(t *testing.T) {
	defer Clear()
	fmt.Println("Started")
}