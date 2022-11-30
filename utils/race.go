package utils

import (
	"fmt"
	"time"
)

func FuncRace() {
	// 定义临界资源
	a := 1
	go func() {
		a = 2
		fmt.Println("goroutine中的a=", a)
	}()

	a = 3
	time.Sleep(2 * time.Second)
	fmt.Println("main goroutine中的a=", a)
}
