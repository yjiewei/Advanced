package utils

import (
	"fmt"
	"log"
	"os"
)

// IOFunc
/**
读写文件操作
1.Reader接口：打开文件、读取数据、关闭文件 Read(b []byte) (n int, err error)
2.Writer接口：打开文件、写入数据、关闭文件 Write(b []byte) (n int, err error)
*/
func IOFunc() {
	read()
	write()
}

func write() {
	// 1.open file
	fileName := "E:\\Goland\\src\\Advanced\\utils\\io.go"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_RDONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// 3.close file
	defer file.Close()

	// 2.write data
	bs := []byte{65, 66, 67, 68, 69, 70}
	_, err = file.Write(bs)
	HandleErr(err)
	file.WriteString("\n")

	// 可以直接写出字符串 file.WriteString("hello world") 也可以将字符串转字节数组 []byte("hello world")
	_, err = file.WriteString("hello world")
	HandleErr(err)
	file.WriteString("\n")

	_, err = file.Write([]byte("hello world"))
	HandleErr(err)
	fmt.Println()

}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
