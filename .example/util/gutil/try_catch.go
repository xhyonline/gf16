package main

import (
	"fmt"

	"github.com/xhyonline/gf16/util/gutil"
)

func main() {
	gutil.TryCatch(func() {
		fmt.Println(1)
		gutil.Throw("error")
		fmt.Println(2)
	}, func(err error) {
		fmt.Println(err)
	})
}
