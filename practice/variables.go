package main

import "fmt"

/**
    0.标识符：字母、数字、下划线，第一个字符不能是数字
	1.fmt包提供多种格式化方法
	2.数据类型：布尔型、整型、浮点型、复数、字符串、派生类型（指针、数组、结构化、channel、函数、切片、接口、map）
	3.变量：未初始化变量默认零值、派生类型则为nil、初始化了值但没有类型会自动判断类型、:= 是短变量声明语句、多变量声明
    4.值类型、引用类型：基本类型int/float/bool/string属于值类型，直接指向存在内存中的值；而引用类型存放的则是值的地址
	5.空白标识符：_, b = 1, 2 表示有这么个变量，但是我不打算用他的值，因为不使用但是要接收，所以用空白标识符接收
*/

// 全局变量
var h, i int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	j int
	k bool
)

var l, m int = 1, 2
var n, o = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//p, q := 123, "hello"

func main() {
	// 未初始化变量
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) // 0 0 false ""

	// 自动判断类型变量
	var d = true
	fmt.Println(d)

	// 短变量 类似局部变量
	e := "Runoob" // var e string = "Runoob"
	fmt.Println(e)

	// 输出全局变量
	p, q := 123, "hello"
	println(h, i, j, k, l, m, n, o, p, q)

	_, _, strs := numbers()
	println(strs)

}

// 一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
