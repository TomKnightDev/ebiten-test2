package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type Camera struct {
	container  *entity
	ViewPort   f64.Vec2
	Position   f64.Vec2
	ZoomFactor int
	Rotation   int
}

func newCamera(container *entity) *Camera {
	c := &Camera{
		container: container,
		ViewPort:  f64.Vec2{screenWidth, screenHeight},
	}

	return c
}

func (c *Camera) Update(game *Game) error {
	c.Position = f64.Vec2{
		0: c.container.position[0] - c.viewportCenter()[0],
		1: c.container.position[1] - c.viewportCenter()[1],
	}

	// if ebiten.IsKeyPressed(ebiten.KeyQ) {
	// 	if c.ZoomFactor > -2400 {
	// 		c.ZoomFactor -= 1
	// 	}
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyE) {
	// 	if c.ZoomFactor < 2400 {
	// 		c.ZoomFactor += 1
	// 	}
	// }

	return nil
}

func (c *Camera) Draw(screen *ebiten.Image, game *Game) []imageTile {
	c.Render(screen, game)

	return []imageTile{}
}

func (c *Camera) viewportCenter() f64.Vec2 {
	return f64.Vec2{
		c.ViewPort[0] * 0.5,
		c.ViewPort[1] * 0.5,
	}
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.Position[0], -c.Position[1])
	// We want to scale and rotate around center of image / screen
	m.Translate(-c.viewportCenter()[0], -c.viewportCenter()[1])
	m.Scale(
		math.Pow(1.01, float64(c.ZoomFactor)),
		math.Pow(1.01, float64(c.ZoomFactor)),
	)
	m.Rotate(float64(c.Rotation) * 2 * math.Pi / 360)
	m.Translate(c.viewportCenter()[0], c.viewportCenter()[1])
	return m
}

func (c *Camera) Render(screen *ebiten.Image, game *Game) {
	screen.DrawImage(game.worldImage, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix(),
	})
}

// func (c *Camera) ScreenToWorld(posX, posY int) (float64, float64) {
// 	inverseMatrix := c.worldMatrix()
// 	if inverseMatrix.IsInvertible() {
// 		inverseMatrix.Invert()
// 		return inverseMatrix.Apply(float64(posX), float64(posY))
// 	} else {
// 		// When scaling it can happend that matrix is not invertable
// 		return math.NaN(), math.NaN()
// 	}
// }

func (c *Camera) Reset() {
	c.Position[0] = 0
	c.Position[1] = 0
	c.Rotation = 0
	c.ZoomFactor = 0
}
