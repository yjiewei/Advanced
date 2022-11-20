package main

import "fmt"

/*
*

			函数
			1.func function_name( [parameter list] ) [return_types] {
		   		函数体
			  }
			2.默认使用值传递
			3.可以返回多个值
			4.函数用法：作为另外一个函数的参数、闭包（匿名函数）、方法
	          todo:闭包需要看 https://www.runoob.com/go/go-function-closures.html
*/
func main() {
	println("函数使用")
	result, _ := max(1, 2)
	fmt.Printf("最大值是：%d \n", result)
}

func max(num1, num2 int) (int, int) {
	/** 声明局部变量 */
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result, num2
}
