package main

import (
	"github.com/xhyonline/gf16/os/gview"
	"github.com/xhyonline/gf16/util/gutil"
)

func main() {
	gutil.Dump(gview.ParseContent(`{{"<div>测试</div>去掉HTML标签<script>var v=1;</script>"|text}}`, nil))
}
