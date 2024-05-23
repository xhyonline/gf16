package main

import (
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("halo 世界！")
	})
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.Writeln("This is customized 404 page")
	})
	s.SetPort(8199)
	s.Run()
}
