package main

import (
	"encoding/json"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/solarlune/resolv"
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

	ur := newUiRenderer(titleScene, titleUpdate)
	titleScene.addComponent(ur)

	game.entities = append(game.entities, titleScene)
	return titleScene
}

type UiUpdate func(*Game) error

func titleUpdate(game *Game) error {
	game.uiManager.Update(1.0 / 60.0)
	game.uiManager.BeginFrame()
	{
		flags := imgui.WindowFlagsNone
		flags |= imgui.WindowFlagsNoTitleBar
		flags |= imgui.WindowFlagsNoResize

		imgui.SetNextWindowPos(imgui.Vec2{(screenWidth - imgui.WindowSize().X) / 2, (screenHeight - imgui.WindowSize().Y) / 2})
		imgui.BeginV("Main Menu", nil, flags)

		imgui.SetNextWindowPos(imgui.Vec2{imgui.WindowSize().X / 2, imgui.WindowSize().Y / 2})
		if imgui.Button("Start Game") {
			game.sceneManager.GoTo(newMainScene(game))
		}

		imgui.End()
	}
	game.uiManager.EndFrame()
	return nil
}

func newMainScene(game *Game) *entity {
	mainScene := &entity{}

	mainScene.position = f64.Vec2{
		0: 0,
		1: 0,
	}

	mainScene.active = true
	mainScene.tags = append(mainScene.tags, Scene)

	ur := newUiRenderer(mainScene, mainUpdate)
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
		1: 18 * tileSize,
	})

	newMouseCursor(game)

	newEnemy(game, f64.Vec2{
		0: 20 * tileSize,
		1: 18 * tileSize,
	})

	newEnemy(game, f64.Vec2{
		0: 21 * tileSize,
		1: 18 * tileSize,
	})

	newEnemy(game, f64.Vec2{
		0: 20 * tileSize,
		1: 19 * tileSize,
	})

	newShip(game, f64.Vec2{
		0: 2 * tileSize,
		1: 2 * tileSize,
	})

	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	return mainScene
}

func mainUpdate(game *Game) error {
	game.uiManager.Update(1.0 / 60.0)
	game.uiManager.BeginFrame()
	{
		if imgui.BeginMainMenuBar() {
			imgui.Text("Health: 100")
			if imgui.Button("Quit") {
				os.Exit(0)
			}
			imgui.EndMainMenuBar()
		}
	}
	game.uiManager.EndFrame()
	return nil
}
