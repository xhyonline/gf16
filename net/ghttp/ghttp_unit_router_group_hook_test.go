// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/test/gtest"
)

func Test_Router_Group_Hook1(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	group := s.Group("/api")
	group.GET("/handler", func(r *ghttp.Request) {
		r.Response.Write("1")
	})
	group.ALL("/handler", func(r *ghttp.Request) {
		r.Response.Write("0")
	}, ghttp.HookBeforeServe)
	group.ALL("/handler", func(r *ghttp.Request) {
		r.Response.Write("2")
	}, ghttp.HookAfterServe)

	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(client.GetContent("/api/handler"), "012")
		t.Assert(client.PostContent("/api/handler"), "02")
		t.Assert(client.GetContent("/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Hook2(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	group := s.Group("/api")
	group.GET("/handler", func(r *ghttp.Request) {
		r.Response.Write("1")
	})
	group.GET("/*", func(r *ghttp.Request) {
		r.Response.Write("0")
	}, ghttp.HookBeforeServe)
	group.GET("/*", func(r *ghttp.Request) {
		r.Response.Write("2")
	}, ghttp.HookAfterServe)

	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(client.GetContent("/api/handler"), "012")
		t.Assert(client.PostContent("/api/handler"), "Not Found")
		t.Assert(client.GetContent("/api/ThisDoesNotExist"), "02")
		t.Assert(client.PostContent("/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Hook3(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/api").Bind([]g.Slice{
		{"ALL", "handler", func(r *ghttp.Request) {
			r.Response.Write("1")
		}},
		{"ALL", "/*", func(r *ghttp.Request) {
			r.Response.Write("0")
		}, ghttp.HookBeforeServe},
		{"ALL", "/*", func(r *ghttp.Request) {
			r.Response.Write("2")
		}, ghttp.HookAfterServe},
	})

	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(client.GetContent("/api/handler"), "012")
		t.Assert(client.PostContent("/api/handler"), "012")
		t.Assert(client.DeleteContent("/api/ThisDoesNotExist"), "02")
	})
}
