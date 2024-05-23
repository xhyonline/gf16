package main

import (
	"fmt"
	"github.com/xhyonline/gf16/util/guid"
)

func main() {
	for i := 0; i < 100; i++ {
		s := guid.S()
		fmt.Println(s, len(s))
	}
}
