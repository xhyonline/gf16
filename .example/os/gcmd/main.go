package main

import (
	"fmt"

	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/os/gcmd"
)

func main() {
	p, err := gcmd.Parse(g.MapStrBool{
		"n,name":        true,
		"p,prefix":      true,
		"f,force":       false,
		"t,tail":        false,
		"i,interactive": false,
	})
	if err != nil {
		fmt.Println(err)
	}
	g.Dump(p)
}
