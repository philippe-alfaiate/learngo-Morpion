package main

import (
	"github.com/faiface/pixel/pixelgl"
)

var nonQuitté bool

func main() {
	nonQuitté = true
	for nonQuitté {
		pixelgl.Run(run)
	}

}

func run() {

}
