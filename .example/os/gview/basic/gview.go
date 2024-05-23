package main

import (
	"fmt"

	"github.com/xhyonline/gf16/frame/g"
)

func main() {
	v := g.View()
	b, err := v.Parse("gview.tpl", map[string]interface{}{
		"k": "v",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
