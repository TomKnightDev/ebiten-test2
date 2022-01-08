package main

import (
	"encoding/json"
	"image"

	"golang.org/x/image/math/f64"
)

func newTitleScene() *entity {
	titleScene := &entity{}
	titleScene.name = "Title Scene"
	titleScene.position = f64.Vec2{
		0: 0,
		1: 0,
	}

	titleScene.active = true

	ur := newUiRenderer(titleScene)
	titleScene.addComponent(ur)

	Entities = append(Entities, titleScene)
	return titleScene
}

func newMainScene() *entity {
	mainScene := &entity{}

	mainScene.position = f64.Vec2{
		0: 0,
		1: 0,
	}

	mainScene.active = true

	ur := newUiRenderer(mainScene)
	mainScene.addComponent(ur)

	// Tiles
	if err := json.Unmarshal(map001, &tileMap); err != nil {
		panic(err)
	}

	ips := []imagePos{}

	for l := len(tileMap.Layers) - 1; l >= 0; l-- {
		for _, t := range tileMap.Layers[l].Tiles {
			if tileMap.Layers[l].Name == "Obstacles" {
				continue
			}

			sx := (t.Tile % tileXNum) * tileSize
			sy := (t.Tile / tileXNum) * tileSize

			ips = append(ips, imagePos{
				sheetPos: image.Rect(sx, sy, sx+tileSize, sy+tileSize),
				worldPos: f64.Vec2{
					0: float64(t.X * tileSize),
					1: float64(t.Y * tileSize),
				},
			})
		}
	}

	sr := newSpritesRenderer(mainScene, ips)
	mainScene.addComponent(sr)
	Entities = append(Entities, mainScene)

	// TODO: Move to better place
	newPlayer()

	return mainScene
}
