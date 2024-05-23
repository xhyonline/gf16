package main

import (
	"fmt"
	"github.com/xhyonline/gf16/encoding/gini"
)

func main() {
	s := `
a = b

`
	m, err := gini.Decode([]byte(s))
	fmt.Println(err)
	fmt.Println(m)
}
