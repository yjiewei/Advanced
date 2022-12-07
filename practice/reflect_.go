package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1.反射操作：通过反射，可以获取一个接口类型变量的类型和数值
	var x = /*float64*/ 3.14 // 空接口的实现
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("type:", reflect.ValueOf(x))

	fmt.Println("nil type:", reflect.TypeOf(nil))
	fmt.Println("nil type:", reflect.ValueOf(nil))

	reflectX := reflect.ValueOf(x)
	fmt.Println("kind:", reflectX.Kind())
	fmt.Println("value:", reflectX.Float())

	var str = "杨杰炜"
	reflectStr := reflect.ValueOf(str) // 反射获取值和类型
	fmt.Println("kind:", reflectStr.Kind())
	fmt.Println("value:", reflectStr.String())

	typeStr := reflect.TypeOf(str)
	fmt.Println("kind:", typeStr.Kind())
	fmt.Println("value:", typeStr.Name())

	fmt.Println("----------------------------")
	var num = 1.23
	// 接口类型变量->反射类型变量
	value := reflect.ValueOf(num)

	// 反射类型对象->接口类型变量
	convertValue := value.Interface()
	// convertValue := value.Interface().(float64)
	fmt.Println("反射类型对象->接口类型变量:", convertValue)

	// 反射类型对象->接口类型变量，理解为强制转换  golang对类型要求非常严格
	pointer := reflect.ValueOf(&num)
	//convertPointer := pointer.Interface()
	convertPointer := pointer.Interface().(float64) // panic: interface conversion: interface {} is *float64, not float64 // 类型对不上，这里是指针
	fmt.Println("反射类型对象->接口类型变量，理解为强制转换:", convertPointer)
}
