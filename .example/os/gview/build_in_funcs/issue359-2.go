package main

import (
	"fmt"

	"github.com/xhyonline/gf16/frame/g"
)

func main() {
	s := "我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人"
	tplContent := `
{{.str | strlimit 10  "..."}}
`
	content, err := g.View().ParseContent(tplContent, g.Map{
		"str": s,
	})
	fmt.Println(err)
	fmt.Println(content)
}
