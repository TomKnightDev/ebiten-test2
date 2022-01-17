package main

import (
	"math"

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
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && s.current.name == "Title Scene" {
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
	imageTiles := []struct {
		e    entity
		tile imageTile
	}{}

	for _, e := range game.entities {
		if !e.active {
			continue
		}
		for _, c := range e.components {
			for _, im := range c.Draw(screen, game) {

				imageTiles = append(imageTiles, struct {
					e    entity
					tile imageTile
				}{e: *e, tile: im})

			}
		}
	}

	for _, imageTile := range imageTiles {
		m := ebiten.GeoM{}

		m.Translate(-(tileSize / 2), -(tileSize / 2))
		m.Rotate(imageTile.e.rotation * 2 * math.Pi / 360)
		m.Translate(imageTile.tile.position[0], imageTile.tile.position[1])

		game.worldImage.DrawImage(imageTile.tile.image, &ebiten.DrawImageOptions{
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
