package main

import (
	"fmt"
	"unsafe"
)

/*
	常量
	1.在程序运行时不会被修改的量，没有指定type类型的则为隐式类型定义
		const identifier [type] = value
	2.可以用作枚举
	3.iota：特殊常量，可以认为是一个可以被编译器修改的常量。常量里面的累加值
	4.我感觉这个输出格式化似乎有些问题
*/

const (
	d = "abc"
	e = len(d)
	f = unsafe.Sizeof(e) // 3个字符为什么是8个字节
)

const (
	oo = 1 << iota
	p  = 3 << iota
	q  // iota从0开始，到这是2，q=3<<2=12
	r
)

func main() {
	// 整型常量
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d \n", area)
	println()
	println(a, b, c)
	println(d, e, f)

	test()

	fmt.Println("oo=", oo)
	fmt.Println("p=", p)
	fmt.Println("q=", q)
	fmt.Println("r=", r)
}

func test() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	println()
	// 0 1 2 ha ha 100 100 7 8
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
