package main

import (
	"context"
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/util/gvalid"
)

// string默认值校验
func main() {
	type User struct {
		Uid string `gvalid:"uid@integer"`
	}

	user := &User{}

	g.Dump(gvalid.CheckStruct(context.TODO(), user, nil))
}
