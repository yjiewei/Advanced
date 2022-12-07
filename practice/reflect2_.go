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
}
