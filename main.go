package main

import (
	"bytes"
	"image/png"
	_ "image/png"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
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
	player1 []byte
	//go:embed resources/enemy1.json
	enemy1 []byte
	//go:embed resources/ship1.json
	ship1 []byte
	//go:embed resources/bullet1.json
	bullet1 []byte
	//go:embed resources/cursor1.json
	cursor1    []byte
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
	g := &Game{
		space: resolv.NewSpace(384, 384, 4, 4),
	}
	g.worldImage = ebiten.NewImage(worldWidth, worldHeight)

	ebiten.SetWindowResizable(true)

	ebiten.SetWindowSize(screenWidth*3, screenHeight*3)
	ebiten.SetWindowTitle("Ebiten-test2")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
