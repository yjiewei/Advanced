package main

import "fmt"

/*
*

		数组
		1.需要声明类型和数量 var variable_name [SIZE] variable_type
		2.可以通过索引获取、修改数据等
		3.初始化数组
			var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
			balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}  // 不确定长度，由编译器确认
			// 将索引为 1 和 3 的元素初始化
			balance := [5]float32{1:2.0,3:7.0}
		4.其他：多维数组、数组参数
	      todo 这里还没看呢
	      https://www.runoob.com/go/go-multi-dimensional-arrays.html
		  https://www.runoob.com/go/go-passing-arrays-to-functions.html
*/
func main() {
	var n [10]int /** 长度为10的数组 */
	var i, j, k int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}
