package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	w int32 = 800
	h int32 = 600
)

var d uint32 = 7

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("helloWorld", 100, 100, w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)

	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)

	}
	defer renderer.Destroy()

	bmp, err := sdl.LoadBMP("/home/runesmaker/go/src/gwg/basicks/sdl/Dragon.bmp")
	if err != nil {
		panic(err)

	}

	texture, err := renderer.CreateTextureFromSurface(bmp)
	if err != nil {
		panic(err)

	}
	defer texture.Destroy()

	renderer.Clear()
	fmt.Println("Очищаю визуализатор")
	renderer.Copy(texture, nil, nil)
	fmt.Println("Копирую текстуры")
	renderer.Present()
	fmt.Println("Визуализирую")
	fmt.Println("через", d, "секунд окно закроется.")
	sdl.Delay(d * 1000)

}
