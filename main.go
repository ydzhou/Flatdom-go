package main

import (
	"github.com/ydzhou/flatdom"
)

func main() {
	dom := &flatdom.Dom{}
	dom.Init()
	dom.Run(10)
}
