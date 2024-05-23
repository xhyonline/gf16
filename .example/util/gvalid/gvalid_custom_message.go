package main

import (
	"context"
	"fmt"
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/util/gvalid"
)

func main() {
	g.I18n().SetLanguage("cn")
	err := gvalid.Check(context.TODO(), "", "required", nil)
	fmt.Println(err.String())
}
