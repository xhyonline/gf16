package main

import (
	"time"

	"github.com/xhyonline/gf16/os/gcron"
	"github.com/xhyonline/gf16/os/glog"
)

func test() {
	glog.Println(111)
}

func main() {
	_, err := gcron.AddOnce("@every 2s", test)
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
}
