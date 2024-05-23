package main

import (
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln(r.Get("amount"))
		r.Response.Writeln(r.GetInt("amount"))
		r.Response.Writeln(r.GetFloat32("amount"))
	})
	s.SetPort(8199)
	s.Run()
}
