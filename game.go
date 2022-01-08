package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world        *ebiten.Image
	entities     []entity
	sceneManager *SceneManager
}

func init() {
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoTo(newTitleScene())
	}

	// g.input.Update()
	if err := g.sceneManager.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.sceneManager.Draw(screen, g)

	// for _, l := range g.layers {
	// 	for i, t := range l {
	// 		op := &ebiten.DrawImageOptions{}
	// 		op.GeoM.Translate(float64((i%worldSizeX)*tileSize), float64((i/worldSizeX)*tileSize))

	// 		sx := (t % tileXNum) * tileSize
	// 		sy := (t / tileXNum) * tileSize
	// 		g.world.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
	// 	}
	// }
	// g.camera.Render(g.world, screen)

	// worldX, worldY := g.camera.ScreenToWorld(ebiten.CursorPosition())
	// ebitenutil.DebugPrint(
	// 	screen,
	// 	fmt.Sprintf("TPS: %0.2f\nMove (WASD/Arrows)\nZoom (QE)\nRotate (R)\nReset (Space)", ebiten.CurrentTPS()),
	// )

	// ebitenutil.DebugPrintAt(
	// 	screen,
	// 	fmt.Sprintf("%s\nCursor World Pos: %.2f,%.2f",
	// 		g.camera.String(),
	// 		worldX, worldY),
	// 	0, screenHeight-32,
	// )
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
