package main

import (
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/net/ghttp"
	"github.com/xhyonline/gf16/os/gtime"
)

func main() {
	s := g.Server()
	s.BindHandler("/cookie", func(r *ghttp.Request) {
		datetime := r.Cookie.Get("datetime")
		r.Cookie.Set("datetime", gtime.Datetime())
		r.Response.Write("datetime:", datetime)
	})
	s.SetPort(8199)
	s.Run()
}
