package main

import (
	"fmt"

	"github.com/xhyonline/gf16/frame/g"

	_ "github.com/xhyonline/gf16/os/gres/testdata"
)

func main() {
	m := g.I18n()
	m.SetLanguage("ja")
	err := m.SetPath("/i18n-dir")
	if err != nil {
		panic(err)
	}
	fmt.Println(m.Translate(`hello`))
	fmt.Println(m.Translate(`{#hello}{#world}!`))
}
