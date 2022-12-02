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
	go func() {
		for i := 0; i < 10; i++ {
			// 给通道写入数据给主协程去读
			channel <- i
			//fmt.Println("执行子协程中的循环：", i)
		}
		fmt.Println("子协程执行结束...")
		close(channel)
	}()
	// 读不到数据就会阻塞，所以会等在这里
	//for {
	//	data, ok := <-channel
	//	fmt.Println("从通道中读取到的数据是：", data, ok)
	//	if !ok {
	//		break
	//	}
	//}

	// 4.范围循环
	for v := range channel {
		fmt.Println("从通道中读取到的数据是：", v)
	}
	fmt.Println("主协程执行结束...")
}
