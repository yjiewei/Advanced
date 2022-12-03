package main

import "fmt"

func main() {
	var channel chan int
	channel = make(chan int)
	fmt.Printf("通道类型和值分别是：%T, %v\n", channel, channel)
	fmt.Println("通道参数给函数传递的是地址值")

	// 1.如果不用通道去给主协程和子协程通信，可能会导致主协程结束导致程序退出
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println("执行子协程中的循环：", i)
	//	}
	//	fmt.Println("子协程执行结束...")
	//}()
	//fmt.Println("主协程执行结束...")

	// 2.利用通道进行通信
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println("执行子协程中的循环：", i)
	//	}
	//	// 给通道写入数据给主协程去读
	//	channel <- 1
	//	fmt.Println("子协程执行结束...")
	//}()
	//// 读不到数据就会阻塞，所以会等在这里
	//data := <-channel
	//fmt.Println("从通道中读取到的数据是：", data)
	//fmt.Println("主协程执行结束...")

	// 3.通道通信还要关闭通道
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		// 给通道写入数据给主协程去读
	//		channel <- i
	//		//fmt.Println("执行子协程中的循环：", i)
	//	}
	//	fmt.Println("子协程执行结束...")
	//	close(channel)
	//}()
	// 读不到数据就会阻塞，所以会等在这里
	//for {
	//	data, ok := <-channel
	//	fmt.Println("从通道中读取到的数据是：", data, ok)
	//	if !ok {
	//		break
	//	}
	//}

	//// 4.范围循环
	//for v := range channel {
	//	fmt.Println("从通道中读取到的数据是：", v)
	//}
	//fmt.Println("主协程执行结束...")

	// 5.缓冲通道 声明通道的时候同时声明capacity
	//cacheChannel := make(chan int, 5)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		cacheChannel <- i
	//		fmt.Println("往子协程写入数据：", i)
	//	}
	//	fmt.Println("子协程写入数据结束")
	//	// 要关通道，不然主协程会等着读，导致死锁
	//	close(cacheChannel)
	//}()
	//
	//for v := range cacheChannel {
	//	fmt.Println("从通道中读取到的数据是：", v)
	//}
	//fmt.Println("主协程执行结束...")

	// 6.定向通道 也就是单向，其实这个不在声明的时候使用，没有管道是只读或者只写的
	// 所以他通常用在函数声明上，限制函数对通道的操作，实现职责分明
	directChannel := make(chan int)
	go directChannelWrite(directChannel)
	directValue := <-directChannel
	fmt.Println("从通道中读取到的数据是：", directValue)
	fmt.Println("主协程执行结束...")

}

func directChannelWrite(directChannel chan<- int) {

	directChannel <- 1
	fmt.Println("往子协程写入数据：", 1)

	fmt.Println("子协程写入数据结束")
	// 要关通道，不然主协程会等着读，导致死锁
	close(directChannel)
}
