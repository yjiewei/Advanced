package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁、同步等待组
var rwMutex *sync.RWMutex
var waitGroup *sync.WaitGroup

// 可以随便读、写的时候其他goroutine不能读写
func main() {
	rwMutex = new(sync.RWMutex)
	waitGroup = new(sync.WaitGroup)
	waitGroup.Add(3)
	//// 读操作可以同时进行
	//go readData(1)
	//go readData(2)

	// 写操作 不能同时进行
	go writeData(1)
	go readData(2)
	go writeData(3)
	waitGroup.Wait()
	fmt.Println("main...over...")
}

func readData(name int) {
	defer waitGroup.Done()
	fmt.Println(name, "开始读操作...")
	rwMutex.RLock() // 读操作加锁，如果写锁先用了，这里要等锁
	fmt.Println(name, "正在读取数据...")
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock()
	fmt.Println(name, "读操作结束...")
}

func writeData(name int) {
	defer waitGroup.Done()
	fmt.Println(name, "开始写操作...")
	rwMutex.Lock() // 写操作加锁
	fmt.Println(name, "正在写入数据...")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
	fmt.Println(name, "写入操作结束...")
}
