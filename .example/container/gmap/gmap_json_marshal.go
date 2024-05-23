package main

import (
	"encoding/json"
	"fmt"
	"github.com/xhyonline/gf16/frame/g"

	"github.com/xhyonline/gf16/container/gmap"
)

func main() {
	m := gmap.New()
	m.Sets(g.MapAnyAny{
		"name":  "john",
		"score": 100,
	})
	b, _ := json.Marshal(m)
	fmt.Println(string(b))
}
