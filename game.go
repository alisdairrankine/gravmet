package gravmet

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Game struct {
	windowWidth  int32
	windowHeight int32
	window       *sdl.Window
	r            Renderer
}

func NewGame(windowWidth, windowHeight int32, r Renderer) *Game {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Gravitation Metal",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		windowWidth,
		windowHeight,
		sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	return &Game{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		window:       window,
		r:            r,
	}

}

func (g *Game) Run() {

	font, err := ttf.OpenFont("Roboto-Regular.ttf", 10)
	if err != nil {
		panic(err)
	}

	defer g.Cleanup()
	surface, err := g.window.GetSurface()
	if err != nil {
		panic(err)
	}
	t := time.Now()
	var dt time.Duration = 0
	white := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	showFPS := false
	for {
		t = time.Now()
		surface.FillRect(nil, white.Uint32())
		g.r.Render(surface)

		if showFPS {
			fps, err := font.RenderUTF8Solid(fmt.Sprintf("%.0f", 1/dt.Seconds()), white)
			if err == nil && fps != nil {
				srect := &sdl.Rect{X: 0, Y: 0, W: fps.W, H: fps.H}
				drect := &sdl.Rect{X: g.windowWidth - fps.W - 10, Y: 10, W: fps.W, H: fps.H}
				fps.Blit(srect, surface, drect)
			}
		}

		g.window.UpdateSurface()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		dt = time.Now().Sub(t)
		if dt.Nanoseconds() < 16666666 {
			time.Sleep(time.Duration((16666666 - dt.Nanoseconds())) * time.Nanosecond)
		}

	}
}

func (g *Game) Cleanup() {
	g.window.Destroy()
	sdl.Quit()
}
