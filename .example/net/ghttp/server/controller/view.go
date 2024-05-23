package main

import (
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/frame/gmvc"
)

type Controller struct {
	gmvc.Controller
}

func (c *Controller) Index() {
	c.View.Assign("name", "john")
	c.View.Assign("mainTpl", "main/main2.html")
	c.View.Display("layout.html")
}

func main() {
	s := g.Server()
	s.BindController("/view", new(Controller))
	s.SetPort(8199)
	s.Run()
}
