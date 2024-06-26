package main

import (
	"github.com/xhyonline/gf16/os/gfsnotify"
	"github.com/xhyonline/gf16/os/glog"
)

// 对同一个文件多次Add是否超过系统inotify限制
func main() {
	path := "/Users/john/temp/log"
	for i := 0; i < 9999999; i++ {
		_, err := gfsnotify.Add(path, func(event *gfsnotify.Event) {
			glog.Println(event)
		})
		if err != nil {
			glog.Fatal(err)
		}
	}
}
