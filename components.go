package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/math/f64"
)

type component interface {
	Update() error
	Draw(screen *ebiten.Image, game *Game)
}

type spriteRenderer struct {
	container *entity
	imageTile *imageTile
}

func newSpriteRenderer(container *entity, ip imagePos) *spriteRenderer {
	return &spriteRenderer{
		container: container,
		imageTile: &imageTile{
			image:    tilesImage.SubImage(ip.sheetPos).(*ebiten.Image),
			position: ip.worldPos,
		},
	}
}

func (s *spriteRenderer) Update() error {
	return nil
}

func (s *spriteRenderer) Draw(screen *ebiten.Image, game *Game) {
	m := ebiten.GeoM{}
	// m.Scale(1, 1)
	m.Translate(s.container.position[0], s.container.position[1])

	game.world.DrawImage(s.imageTile.image, &ebiten.DrawImageOptions{
		GeoM: m,
	})
}

type spritesRenderer struct {
	container  *entity
	imageTiles []imageTile
}

type imageTile struct {
	image    *ebiten.Image
	position f64.Vec2
}

type imagePos struct {
	sheetPos image.Rectangle
	worldPos f64.Vec2
}

func newSpritesRenderer(container *entity, images []imagePos) *spritesRenderer {
	imageTiles := []imageTile{}

	for _, ip := range images {
		imageTiles = append(imageTiles, imageTile{
			image:    tilesImage.SubImage(ip.sheetPos).(*ebiten.Image),
			position: ip.worldPos,
		})
	}

	return &spritesRenderer{
		container:  container,
		imageTiles: imageTiles,
	}
}

func (s *spritesRenderer) Update() error {
	return nil
}

func (s *spritesRenderer) Draw(screen *ebiten.Image, game *Game) {
	for _, imageTile := range s.imageTiles {
		m := ebiten.GeoM{}
		// m.Scale(1, 1)
		m.Translate(imageTile.position[0], imageTile.position[1])

		game.world.DrawImage(imageTile.image, &ebiten.DrawImageOptions{
			GeoM: m,
		})
	}
}

type uiRenderer struct {
	container *entity
}

func newUiRenderer(container *entity) *uiRenderer {
	return &uiRenderer{
		container: container,
	}
}

func (u *uiRenderer) Update() error {
	return nil
}

func (u *uiRenderer) Draw(screen *ebiten.Image, game *Game) {
	ebitenutil.DebugPrint(screen, u.container.name)
}

type input struct {
	container *entity
}

func newInput(container *entity) *input {
	return &input{
		container: container,
	}
}

func (i *input) Update() error {
	c := i.container
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		c.position[0] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		c.position[0] += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		c.position[1] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		c.position[1] += 1
	}

	// if ebiten.IsKeyPressed(ebiten.KeyR) {
	// 	g.camera.Rotation += 1
	// }

	// if ebiten.IsKeyPressed(ebiten.KeySpace) {
	// 	g.camera.Reset()
	// }

	return nil
}

func (i *input) Draw(screen *ebiten.Image, game *Game) {

}
