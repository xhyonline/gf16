package main

import (
	"fmt"

	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/os/gfile"
)

func main() {
	v := g.View()
	v.AddPath(gfile.MainPkgPath() + gfile.Separator + "template")
	b, err := v.Parse("index.html", map[string]interface{}{
		"k": "v",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
