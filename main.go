package main

import (
	"bytes"
	"image/png"
	_ "image/png"
	"log"

	_ "embed"

	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

const (
	screenWidth  = 768
	screenHeight = 768
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
		space:     resolv.NewSpace(384, 384, 4, 4),
		uiManager: renderer.New(nil),
	}

	// Not sure why the neg half tile size is needed here, but it is
	g.worldImage = ebiten.NewImage(worldWidth-(tileSize/2), worldHeight-(tileSize/2))

	ebiten.SetWindowResizable(true)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten-test2")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
