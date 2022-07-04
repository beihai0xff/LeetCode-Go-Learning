package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		x, y := robotgo.GetMousePos()
		robotgo.MoveSmooth(x-1, y-1)
		// Sleep time.Sleep tm second
		robotgo.Sleep(180)
	}
}
