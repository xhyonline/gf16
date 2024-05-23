package main

import (
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writef("name: %v, pass: %v", r.Get("name"), r.Get("pass"))
	})
	s.SetPort(8199)
	s.Run()
}
