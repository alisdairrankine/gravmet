package gravmet

import "github.com/veandco/go-sdl2/sdl"

type TileMap struct {
	tileDimension uint
	width         uint
	height        uint
	terrain       []Tile
	entities      []Tile
}

func NewTileMap(w, h uint) *TileMap {

	terrain := make([]Tile, w*h)
	entities := make([]Tile, w*h)

	return &TileMap{
		tileDimension: 16,
		width:         w,
		height:        h,
		terrain:       terrain,
		entities:      entities,
	}
}

type Tile uint32

func (t *TileMap) AddTerrain(x, y uint, tile Tile) {
	t.terrain[t.coordToIndex(x, y)] = tile
}

func (t *TileMap) indexToCoord(i uint) (uint, uint) {
	y := i / t.width
	x := i % t.width
	return x, y
}

func (t *TileMap) coordToIndex(x, y uint) uint {
	return y*t.width + x
}

func (t *TileMap) Render(surface *sdl.Surface) {
	for i, tile := range t.terrain {
		x, y := t.indexToCoord(uint(i))
		rect := sdl.Rect{int32(x * t.tileDimension), int32(y * t.tileDimension), int32(t.tileDimension), int32(t.tileDimension)}
		surface.FillRect(&rect, uint32(tile))
	}
}
