package main

import (
	"fmt"

	"github.com/xhyonline/gf16/text/gstr"
)

func main() {
	fmt.Println(gstr.SubStr("我是中国人", 2))
	fmt.Println(gstr.SubStr("我是中国人", 2, 2))
}
