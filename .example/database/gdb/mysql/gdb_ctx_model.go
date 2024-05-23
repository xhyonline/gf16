package main

import (
	"context"
	"github.com/xhyonline/gf16/frame/g"
)

func main() {
	ctx := context.WithValue(context.Background(), "Trace-Id", "123456789")
	_, err := g.DB().Model("user").Ctx(ctx).All()
	if err != nil {
		panic(err)
	}
}
