package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	tit      string = "Вывод текста"
	w        int32  = 800
	h        int32  = 600
	fontSize int    = 42
	fontPath string = "Albionic.ttf"
	text     string = "it is a goodday ot die"
	delay    uint32 = 7
)

func errorer(e error, n int) int {
	if e != nil {
		panic(e)
		return n
	}
	return 255
}
func run() int {
	fmt.Println("Инициализирую Видео")
	sdl.Init(sdl.INIT_VIDEO)

	fmt.Println("Инициализирую текст")
	if err := ttf.Init(); err != nil {
		panic(err)
		return 1
	}

	fmt.Println("Создаю окно")
	window, err := sdl.CreateWindow(tit, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)
	errorer(err, 2)
	defer window.Destroy()

	fmt.Println("Задаю шрифт")
	font, err := ttf.OpenFont(fontPath, fontSize)
	errorer(err, 3)
	defer font.Close()

	fmt.Println("Визуализирую текст")
	solid, err := font.RenderUTF8Solid(text, sdl.Color{0, 126, 52, 255})
	errorer(err, 4)
	defer solid.Free()

	fmt.Println("Передаю поверхность окну")
	surface, err := window.GetSurface()
	errorer(err, 5)

	if err := solid.Blit(nil, surface, nil); err != nil {
		panic(err)
		return 6
	}

	fmt.Println("Обновляю поверхность окна")
	window.UpdateSurface()
	sdl.PollEvent()
	fmt.Println("Окно закроется через", delay, "секунд.")
	sdl.Delay(delay * 1000)

	return 0

}

func main() {
	os.Exit(run())

}
