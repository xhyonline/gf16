package main

import (
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"header":    "This is header",
			"container": "This is container",
			"footer":    "This is footer",
		})
	})
	s.SetPort(8199)
	s.Run()
}
