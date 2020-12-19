package main

import (
	"fmt"
	"time"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

func drawRow(y int64) {
	for x := 0; x < 550; x++ {
		rand := nextRandFloat()
		if rand < 0.5 {
			cv.SetFillStyle("#fff")
			cv.FillRect(float64(x), float64(y), 1, 1)
		}
	}
}

func drawRand() {
	fmt.Println("drawRand")
	var y int64 = 0
	wnd.MainLoop(func() {
		drawRow(y)
		y++
		if y > 550 {
			time.Sleep(10 * time.Second)
			wnd.Close()
		}
	})
}

func drawFractal() {
	cv.Translate(600, 1500)
	cv.Rotate(4)

	wnd.MainLoop(func() {
		iterate()
	})
}

var cv *canvas.Canvas
var wnd *sdlcanvas.Window

func main() {
	fmt.Println("Hello World")
	wnd1, cv1, err := sdlcanvas.CreateWindow(1200, 1200, "Hello")
	cv = cv1
	wnd = wnd1
	if err != nil {
		panic(err)
	}
	defer wnd.Destroy()
	drawFractal()
	// drawRand() // Uncomment to see the distribution of random function
}
