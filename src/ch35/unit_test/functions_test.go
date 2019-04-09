package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T)  {
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, but the actual is %d",
				inputs[i], expected[i], ret)
		}
	}
}

func TestErrorInCode(t *testing.T)  {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFatalInCode(t *testing.T)  {
	fmt.Println("Start")
	t.Fatal("Fatal")
	fmt.Println("End")
}

func TestSquareAssert(t *testing.T)  {
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}
