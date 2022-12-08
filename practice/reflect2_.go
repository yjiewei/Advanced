package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%s, 年龄：%d, 性别：%s \n", p.Name, p.Age, p.Gender)
}

func main() {
	person := Person{"yangjiewei", 24, "Male"}
	GetMessage(person)
	person.PrintInfo()
	changePersonValue(person)
	callPersonMethod(person)
}

func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input)   // 获取input类型
	getValue := reflect.ValueOf(input) // 获取input的值

	// 获取类型值
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Printf("字段名：%s，对应值：%v\n", field.Name, value)
	}

	// 获取方法
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称：%s，方法类型：%v\n", method.Name, method.Type)
	}

	// reflect设置实际变量的值 需要是指针
	var change = 3.14
	//changeValue := reflect.ValueOf(change)
	//elem0 := changeValue.Elem()
	//elem0.SetFloat(1.23) // panic: reflect: call of reflect.Value.Elem on float64 Value

	changeValue2 := reflect.ValueOf(&change)
	elem := changeValue2.Elem() // 获取修改的元素
	if elem.CanSet() {
		fmt.Println("change:", change)
		fmt.Println("change变量可以修改值")
		elem.SetFloat(1.23)
		fmt.Println("修改参数值change:", change)
	}
}

// 通过反射修改元素值
func changePersonValue(person Person) {
	fmt.Println(person)
	changePersonValue := reflect.ValueOf(&person) // 必须是指针地址，如果是引用类型则不需要
	if changePersonValue.Kind() == reflect.Ptr {
		elem := changePersonValue.Elem()
		fmt.Println(elem.CanSet())
		name := elem.FieldByName("Name")
		name.SetString("杨杰炜")
		fmt.Println(person) // 改掉了名字，其实我们需要判断字段类型，可能是未知的 用Type方法
	}
}

// 通过反射进行方法调用
func callPersonMethod(person Person) {
	fmt.Println(person)
	callPersonMethod := reflect.ValueOf(&person) // 必须是指针地址，如果是引用类型则不需要
	if callPersonMethod.Kind() == reflect.Ptr {
		elem := callPersonMethod.Elem()
		fmt.Println(elem.CanSet())
		method := elem.MethodByName("PrintInfo")
		fmt.Printf("kind:%s, type:%s\n", method.Kind(), method.Type()) // kind:func, type:func()
		fmt.Println("利用反射调用方法")
		method.Call(nil) // 或者传一个空的切片
		// args:= []reflect.Value{reflect.ValueOf("反射机制")} // 直接声明类型为reflect.Value参数的切片
	}
}
