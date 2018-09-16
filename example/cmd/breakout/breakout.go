package main

import (
	"github.com/jakecoffman/gam"
	"github.com/jakecoffman/gam/example/breakout"
)

const (
	width  = 800
	height = 600
)

func main() {
	Breakout := &breakout.Game{}
	gam.Run(Breakout, 1./60, width, height)
}
