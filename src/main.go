package main

import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth int = 800
const winHeight int = 600

type color struct {
	r, g, b byte
}

func setPixel(x int, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index > 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Testing SDL", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	pixels := make([]byte, winHeight*winWidth*4)
	paddle1, ball, paddle2 := initializePong()
	keyState := sdl.GetKeyboardState()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		startPong(&paddle1, &ball, &paddle2, pixels, keyState)

		tex.Update(nil, unsafe.Pointer(&pixels[0]), winWidth*4)

		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(16)
	}

	// holdThatWindowForMe()
}

func holdThatWindowForMe() {
	for i := 0; i < 2000; i++ {
		sdl.PumpEvents()
		sdl.Delay(1)
	}
}
