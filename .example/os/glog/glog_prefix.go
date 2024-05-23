package main

import (
	"github.com/xhyonline/gf16/os/glog"
)

func main() {
	l := glog.New()
	l.SetPrefix("[API]")
	l.Println("hello world")
	l.Error("error occurred")
}
