package main

import (
	"context"
	"fmt"
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/i18n/gi18n"
)

func main() {
	var (
		orderId     = 865271654
		orderAmount = 99.8
	)
	fmt.Println(g.I18n().Tf(
		gi18n.WithLanguage(context.TODO(), `en`),
		`{#OrderPaid}`, orderId, orderAmount,
	))
	fmt.Println(g.I18n().Tf(
		gi18n.WithLanguage(context.TODO(), `zh-CN`),
		`{#OrderPaid}`, orderId, orderAmount,
	))
}
