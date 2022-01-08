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

func (s *SceneManager) Draw(screen *ebiten.Image, game *Game) {
	for _, e := range game.entities {
		if !e.active {
			continue
		}
		for _, c := range e.components {
			c.Draw(screen, game)
		}
	}
}

func (s *SceneManager) GoTo(scene *entity) {
	if s.current != nil {
		s.current.active = false
	}

	s.current = scene
}
