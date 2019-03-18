package series

import "fmt"

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}
//小写字母开头的函数是包内私有，包外代码不可见
func square(n int) int {
	return n * n
}

//只有首字母大写的才允许保外访问
func GetFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}
