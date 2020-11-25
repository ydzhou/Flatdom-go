package main

import (
	"github.com/Flatdom-go/flatdom"
)

func main() {
	dom := &flatdom.Dom{}
	dom.Init()
	dom.Run(10)
}
