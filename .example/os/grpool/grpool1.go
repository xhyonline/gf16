package main

import (
	"fmt"
	"time"

	"github.com/xhyonline/gf16/os/grpool"
	"github.com/xhyonline/gf16/os/gtimer"
)

func job() {
	time.Sleep(1 * time.Second)
}

func main() {
	pool := grpool.New(100)
	for i := 0; i < 1000; i++ {
		pool.Add(job)
	}
	fmt.Println("worker:", pool.Size())
	fmt.Println("  jobs:", pool.Jobs())
	gtimer.SetInterval(time.Second, func() {
		fmt.Println("worker:", pool.Size())
		fmt.Println("  jobs:", pool.Jobs())
		fmt.Println()
		gtimer.Exit()
	})

	select {}
}
