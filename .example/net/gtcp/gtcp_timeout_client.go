package main

import (
	"time"

	"github.com/xhyonline/gf16/net/gtcp"
	"github.com/xhyonline/gf16/os/glog"
	"github.com/xhyonline/gf16/os/gtime"
)

func main() {
	conn, err := gtcp.NewConn("127.0.0.1:8999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if err := conn.Send([]byte(gtime.Now().String())); err != nil {
		glog.Error(err)
	}

	time.Sleep(time.Minute)
}
