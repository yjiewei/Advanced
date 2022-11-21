package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// FileFunc
/**
文件操作
1.路径：相对路径、绝对路径、当前目录(.)、上一层(..)
2.目录：创建文件夹(os.Mkdir()/os.MkdirAll())、创建文件、打开文件、删除目录
*/
func FileFunc() {
	// 1.路径
	fileName1 := "E:\\Goland\\src\\Advanced\\utils\\file.go"
	fileName2 := "file.go"
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))

	fmt.Println("获取上一级目录：", path.Join(fileName1, ".."))

	// 2.创建目录
	//err := os.Mkdir("E:\\Goland\\src\\Advanced\\utils\\fileDir", os.ModePerm)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("文件夹创建成功。")

	// 创建多层级的文件夹要用 os.MkdirAll()
	//err := os.MkdirAll("E:\\Goland\\src\\Advanced\\utils\\fileDir\\aa", os.ModePerm)
	//if err != nil {
	//	// err: mkdir E:\Goland\src\Advanced\utils\fileDir\aa: The system cannot find the path specified.
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("多层级文件夹创建成功。")

	// 3.创建文件 如果是已存在文件则会被截断
	//_, err := os.Create("E:\\Goland\\src\\Advanced\\utils\\file.txt")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("创建文件成功。")

	// 4.打开文件
	file1, err := os.Open(fileName1) // 只读的方式
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file1)
	_ = file1.Close()

	file2, err := os.OpenFile(fileName1, os.O_RDONLY|os.O_WRONLY, os.ModePerm) // 读写的方式 如果是新建则 777
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file2)
	_ = file2.Close()

	// 5.移除空目录、文件
	//err := os.Remove(fileName1) // os.RemoveAll()
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
}
