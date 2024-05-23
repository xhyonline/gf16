package main

import (
	"errors"

	"github.com/xhyonline/gf16/os/glog"

	"github.com/xhyonline/gf16/errors/gerror"
)

func Error1() error {
	return errors.New("test1")
}

func Error2() error {
	return gerror.New("test2")
}

func main() {
	glog.Println(Error1())
	glog.Println(Error2())
}
