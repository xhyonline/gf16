package main

import (
	"fmt"
	"github.com/xhyonline/gf16/net/ghttp"
)

func main() {
	r, err := ghttp.Get("http://127.0.0.1:8199/11111/11122")
	fmt.Println(err)
	fmt.Println(r.Header)
}
