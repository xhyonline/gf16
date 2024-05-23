package main

import (
	"github.com/xhyonline/gf16/frame/g"
)

func main() {
	g.DB().Model("user").Distinct().CountColumn("uid,name")
}
