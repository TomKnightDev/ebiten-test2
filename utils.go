package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type TileMap struct {
	Tileshigh int `json:"tileshigh"`
	Layers    []struct {
		Tiles []struct {
			X     int  `json:"x"`
			Rot   int  `json:"rot"`
			Y     int  `json:"y"`
			Index int  `json:"index"`
			FlipX bool `json:"flipX"`
			Tile  int  `json:"tile"`
		} `json:"tiles"`
		Name   string `json:"name"`
		Number int    `json:"number"`
	} `json:"layers"`
	Tileheight int `json:"tileheight"`
	Tileswide  int `json:"tileswide"`
	Tilewidth  int `json:"tilewidth"`
}

func HasTag(e *entity, tag Tag) bool {
	for _, t := range e.tags {
		if t == tag {
			return true
		}
	}

	return false
}

func GetEntsWithTag(game *Game, tag Tag) []*entity {
	ents := []*entity{}

	for _, e := range game.entities {
		if HasTag(e, tag) {
			ents = append(ents, e)
		}
	}

	return ents
}

func GetMag(vec2 f64.Vec2) float64 {
	return math.Sqrt(vec2[0]*vec2[0] + vec2[1]*vec2[1])
}

func GetCursorPos(game *Game) f64.Vec2 {
	x, y := ebiten.CursorPosition()
	camx := game.camera.Position[0]
	camy := game.camera.Position[1]

	return f64.Vec2{
		float64(x) + camx,
		float64(y) + camy,
	}
}
