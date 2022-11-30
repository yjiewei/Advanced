package utils

import (
	"fmt"
	"runtime"
)

func RuntimeFunc() {
	fmt.Println("runtime可以输出运行时信息,GOROOT:", runtime.GOROOT())
	fmt.Println("runtime可以输出运行时信息,版本信息:", runtime.Version())
	fmt.Println("逻辑CPU的数量:", runtime.NumCPU())
	fmt.Println("go程序执行的最大CPU数量:", runtime.GOMAXPROCS(runtime.NumCPU()))
	// 还有其他Goexit、Gosched等方法，对协程进行操作
	// https://pkg.go.dev/runtime#Gosched
}
