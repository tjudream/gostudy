package condition

import "testing"

/**
08 条件和循环
*/

func TestIfMultiSec(t *testing.T) {
	if a := 1==1; a {
		t.Log("1==1")
	}
	//if v,err := someFun(); err==nil {
	//	t.Log(v)
	//} else {
	//	t.Log("there is an error")
	//}
}

/**
	// go 的switch的case中不需有break，go会默认break
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		//break
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	switch {
		case 0 <= Num && Num <= 3:
			fmt.Printf("0-3")
		case 4 <= Num && Num <=6:
			fmt.Printf("4-6")
		case 7 <= Num && Num <= 9:
			fmt.Printf("7-9")
	}
*/

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0,2:
			t.Log("Even")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := -1; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}