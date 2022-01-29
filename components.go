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
	if !s.container.active {
		return nil
	}

	if s.container.direction[0] != 0 || s.container.direction[1] != 0 {
		mag := GetMag(s.container.direction)

		s.container.position[0] += s.container.direction[0] / mag * s.container.velocity
		s.container.position[1] += s.container.direction[1] / mag * s.container.velocity
	}

	s.imageTile.position = s.container.position
	its := []imageTile{}
	its = append(its, *s.imageTile)
	return its
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
	update    UiUpdate
}

func newUiRenderer(container *entity, update UiUpdate) *uiRenderer {

	return &uiRenderer{
		container: container,
		update:    update,
	}
}

func (u *uiRenderer) Update(game *Game) error {
	return u.update(game)
}

func (u *uiRenderer) Draw(screen *ebiten.Image, game *Game) []imageTile {
	ebitenutil.DebugPrint(screen, u.container.name)
	game.uiManager.Draw(screen)
	return []imageTile{}
}

type input struct {
	container             *entity
	actionTurnTime        int
	currentActionTurnTime int
}

func newInput(container *entity, actionTurnTime int) *input {
	return &input{
		container:             container,
		actionTurnTime:        actionTurnTime,
		currentActionTurnTime: 0,
	}
}

func (i *input) Update(game *Game) error {
	c := i.container

	i.currentActionTurnTime++

	if i.currentActionTurnTime >= i.actionTurnTime && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		i.currentActionTurnTime = 0

		if shoot := c.getComponent(&shoots{}); shoot != nil {

			if HasTag(c, Ship) {
				// xdir := c.direction[0] - c.position[0]
				// ydir := c.direction[1] - c.position[1]
				b := newBullet(game, c.position, c.direction, Ship)
				b.velocity *= 2
			} else if HasTag(c, Player) {
				cpos := GetCursorPos(game)
				xdir := cpos[0] - c.position[0]
				ydir := cpos[1] - c.position[1]
				newBullet(game, c.position, f64.Vec2{xdir, ydir}, Player)
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		if HasTag(c, Player) {
			ship := GetEntsWithTag(game, Ship)

			c.active = false
			ship[0].active = true
			return nil
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		if HasTag(c, Ship) {
			player := GetEntsWithTag(game, Player)

			c.active = false
			player[0].active = true
			return nil
		}
	}

	x := 0.0
	y := 0.0

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		x += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		y += 1
	}

	if c.position[0]+x >= screenWidth || c.position[0]+x <= 0 {
		x = 0
	}

	if c.position[1]+y >= screenHeight || c.position[1]+y <= 0 {
		y = 0
	}

	if HasTag(c, Ship) {
		if x == 1 && y == 0 {
			c.rotation = 0
		} else if x == 1 && y == 1 {
			c.rotation = 45
		} else if x == 0 && y == 1 {
			c.rotation = 90
		} else if x == -1 && y == 1 {
			c.rotation = 135
		} else if x == -1 && y == 0 {
			c.rotation = 180
		} else if x == -1 && y == -1 {
			c.rotation = 225
		} else if x == 0 && y == -1 {
			c.rotation = 270
		} else if x == 1 && y == -1 {
			c.rotation = 315
		}

		if x != 0 || y != 0 {
			c.direction = f64.Vec2{x, y}
		}

		x *= 2
		y *= 2
	}

	c.position[0] += float64(x)
	c.position[1] += float64(y)

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

func newBoxCollider(container *entity, game *Game, originTag []string, size float64) *boxCollider {
	bc := &boxCollider{
		container: container,
		collider:  resolv.NewObject(container.position[0], container.position[1], size, size, originTag...),
	}

	game.space.Add(bc.collider)
	return bc
}

func (b *boxCollider) Update(game *Game) error {
	x := b.container.position[0] - b.collider.X
	y := b.container.position[1] - b.collider.Y

	if collision := b.collider.Check(x, 0); collision != nil {
		if !collision.HasTags(b.collider.Tags()...) {

			if HasTag(b.container, Bullet) {
				b.container.active = false
				return nil
			}

			if HasTag(b.container, Enemy) {
				b.container.active = false
				return nil
			}

			b.container.position[0] = b.collider.X
		}
	} else {
		b.collider.X = b.container.position[0]
		b.collider.Update()
	}

	if collision := b.collider.Check(0, y); collision != nil {
		if !collision.HasTags(b.collider.Tags()...) {

			if HasTag(b.container, Bullet) {
				b.container.active = false
				return nil
			}

			if HasTag(b.container, Enemy) {
				b.container.active = false
				return nil
			}

			b.container.position[1] = b.collider.Y
		}
	} else {
		b.collider.Y = b.container.position[1]
		b.collider.Update()
	}

	return nil
}

func (b *boxCollider) Draw(screen *ebiten.Image, game *Game) []imageTile {
	return []imageTile{}
}

type shoots struct {
	container *entity
}

func newShoots(container *entity) *shoots {
	return &shoots{
		container: container,
	}
}

func (s *shoots) Update(game *Game) error {
	return nil
}

func (s *shoots) Draw(screen *ebiten.Image, game *Game) []imageTile {
	return []imageTile{}
}

type mouseInput struct {
	container *entity
}

func newMouseInput(container *entity) *mouseInput {
	return &mouseInput{
		container: container,
	}
}

func (m *mouseInput) Update(game *Game) error {
	cpos := GetCursorPos(game)
	m.container.position[0] = cpos[0]
	m.container.position[1] = cpos[1]

	return nil
}

func (m *mouseInput) Draw(screen *ebiten.Image, game *Game) []imageTile {
	return []imageTile{}
}
