package main

import "fmt"

/**
函数作用域
1.函数内定义的变量称为局部变量
2.函数外定义的变量称为全局变量
3.函数定义中的变量称为形式参数
*/

/* 声明全局变量 */
var g int

func main() {
	localVariable()
	globalVariable()
	formalParameter()
}

func formalParameter() {
	// helloworld/practice/function.go:20
}

func globalVariable() {
	/* 声明局部变量 */
	var a, b int

	/* 初始化参数 */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("全局变量结果： a = %d, b = %d and g = %d\n", a, b, g)
}

func localVariable() {
	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("局部变量结果： a = %d, b = %d and c = %d\n", a, b, c)
}
