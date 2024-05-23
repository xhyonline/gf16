package main

import (
	"time"

	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
)

func main() {
	gcron.SetLogLevel(glog.LEVEL_ALL)
	gcron.Add("* * * * * ?", func() {
		glog.Println("test")
	})
	time.Sleep(3 * time.Second)
}
