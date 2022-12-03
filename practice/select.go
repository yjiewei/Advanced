package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case <-ch1:
		fmt.Println("case1可以执行...")
	case <-ch2:
		fmt.Println("case2可以执行...")
	case <-time.After(3 * time.Second):
		fmt.Println("case3可以执行...超时3秒")
	}
}
