package main

import (
	"github.com/alisdairrankine/gravmet"
)

func main() {

	renderer := gravmet.NewTileMap(8, 8)
	renderer.AddTerrain(0, 0, 0xffff0000)
	renderer.AddTerrain(1, 1, 0xffff0000)
	renderer.AddTerrain(2, 2, 0xffff0000)
	renderer.AddTerrain(0, 3, 0xffff0000)
	renderer.AddTerrain(1, 3, 0xffff0000)
	renderer.AddTerrain(2, 3, 0xffff0000)
	renderer.AddTerrain(3, 3, 0xffff0000)

	game := gravmet.NewGame(800, 600, renderer)
	game.Run()

}
