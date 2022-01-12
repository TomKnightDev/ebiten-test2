package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	transitionFrom = ebiten.NewImage(384, 384)
	transitionTo   = ebiten.NewImage(384, 384)
)

// const transitionMaxCount = 20

type SceneManager struct {
	current *entity
	next    *entity
}

func (s *SceneManager) Update(game *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.GoTo(newMainScene(game))
	}

	for _, e := range game.entities {
		if !e.active {
			continue
		}
		for _, c := range e.components {
			err := c.Update(game)
			if err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func (s *SceneManager) Draw(screen *ebiten.Image, game *Game) []imageTile {
	imageTiles := []imageTile{}

	for _, e := range game.entities {
		if !e.active {
			continue
		}
		for _, c := range e.components {
			for _, im := range c.Draw(screen, game) {

				imageTiles = append(imageTiles, im)

			}
		}
	}

	for _, imageTile := range imageTiles {
		m := ebiten.GeoM{}
		m.Translate(imageTile.position[0], imageTile.position[1])

		game.worldImage.DrawImage(imageTile.image, &ebiten.DrawImageOptions{
			GeoM: m,
		})
	}

	return []imageTile{}
}

func (s *SceneManager) GoTo(scene *entity) {
	if s.current != nil {
		s.current.active = false
	}

	s.current = scene
}
