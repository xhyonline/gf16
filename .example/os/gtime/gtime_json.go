package main

import (
	"encoding/json"
	"fmt"

	"github.com/xhyonline/gf16/os/gtime"
)

func main() {
	t := gtime.Now()
	b, err := json.Marshal(t)
	fmt.Println(err)
	fmt.Println(string(b))
}
