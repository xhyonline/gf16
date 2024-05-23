package main

import (
	"fmt"

	"github.com/xhyonline/gf16/frame/g"
)

func main() {
	conn := g.Redis().Conn()
	defer conn.Close()
	if _, err := conn.Do("SET", "k", "v"); err != nil {
		panic(err)
	}
	v, _ := conn.DoVar("GET", "k")
	fmt.Println(v.String())
}
