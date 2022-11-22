package main

import (
	"Advanced/utils"
	"fmt"
)

/*
*
关于包的使用：

	1.一个目录下的统计文件归属一个包，package的声明要一致
	2.package声明的包和对应的目录名可以不一致，但习惯上还是写一致的
	3.包可以嵌套，写在多个文件夹里面
	4.同包下的函数不需要导入包，可以直接使用
	5.main包，main()函数所在的包，其他的包不能使用
	6.导入包时，路径要从src下开始写
*/
func main() {
	// time()
	// file()
	io()
}

func time() {
	utils.Test()
	person := utils.Person{
		Name:  "杨杰炜",
		Email: "yang_7131@163.com",
	}
	fmt.Println("个人姓名：" + person.Name + " 联系方式：" + person.Email)

	utils.TimeFunc()
}

func file() {
	utils.FileFunc()
}

func io() {
	utils.IOFunc()
}
