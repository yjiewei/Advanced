package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 并发编程中的共享资源安全问题
var ticket = 10

func Sale() {
	// 模拟4个售票口，即4个goroutine
	go tickets("售票口1")
	go tickets("售票口2")
	go tickets("售票口3")
	go tickets("售票口4")

	// 防止主线程执行完直接结束了，协程还没开始的情况
	time.Sleep(2 * time.Second)
}

func tickets(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		// 线程进入之后条件可能会改变，那输出可能就不对了，是否考虑加锁？go有其他方式处理吗
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出电影票：", ticket)
			ticket--
		} else {
			fmt.Println(name, "已售罄，没有电影票了。")
			break
		}
	}
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
