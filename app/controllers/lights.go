package controllers

import (
	"github.com/revel/revel"
        "rpi_ws281x/golang/ws2811"
        "math"
)

type Lights struct {
	*revel.Controller
	Status string ` json:"status" `
}

func (c Lights) Off() revel.Result {
	off()
	json := Lights{Status: "ok"}
	return c.RenderJson(json)
}

func (c Lights) Colour(r int, g int, b int, a int) revel.Result {
	status := "unknown"
	if ((r >= 0 && r < 256) && (g >= 0 && g < 256) && (b >= 0 && b < 256) && (a >= 0 && a < 256)) {
		status = "ok"
		setColour(uint32(r), uint32(g), uint32(b), a)
	} else {
		status = "invalid colours"
	}
	json := Lights{Status: status}
	return c.RenderJson(json)
}

func setColour(red uint32, green uint32, blue uint32, brightness int) {
        redShift := red * uint32(math.Pow(2, 8))
        greenShift := green * uint32(math.Pow(2, 16))
        colour := redShift + greenShift + blue
        ws2811.Init(18, 60, brightness)
        for i := 0; i < 60; i++ {
                ws2811.SetLed(i, colour)
        }
        ws2811.Render()
}

func off() {
        ws2811.Init(18, 60, 100)
        ws2811.Clear()
        ws2811.Render()
}
