package main

import (
	"fmt"

	"github.com/xhyonline/gf16/os/glog"
)

func main() {

	glog.PrintStack()
	glog.New().PrintStack()

	fmt.Println(glog.GetStack())
	fmt.Println(glog.New().GetStack())
}
