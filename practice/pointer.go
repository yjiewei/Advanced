package main

import "fmt"

/*
*

		指针
		1.取地址符是 &
		2.一个指针变量指向了一个值的内存地址。 var var_name *var-type
		3.空指针，声明后没有初始化变量 nil 也代表零值或空值
		4.指针数组 下面的例子我看不出来使用场景是什么
	    5.指向指针的指针：如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
		6.指针作为函数参数：向函数传递指针，只需要在函数定义的参数上设置为指针类型即可
*/
const MAX int = 3

func main() {
	var a int = 10
	fmt.Printf("变量的地址: %x\n", &a) // c00000e0c8

	var ip *int
	ip = &a
	fmt.Printf("a 的变量的地址是：%x\n", &a)

	// 指针变量的存储地址
	fmt.Printf("ip变量存储的指针地址是：%x\n", ip)

	// 使用指针访问值，也就是取出指针指向的值
	fmt.Printf("*ip变量的值：%d\n", *ip)

	var nullpointer *int
	fmt.Printf("空指针：%d \n", nullpointer)

	println(nullpointer == nil)

	b := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &b[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

	ptrptr()

	/* 定义局部变量 */
	var e int = 100
	var f int = 200
	/* 调用函数用于交换两个值
	 * &e 指向 e 变量的地址
	 * &f 指向 f 变量的地址
	 */
	fmt.Printf("交换前 e 的值 : %d\n", e)
	fmt.Printf("交换前 f 的值 : %d\n", f)
	swap(&e, &f)
	fmt.Printf("交换后 e 的值 : %d\n", e)
	fmt.Printf("交换后 f 的值 : %d\n", f)

}

// 地址不变，把里面的值换了 这里的go语言特性
func swap(x *int, y *int) {
	//var temp int
	//temp = *x    /* 保存 x 地址的值 */
	//*x = *y      /* 将 y 赋值给 x */
	//*y = temp    /* 将 temp 赋值给 y */

	*x, *y = *y, *x
}

// 指向指针的指针
func ptrptr() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

	fmt.Printf("指针变量 ptr 的地址 = %d\n", ptr)
	fmt.Printf("指向指针的指针变量 pptr 的地址 = %d\n", pptr)
}
