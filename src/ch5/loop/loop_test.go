package loop

import "testing"

/**
08 条件和循环
go 中没有while， 只有 for
 */
func TestWhileLoop(t *testing.T) {
	n := 0
	/* while(n<5) */
	for n < 5 {
		t.Log(n)
		n++
	}

	//无限循环
	//for {
	//}
}


