// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/xhyonline/gf16.

package ghttp_test

import (
	"context"
	"fmt"
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/net/ghttp"
	"github.com/xhyonline/gf16/test/gtest"
	"testing"
	"time"
)

func Test_Request_SetCtx(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			ctx := context.WithValue(r.Context(), "test", 1)
			r.SetCtx(ctx)
			r.Middleware.Next()
		})
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.Write(r.Context().Value("test"))
		})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(c.GetContent("/"), "1")
	})
}
