package main

import (
	"fmt"

	"github.com/xhyonline/gf16/encoding/gcompress"
	"github.com/xhyonline/gf16/os/gfile"
)

func main() {
	err := gcompress.UnZipContent(
		gfile.GetBytes(`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg\encoding\gcompress\data.zip`),
		`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg`,
	)
	fmt.Println(err)
}
