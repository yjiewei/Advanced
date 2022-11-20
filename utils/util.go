package utils

import "fmt"

type Person struct {
	Name  string
	Email string
}

func Test() {
	fmt.Println("测试包的权限，大写方法则能导出...小写不行。")
}

func init() {
	fmt.Println("utils包下的初始化函数，用于初始化一些信息。")
}
