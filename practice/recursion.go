package main

import "fmt"

/*
*

	递归函数
	1.概念：递归就是在运行的过程中调用自己
*/
func main() {
	recursion()

}

func recursion() {
	//recursion() /* 函数调用自身 但是没有设置退出条件，则会一直在循环*/

	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	var n int
	for n = 0; n < 10; n++ {
		fmt.Printf("%d\t", fibonacci(n))
	}
}

// 实现阶乘
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// 斐波那契数列 等于前两个之和
func fibonacci(i int) int {
	if i < 2 {
		return i
	}
	return fibonacci(i-2) + fibonacci(i-1)
}
