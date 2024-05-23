package main

import (
	"time"

	"github.com/xhyonline/gf16/net/gtcp"
	"github.com/xhyonline/gf16/os/glog"
	"github.com/xhyonline/gf16/util/gconv"
)

func main() {
	// Client
	conn, err := gtcp.NewConn("127.0.0.1:8999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for i := 0; i < 3; i++ {
		if err := conn.Send([]byte(gconv.String(i))); err != nil {
			glog.Error(err)
		}
		time.Sleep(time.Second)
	}
}
