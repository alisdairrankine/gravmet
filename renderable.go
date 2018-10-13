package gravmet

import "github.com/veandco/go-sdl2/sdl"

type Renderable interface {
	Surface() *sdl.Surface
}
