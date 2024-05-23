package main

import (
	"fmt"
	"github.com/xhyonline/gf16/net/ghttp"
	"github.com/xhyonline/gf16/os/glog"
)

func main() {
	path1 := "/Users/john/Pictures/logo1.png"
	path2 := "/Users/john/Pictures/logo2.png"
	r, e := ghttp.Post(
		"http://127.0.0.1:8199/upload",
		fmt.Sprintf(`upload-file=@file:%s&upload-file=@file:%s`, path1, path2),
	)
	if e != nil {
		glog.Error(e)
	} else {
		fmt.Println(string(r.ReadAll()))
		r.Close()
	}
}
