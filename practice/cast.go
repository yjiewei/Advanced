package main

import "fmt"

/*
*

	类型转换
	1.定义：类型转换用于将一种数据类型的变量转换为另外一种类型的变量
	2.type_name(expression)
	3.go 不支持隐式转换类型
*/
func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	intDivide := sum / count
	fmt.Printf("intDivide 的值为: %d\n", intDivide)
	// sum int转float32
	mean = float32(sum) / float32(count)
	fmt.Printf("float32类型转换之后的mean 的值为: %f\n", mean)

	float64mean := float64(sum) / float64(count)
	fmt.Printf("float64mean 的值为: %f\n", float64mean)
}
