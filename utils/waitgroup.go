package utils

import "sync"

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func WG() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.设置等待组中要执行的子goroutine的数量
		wg.Add(1)
		// Launch a goroutine to fetch the URL. 启动goroutine
		go func(url string) {
			// Decrement the counter when the goroutine completes. goroutine执行完成时done wg.Add(-1)
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete. 所以协程执行结束counter=0
	wg.Wait()
}
