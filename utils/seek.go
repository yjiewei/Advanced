package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Seek
/**
读文件时设置偏移量
	Seek(offset int64, whence int) (ret int64, err error) 设置光标的位置
	第一个参数：偏移量
	第二个参数：如何设置
断点续传
	问题：1.传的文件比较大，是否有方法可以缩短耗时 2.如果传递过程中，被迫中断，下次还需要重新开始吗？ 3.传递文件的时候，支持暂停和恢复吗？
	解决：可以通过断点续传操作恢复到读写的位置，在继续传递数据。主要是记住上一次已经传递了多少数据(通过临时文件的方式记录)。
*/
func Seek() {
	// seekOperate()
	copyFileBySeek()
}

// 断点续传
func copyFileBySeek() {
	srcFile := "E:\\Goland\\src\\Advanced\\utils\\Go路线.png"
	destFile := srcFile[strings.LastIndex(srcFile, "\\")+1:]
	fmt.Println("复制文件到：", destFile)
	tempFile := destFile + "temp.txt"
	fmt.Println("临时文件：", tempFile)

	file1, err := os.Open(srcFile)
	HandleErr(err)
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandleErr(err)
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_WRONLY|os.O_RDONLY, os.ModePerm)
	HandleErr(err)

	defer file1.Close()
	defer file2.Close()

	// 1.先读取临时文件中的数据，再seek
	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 1024, 1024)
	n1, err := file3.Read(bs)
	countStr := string(bs[:n1])
	count, err := strconv.ParseInt(countStr, 10, 64)
	//HandleErr(err)
	fmt.Println(count)

	// 2.设置读，写的位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1024)
	//读、写、读取的数据量
	n2 := -1
	n3 := -1
	total := int(count)

	// 3.复制文件
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件已经复制完了。")
			file3.Close()
			os.Remove(tempFile)
			return
		}
		n3, err = file2.Write(data[:n2])
		total += n3

		// 将复制到的文件存储到临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		//if total > 200 {
		//	fmt.Println("模仿断电操作...")
		//	return
		//}
	}
}

func seekOperate() {
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
