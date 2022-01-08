package main

import (
	"bytes"
	"image/png"
	_ "image/png"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 384
	screenHeight = 384
)

const (
	tileSize = 8
	tileXNum = 32
)

const (
	worldWidth  = 384
	worldHeight = 384
	worldSizeX  = worldWidth / tileSize
)

var (
	//go:embed resources/map001.png
	sprite_sheet []byte
	//go:embed resources/map001.json
	map001 []byte
	//go:embed resources/player1.json
	player1    []byte
	tilesImage *ebiten.Image
	tileMap    TileMap
)

func init() {
	img, err := png.Decode(bytes.NewReader(sprite_sheet))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

func main() {
	g := &Game{}
	g.world = ebiten.NewImage(worldWidth, worldHeight)

	ebiten.SetWindowSize(screenWidth*3, screenHeight*3)
	ebiten.SetWindowTitle("Ebiten-test2")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
