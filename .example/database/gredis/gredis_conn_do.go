package main

import (
	"fmt"

	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/util/gconv"
)

func main() {
	conn := g.Redis().Conn()
	defer conn.Close()
	conn.Do("SET", "k", "v")
	v, _ := conn.Do("GET", "k")
	fmt.Println(gconv.String(v))
}
