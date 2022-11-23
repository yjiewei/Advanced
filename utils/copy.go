package utils

import (
	"fmt"
	"io"
	"os"
)

// Copy
/**
拷贝文件
	1.通过io读写文件的拷贝，Read&Write
	2.通过io复制文件实现拷贝，io.Copy()
	3.通过ioutil.WriteFile() ioutil.ReadFile() 一次性读取文件，再一次性写入文件，不适合用于大文件，容易内存溢出。1.16版本已经淘汰，已修改为推荐方式。
*/
func Copy() {
	srcFile := "E:\\Goland\\src\\Advanced\\utils\\copy.go"
	destFile := "copyFile.go"
	//total, err := copyFileByIORW(srcFile, destFile)
	//fmt.Println(total, err)

	//total64, err := copyFileByIOCopy(srcFile, destFile)
	//fmt.Println(total64, err)

	total, err := copyFileByIoutil(srcFile, destFile)
	fmt.Println(total, err)
}

// 1.copy file by io read write
func copyFileByIORW(srcFile, destFile string) (int, error) {
	// 1.1打开文件
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(destFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}

	// 1.2关闭文件
	defer file1.Close()
	defer file2.Close()

	// 1.3读写文件
	bs := make([]byte, 1024, 1024)
	total := 0
	for {
		n, err := file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("文件拷贝完毕。")
			break
		} else if err != nil {
			fmt.Println("文件拷贝出错...")
			return total, err
		}
		total += n
		_, err = file2.Write(bs[:n]) // 写出去
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}

// 2.copy file by io.Copy()
func copyFileByIOCopy(srcFile, destFile string) (int64, error) {
	// 1.1打开文件
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(destFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}

	// 1.2关闭文件
	defer file1.Close()
	defer file2.Close()

	// 1.3复制文件
	return io.Copy(file2, file1)
}

// 3.copy file by ioutil 这方法已经在1.16版本淘汰
func copyFileByIoutil(srcFile, destFile string) (int, error) {
	input, err := /*ioutil.ReadFile(srcFile)*/ os.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}

	err = /*ioutil.WriteFile(destFile, input, 0644)*/ os.WriteFile(destFile, input, os.ModePerm)
	if err != nil {
		return 0, err
	}

	return len(input), nil
}
