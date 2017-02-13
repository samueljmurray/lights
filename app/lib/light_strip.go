package main

import (
	"rpi_ws281x/golang/ws2811"
	"math"
)

func SetColour(red uint32, green uint32, blue uint32) {
	redShift := red * uint32(math.Pow(2, 8))
	greenShift := green * uint32(math.Pow(2, 16))
	colour := redShift + greenShift + blue
	ws2811.Init(18, 60, 100)
	for i := 0; i < 60; i++ {
                ws2811.SetLed(i, colour)
        }
	ws2811.Render()
}

func Off() {
	ws2811.Init(18, 60, 100)
	ws2811.Clear()
	ws2811.Render()
}
