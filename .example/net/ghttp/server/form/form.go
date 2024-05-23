package main

import (
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		g.Dump(r.GetForm("array"))
		r.Response.WriteTpl("form.html")
	})
	s.SetPort(8199)
	s.Run()
}
