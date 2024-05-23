package main

import (
	"time"

	"github.com/xhyonline/gf16/os/gcron"
	"github.com/xhyonline/gf16/os/glog"
)

func main() {
	gcron.SetLogLevel(glog.LEVEL_ALL)
	gcron.Add("* * * * * ?", func() {
		glog.Println("test")
	})
	time.Sleep(3 * time.Second)
}
