package gravmet

import "github.com/veandco/go-sdl2/sdl"

type Renderer interface {
	Render(*sdl.Surface)
}
