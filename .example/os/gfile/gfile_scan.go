package main

import (
	"github.com/xhyonline/gf16/os/gfile"
	"github.com/xhyonline/gf16/util/gutil"
)

func main() {
	gutil.Dump(gfile.ScanDir("/Users/john/Documents", "*.*"))
	gutil.Dump(gfile.ScanDir("/home/john/temp/newproject", "*", true))
}
