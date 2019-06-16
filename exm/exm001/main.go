package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	window_size_w int32 = 600
	window_size_h int32 = 600

	rect_size_w int32 = window_size_w / 3
	rect_size_h int32 = window_size_h / 3
	rectPosX    int32 = (window_size_w / 2) - (rect_size_w / 2)
	rectPosY    int32 = (window_size_h / 2) - (rect_size_h / 2)
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("firsExample", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, window_size_w, window_size_h, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	sureface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	sureface.FillRect(nil, 0)

	rect := sdl.Rect{rectPosX, rectPosY, rect_size_w, rect_size_w}
	sureface.FillRect(&rect, 0x01796F)

	window.UpdateSurface()

	var running bool = true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Exit")
				running = false
				break
			}
		}
	}
}
