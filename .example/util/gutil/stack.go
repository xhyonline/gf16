package main

import (
	"github.com/xhyonline/gf16/util/gutil"
)

func Test(s *interface{}) {
	//debug.PrintStack()
	gutil.PrintStack()
}

func main() {
	Test(nil)
}
