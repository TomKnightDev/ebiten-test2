package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type basicAI struct {
	container *entity
	target    *entity
}

func newBasicAI(container *entity) *basicAI {
	bai := &basicAI{
		container: container,
	}

	return bai
}

func (bai *basicAI) Update(game *Game) error {
	c := bai.container
	t := bai.target

	if t == nil {
		for _, ent := range game.entities {
			if len(ent.tags) > 0 && ent.tags[0] == Player {
				xDist := math.Abs(ent.position[0] - c.position[0])
				yDist := math.Abs(ent.position[1] - c.position[1])

				if xDist < tileSize*4 || yDist < tileSize*4 {
					t = ent
					break
				}
			}
		}
	}

	if t != nil {
		xDist := math.Abs(t.position[0] - c.position[0])
		yDist := math.Abs(t.position[1] - c.position[1])

		if xDist > tileSize || yDist > tileSize {
			if t.position[0] > c.position[0] {
				c.position[0] += 0.5
			}
			if t.position[0] < c.position[0] {
				c.position[0] -= 0.5
			}
			if t.position[1] > c.position[1] {
				c.position[1] += 0.5
			}
			if t.position[1] < c.position[1] {
				c.position[1] -= 0.5
			}
		}
	}

	return nil
}

func (bai *basicAI) Draw(screen *ebiten.Image, game *Game) []imageTile {
	return []imageTile{}
}
