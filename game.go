package main

import (
	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Game struct {
	worldImage   *ebiten.Image
	entities     []*entity
	sceneManager *SceneManager
	space        *resolv.Space
	camera       *Camera
	uiManager    *renderer.Manager
}

func (game *Game) Update() error {
	if game.sceneManager == nil {
		game.sceneManager = &SceneManager{}
		game.sceneManager.GoTo(newTitleScene(game))
	}

	if err := game.sceneManager.Update(game); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.uiManager.SetDisplaySize(screenWidth, screenHeight)
	return screenWidth, screenHeight
}
