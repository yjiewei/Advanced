package utils

import (
	"fmt"
	"os"
)

// IOFunc
/**
读写文件操作
1.Reader接口：打开文件、读取数据、关闭文件  Read(b []byte) (n int, err error)
*/
func IOFunc() {
	read()
}

func read() {
	// 1.open file
	fileName := "E:\\Goland\\src\\Advanced\\utils\\io.go"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 3.close file
	defer file.Close()

	// 2.read data
	bs := make([]byte, 4, 4)
	for {
		_, err := file.Read(bs)
		if err != nil {
			fmt.Println("已经读到文件末尾啦。")
			break
		}
		//fmt.Println("err:", err)
		//fmt.Println("n:", n)
		fmt.Printf(string(bs))
	}
}
