package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.正常执行
	//timer := time.NewTimer(3 * time.Second)
	//fmt.Println(time.Now()) // 2022-12-03 17:40:46.123744 +0800 CST m=+0.003269901
	//fmt.Println(<-timer.C)  // 2022-12-03 17:40:49.1306716 +0800 CST m=+3.010197501

	// 2.中断
	//timer2 := time.NewTimer(5 * time.Second)
	//go func() {
	//	<-timer2.C
	//	fmt.Println("子协程timer结束了。")
	//}()
	//time.Sleep(3 * time.Second)
	//flag := timer2.Stop()
	//if flag {
	//	fmt.Println("中断timer直接结束。")
	//}

	// 3.after操作 返回的是通道，相当于time.NewTimer(5 * time.Second).C
	ch := time.After(5 * time.Second)
	<-ch
	fmt.Println("主协程执行结束了。")
}
