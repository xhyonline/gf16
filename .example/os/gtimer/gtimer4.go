package main

import (
	"time"

	"github.com/xhyonline/gf16/os/glog"
	"github.com/xhyonline/gf16/os/gtimer"
)

func main() {
	interval := time.Second
	gtimer.AddTimes(interval, 2, func() {
		glog.Println("doing1")
	})
	gtimer.AddTimes(interval, 2, func() {
		glog.Println("doing2")
	})

	select {}
}
