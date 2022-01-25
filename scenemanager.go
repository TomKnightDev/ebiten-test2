package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SceneManager struct {
	current *entity
}

func (s *SceneManager) Update(game *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && s.current.name == "Title Scene" {
		s.GoTo(newMainScene(game))
	}

	for i, e := range game.entities {
		if e.lifetime > 0 {
			e.currentLifetime++
			if e.currentLifetime >= e.lifetime {
				bc := e.getComponent(&boxCollider{}).(*boxCollider)
				game.space.Remove(bc.collider)
				RemoveEntity(game, i)
			}
		}
		if !e.active {
			if HasTag(e, Bullet) || HasTag(e, Enemy) {
				bc := e.getComponent(&boxCollider{}).(*boxCollider)
				game.space.Remove(bc.collider)
				RemoveEntity(game, i)
			}
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

	ebitenutil.DebugPrint(screen, fmt.Sprint(len(game.entities)))

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

func RemoveEntity(game *Game, i int) {
	if i >= len(game.entities) {
		return
	}

	game.entities[i] = game.entities[len(game.entities)-1]
	game.entities = game.entities[:len(game.entities)-1]

	// return append(scenes[:s], scenes[s+1:]...)
}
