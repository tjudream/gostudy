package _map

import (
	"testing"
)
/**
11 Map与工厂模式， 在GO语言中实现Set
 */
func TestMapWithFunValue(t *testing.T)  {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op}
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](2), m[2](2), m[3](2))
}

// Go的内置集合中没有Set，可以用map实现： map[type]bool
func TestMapForSet(t *testing.T)  {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n]{
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existint", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	t.Log(mySet)
	delete(mySet, 1)
	t.Log(mySet)
}