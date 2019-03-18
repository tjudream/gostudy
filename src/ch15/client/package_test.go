package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
//	t.Log(series.square(5)) //小写字母开头的是包内私有的，外部无法访问
}
