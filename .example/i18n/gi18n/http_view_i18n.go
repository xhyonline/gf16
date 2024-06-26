package main

import (
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/i18n/gi18n"
	"github.com/xhyonline/gf16/net/ghttp"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			r.SetCtx(gi18n.WithLanguage(r.Context(), r.GetString("lang", "zh-CN")))
			r.Middleware.Next()
		})
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.WriteTplContent(`{#hello}{#world}!`)
		})
	})
	s.SetPort(8199)
	s.Run()
}
