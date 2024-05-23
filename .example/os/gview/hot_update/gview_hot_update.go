package main

import (
	"fmt"
	"time"

	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/os/gtime"
)

func main() {
	v := g.View()
	v.SetPath(`D:\Workspace\Go\GOPATH\src\gitee.com\johng\gf\geg\os\gview`)
	gtime.SetInterval(time.Second, func() bool {
		b, _ := v.Parse("gview.tpl", nil)
		fmt.Println(string(b))
		return true
	})
	select {}
}
