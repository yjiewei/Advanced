package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10
var wg sync.WaitGroup

// 原语级别的 整了一把锁 这里要定义为全局变量。。。。
var mutex = sync.Mutex{}

func main() {
	wg.Add(4) // 如果数量有多余，会出现死锁问题
	// 模拟4个售票口，即4个goroutine
	go tickets("售票口1")
	go tickets("售票口2")
	go tickets("售票口3")
	go tickets("售票口4")

	// 等待协程执行结束
	wg.Wait()
}

func tickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		// 线程进入之后条件可能会改变，那输出可能就不对了，是否考虑加锁？go有其他方式处理吗
		// 用mutex.Lock()，但是这把锁要加在哪里？如何分析？ticket>0，到ticket--，这里可能被其他线程占用了，所以要在这里加锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出电影票：", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "已售罄，没有电影票了。")
			break
		}
		mutex.Unlock()
	}
}
