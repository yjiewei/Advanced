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

}
