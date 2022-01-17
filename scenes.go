package main

import (
	"encoding/json"
	"image"

	"github.com/solarlune/resolv"
	"github.com/yohamta/furex"
	"golang.org/x/image/math/f64"
)

func newTitleScene(game *Game) *entity {
	titleScene := &entity{}
	titleScene.name = "Title Scene"
	titleScene.position = f64.Vec2{
		0: 0,
		1: 0,
	}

	titleScene.active = true
	titleScene.tags = append(titleScene.tags, Scene)

	ur := newUiRenderer(titleScene)

	ur.rootFlex.Direction = furex.Row
	ur.rootFlex.Justify = furex.JustifyCenter
	ur.rootFlex.AlignItems = furex.AlignItemCenter
	ur.rootFlex.AlignContent = furex.AlignContentCenter
	ur.rootFlex.Wrap = furex.Wrap

	ur.rootFlex.AddChild(NewButton(50, 20, "Begin"))

	titleScene.addComponent(ur)

	game.entities = append(game.entities, titleScene)
	return titleScene
}

func newMainScene(game *Game) *entity {
	mainScene := &entity{}

	mainScene.position = f64.Vec2{
		0: 0,
		1: 0,
	}

	mainScene.active = true
	mainScene.tags = append(mainScene.tags, Scene)

	ur := newUiRenderer(mainScene)
	mainScene.addComponent(ur)

	// Tiles
	if err := json.Unmarshal(map001, &tileMap); err != nil {
		panic(err)
	}

	ips := []imagePos{}

	for l := len(tileMap.Layers) - 1; l >= 0; l-- {
		for _, t := range tileMap.Layers[l].Tiles {
			if tileMap.Layers[l].Name == "Obstacles" && t.Tile >= 0 {
				o := resolv.NewObject(float64(t.X*tileSize), float64(t.Y*tileSize), 8, 8, "wall")
				game.space.Add(o)
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
	game.entities = append(game.entities, mainScene)

	// TODO: Move to better place
	newPlayer(game, f64.Vec2{
		0: 18 * tileSize,
		1: 15 * tileSize,
	})

	// newEnemy(game, f64.Vec2{
	// 	0: 20 * tileSize,
	// 	1: 18 * tileSize,
	// })

	// newEnemy(game, f64.Vec2{
	// 	0: 21 * tileSize,
	// 	1: 18 * tileSize,
	// })

	// newEnemy(game, f64.Vec2{
	// 	0: 20 * tileSize,
	// 	1: 19 * tileSize,
	// })

	newShip(game, f64.Vec2{
		0: 2 * tileSize,
		1: 2 * tileSize,
	})

	return mainScene
}
