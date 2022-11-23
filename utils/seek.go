package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Seek
/**
读文件时设置偏移量
	Seek(offset int64, whence int) (ret int64, err error) 设置光标的位置
	第一个参数：偏移量
	第二个参数：如何设置
*/
func Seek() {
	fileName := "E:\\Goland\\src\\Advanced\\utils\\copy.go"
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	// 读写
	bs := []byte{0}
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	_, err = file.Seek(8, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
