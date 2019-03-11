package constant_test

import (
	"testing"
)

const(
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Wirtable
	Executable
)

func TestConstantTry(t *testing.T)  {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1
	t.Log(a & Readable, a & Wirtable, a & Executable)
	t.Log(a & Readable == Readable, a & Wirtable == Wirtable, a & Executable == Executable)
}