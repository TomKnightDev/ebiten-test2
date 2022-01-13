package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
	"golang.org/x/image/math/f64"
)

type component interface {
	Update(game *Game) error
	Draw(screen *ebiten.Image, game *Game) []imageTile
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

func (s *spriteRenderer) Update(game *Game) error {
	return nil
}

func (s *spriteRenderer) Draw(screen *ebiten.Image, game *Game) []imageTile {
	c := s.container.getComponent(&Camera{})

	if c != nil {
		m := c.(*Camera).worldMatrix()
		m.Translate(s.container.position[0], s.container.position[1])

		screen.DrawImage(s.imageTile.image, &ebiten.DrawImageOptions{
			GeoM: m,
		})

		return []imageTile{}
	} else {
		return []imageTile{{
			image:    s.imageTile.image,
			position: s.container.position,
			zOrder:   s.imageTile.zOrder,
		}}

		// m := ebiten.GeoM{}
		// m.Translate(s.container.position[0], s.container.position[1])

		// game.worldImage.DrawImage(s.imageTile.image, &ebiten.DrawImageOptions{
		// 	GeoM: m,
		// })
	}
}

type spritesRenderer struct {
	container  *entity
	imageTiles []imageTile
}

type imageTile struct {
	image    *ebiten.Image
	position f64.Vec2
	zOrder   int
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

func (s *spritesRenderer) Update(game *Game) error {
	return nil
}

func (s *spritesRenderer) Draw(screen *ebiten.Image, game *Game) []imageTile {
	return s.imageTiles

	// for _, imageTile := range s.imageTiles {
	// 	m := ebiten.GeoM{}
	// 	// m.Scale(1, 1)
	// 	m.Translate(imageTile.position[0], imageTile.position[1])

	// 	game.worldImage.DrawImage(imageTile.image, &ebiten.DrawImageOptions{
	// 		GeoM: m,
	// 	})
	// }
}

type uiRenderer struct {
	container *entity
}

func newUiRenderer(container *entity) *uiRenderer {
	return &uiRenderer{
		container: container,
	}
}

func (u *uiRenderer) Update(game *Game) error {
	return nil
}

func (u *uiRenderer) Draw(screen *ebiten.Image, game *Game) []imageTile {
	ebitenutil.DebugPrint(screen, u.container.name)
	return []imageTile{}
}

type input struct {
	container *entity
}

func newInput(container *entity) *input {
	return &input{
		container: container,
	}
}

func (i *input) Update(game *Game) error {
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

func (i *input) Draw(screen *ebiten.Image, game *Game) []imageTile {
	return []imageTile{}
}

type boxCollider struct {
	container *entity
	collider  *resolv.Object
}

func newBoxCollider(container *entity, game *Game) *boxCollider {
	bc := &boxCollider{
		container: container,
		collider:  resolv.NewObject(container.position[0], container.position[1], 6, 6),
	}

	game.space.Add(bc.collider)
	return bc
}

func (b *boxCollider) Update(game *Game) error {
	x := b.container.position[0] - b.collider.X
	y := b.container.position[1] - b.collider.Y

	if collision := b.collider.Check(x, 0); collision != nil {
		b.container.position[0] = b.collider.X
	} else {
		b.collider.X = b.container.position[0]
		b.collider.Update()
	}

	if collision := b.collider.Check(0, y); collision != nil {
		b.container.position[1] = b.collider.Y
	} else {
		b.collider.Y = b.container.position[1]
		b.collider.Update()
	}

	return nil
}

func (b *boxCollider) Draw(screen *ebiten.Image, game *Game) []imageTile {
	return []imageTile{}
}
