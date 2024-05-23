package main

import (
	"github.com/xhyonline/gf16/crypto/gaes"
	"github.com/xhyonline/gf16/os/gfile"
	"github.com/xhyonline/gf16/os/gres"
)

var (
	CryptoKey = []byte("x76cgqt36i9c863bzmotuf8626dxiwu0")
)

func main() {
	binContent := gfile.GetBytes("data.bin")
	binContent, err := gaes.Decrypt(binContent, CryptoKey)
	if err != nil {
		panic(err)
	}
	if err := gres.Add(binContent); err != nil {
		panic(err)
	}
	gres.Dump()
}
