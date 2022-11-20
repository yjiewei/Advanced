package main

import "fmt"

///* 定义接口 */
//type interface_name interface {
//	method_name1 [return_type]
//	method_name2 [return_type]
//	method_name3 [return_type]
//...
//method_namen [return_type]
//}
//
///* 定义结构体 */
//type struct_name struct {
//	/* variables */
//}
//
///* 实现接口方法 */
//func (struct_name_variable struct_name) method_name1() [return_type] {
///* 方法实现 */
//}
//...
//func (struct_name_variable struct_name) method_namen() [return_type] {
///* 方法实现*/
//}
/**
接口：我感觉有点难
定义：Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
*/

// 1.定义了phone接口
type Phone interface {
	call()
}

// 2.定义结构体去实现接口方法，这里不一定是要结构体吧
type NokiaPhone struct{}

type IPhone struct{}

// 3.方法实现 在方法函数的基础上，加限制类型？不是返回值
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("我是诺基亚，我想打电话给你...")
}

func (iPhone IPhone) call() {
	fmt.Println("我是苹果，我想打电话给你...")
}

func main() {
	// 直接使用实现方法
	var phone NokiaPhone
	phone.call()

	// 接口对象
	var interfacePhone Phone
	interfacePhone = new(IPhone)
	interfacePhone.call()
}
