package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"

	"github.com/veandco/go-sdl2/sdl"
)

func errorer(e error, n int) int {
	if e != nil {
		panic(e)
		return n
	}
	return 255
}

func run() int {
	const dellay uint32 = 10
	var title string = "Визуализация картинок"
	//var picturePath string = "img/tp.png"
	var picturePath string = "img/tb.jpeg"
	var w int32 = 740
	var h int32 = 555

	fmt.Println("Создаю окно")
	window, err := sdl.CreateWindow(title, 0, 0, w, h, sdl.WINDOW_SHOWN)
	errorer(err, 1)
	defer window.Destroy()

	fmt.Println("Создаю визуализатор")
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	errorer(err, 2)
	defer renderer.Destroy()

	fmt.Println("Загружаю изображение")
	image, err := img.Load(picturePath)
	errorer(err, 3)
	defer image.Free()

	fmt.Println("Создаю текстуру из картинки")
	texture, err := renderer.CreateTextureFromSurface(image)
	errorer(err, 4)
	defer texture.Destroy()

	fmt.Println("Создаю прямоугольники")
	src := sdl.Rect{0, 0, w, h}
	dst := sdl.Rect{0, 0, w, h}

	fmt.Println("Очищаю визуализатор")
	renderer.Clear()

	fmt.Println("Задаю цвет и поле")
	renderer.SetDrawColor(66, 60, 99, 0)
	renderer.FillRect(&sdl.Rect{0, 0, w, h})

	fmt.Println("Копирую текстуры")
	renderer.Copy(texture, &src, &dst)

	fmt.Println("Показываю")
	renderer.Present()

	sdl.PollEvent()
	fmt.Println("Окно закроется через", dellay, "секнд")
	sdl.Delay(dellay * 1000)
	fmt.Println("Выхожу")

	return 0

}

func main() {
	os.Exit(run())

}
