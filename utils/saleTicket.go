package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// fixme 加了锁之后，原来用的等待组用起来不对。导致不能正常加锁？
// 并发编程中的共享资源安全问题 https://pkg.go.dev/sync#Mutex
var ticket = 10
var wg sync.WaitGroup

func Sale() {
	//wg.Add(4) // 如果数量有多余，会出现死锁问题
	// 模拟4个售票口，即4个goroutine
	go tickets("售票口1")
	go tickets("售票口2")
	go tickets("售票口3")
	go tickets("售票口4")

	// 防止主线程执行完直接结束了，协程还没开始的情况
	time.Sleep(2 * time.Second)

	// 用同步等待的方法取代上面的线程睡眠，但是并没有解决共享变量的问题，上锁之后再用这个，锁没有生效
	/*	fmt.Println("如果先输出这行，说明主线程已经执行完了，我们要等等。")
		wg.Wait()
		fmt.Println("协程也执行完毕了，那就放行吧。")*/

}

func tickets(name string) {
	rand.Seed(time.Now().UnixNano())
	// 原语级别的 整了一把锁
	mutex := sync.Mutex{}
	for {
		// 线程进入之后条件可能会改变，那输出可能就不对了，是否考虑加锁？go有其他方式处理吗
		// 用mutex.Lock()，但是这把锁要加在哪里？如何分析？ticket>0，到ticket--，这里可能被其他线程占用了，所以要在这里加锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出电影票：", ticket)
			ticket--
		} else {
			fmt.Println(name, "已售罄，没有电影票了。")
			break
		}
		mutex.Unlock()
	}
	//wg.Done()
}

/**
售票口1 售出电影票： 10
售票口2 售出电影票： 9
售票口3 售出电影票： 9
售票口4 售出电影票： 9
售票口1 售出电影票： 6
售票口3 售出电影票： 5
售票口1 售出电影票： 4
售票口1 售出电影票： 3
售票口2 售出电影票： 2
售票口3 售出电影票： 1
售票口3 已售罄，没有电影票了。
售票口2 售出电影票： 0
售票口2 已售罄，没有电影票了。
售票口4 售出电影票： -1
售票口4 已售罄，没有电影票了。
售票口1 售出电影票： -2
售票口1 已售罄，没有电影票了。
*/
