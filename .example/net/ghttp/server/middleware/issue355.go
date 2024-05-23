package main

import (
	"fmt"
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindMiddlewareDefault(func(r *ghttp.Request) {
		fmt.Println("cors")
		r.Response.CORSDefault()
		r.Middleware.Next()
	})
	s.BindHandler("/api/captcha", func(r *ghttp.Request) {
		r.Response.Write("captcha")
	})
	s.SetPort(8010)
	s.Run()
}
