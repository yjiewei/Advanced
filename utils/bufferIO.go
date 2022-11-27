package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// BufferIO
/*
	bufio包原理：通过缓冲提高效率，io本身的效率不低，但是频繁访问本地磁盘文件会很低。bufio提供缓冲区，读写现在缓冲区，然后再读写文件到磁盘减少磁盘iO。
	ReadBytes()
	ReadString()
	ReadLine()
	Write()
	WriteByte()
	WriteString()
*/
func BufferIO() {
	fileName := "E:\\Goland\\src\\Advanced\\utils\\bufferIO.go"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// 1.创建Reader对象 Read()高效读取
	b1 := bufio.NewReader(file)
	//p := make([]byte, 1024) // 默认的缓冲区是4096，如果切片比这个值大，则缓冲区没有意义
	//n1, err := b1.Read(p)
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	// 2.ReadLine() 读行不推荐
	data, flag, err := b1.ReadLine()
	fmt.Println(flag)
	fmt.Println(err)
	fmt.Println(data)
	fmt.Println(string(data))

	// 3.读字符串，通过分隔符划分
	//s1, err := b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)

	// 4.scanner
	// fmt包下的scanner不能读取空格后的字符
	// 我已经回想起被C语言支配的恐惧了。。。
	//s2 := ""
	//fmt.Scanln(&s2)
	//fmt.Println(s2)

	//b2 := bufio.NewReader(os.Stdin)
	//s2, err := b2.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s2)

	// 5.写数据
	// 关于写数据的缓冲区和切片规则有点多，用的时候再去网上查吧，死记硬背不管用的
	fileName = "E:\\Goland\\src\\Advanced\\utils\\bufferIO.txt"
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	w1 := bufio.NewWriter(file)
	n2, err := w1.WriteString("hello world \n")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n2)
	w1.Flush() // 刷新缓冲区，如果没有刷新，数据不会到文件，因为缓冲区还没有刷磁盘

	for i := 1; i < 1000; i++ {
		// 1:1000 所占用的字节作为p
		_, err := w1.WriteString(fmt.Sprintf(" %d:hello \n", i))
		if err != nil {
			return
		}
	}
	// 如果不加这一行，前面的缓冲区写满之后到磁盘，但是还有一些数据还在缓存
	w1.Flush()
}
