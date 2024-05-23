package main

import (
	"sync"

	"github.com/xhyonline/gf16/os/glog"
	"github.com/xhyonline/gf16/os/gmlock"
)

// 内存锁 - 给定过期时间
func main() {
	key := "lock"
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			gmlock.Lock(key, 1000)
			glog.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
