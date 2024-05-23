package main

import (
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/net/ghttp"
	"github.com/xhyonline/gf16/os/gtime"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/set", func(r *ghttp.Request) {
			r.Session.Set("time", gtime.Timestamp())
			r.Response.Write("ok")
		})
		group.GET("/get", func(r *ghttp.Request) {
			r.Response.WriteJson(r.Session.Map())
		})
		group.GET("/clear", func(r *ghttp.Request) {
			r.Session.Clear()
		})
	})
	s.SetPort(8199)
	s.Run()
}
