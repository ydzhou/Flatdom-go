package main

import (
	"github.com/ydzhou/Flatdom-go/flatdom"
)

func main() {
	d := &flatdom.Dom{}
	d.Init()
	d.Run(10)
}
