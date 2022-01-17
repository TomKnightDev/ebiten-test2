package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Button struct {
	size       image.Point
	isPressing bool
	margin     []int
	text       string
}

func NewButton(w, h int, text string) *Button {
	b := new(Button)
	b.size = image.Pt(w, h)
	b.text = text
	b.margin = []int{0, 0, 0, 0}
	return b
}

func (b *Button) Size() (int, int) {
	return b.size.X, b.size.Y
}

func (b *Button) Margin() []int {
	return b.margin
}

func (b *Button) SetMargin(m []int) {
	b.margin = m
}

func (b *Button) HandlePress(x, y int, t ebiten.TouchID) {
	b.isPressing = true
}

func (b *Button) HandleRelease(x, y int, isCancel bool) {
	b.isPressing = false
	if isCancel {
		println("The click is cancelled!")
	} else {
		println("clicked!")
	}
}

func (b *Button) Draw(screen *ebiten.Image, frame image.Rectangle) {
	if b.isPressing {
		FillRect(screen, frame, color.RGBA{0xaa, 0, 0, 0xff})
	} else {
		FillRect(screen, frame, color.RGBA{0, 0xaa, 0, 0xff})
	}
	DrawRect(screen, frame, color.RGBA{0xff, 0xff, 0xff, 0xff}, 2)
	ebitenutil.DebugPrintAt(screen, b.text,
		frame.Min.X+((frame.Dx()-36)/2), frame.Min.Y+b.size.Y/2-8)
}

var (
	imgOfAPixel *ebiten.Image
)

func createRectImg() *ebiten.Image {
	if imgOfAPixel != nil {
		return imgOfAPixel
	}
	imgOfAPixel := ebiten.NewImage(1, 1)
	return imgOfAPixel
}

func FillRect(target *ebiten.Image, r image.Rectangle, clr color.Color) {
	img := createRectImg()
	img.Fill(clr)

	op := &ebiten.DrawImageOptions{}

	size := r.Size()
	op.GeoM.Translate(float64(r.Min.X)*(1/float64(size.X)),
		float64(r.Min.Y)*(1/float64(size.Y)))
	op.GeoM.Scale(float64(size.X), float64(size.Y))

	target.DrawImage(img, op)
}

func DrawRect(target *ebiten.Image, r image.Rectangle, clr color.Color, width int) {
	FillRect(target, image.Rect(r.Min.X, r.Min.Y, r.Min.X+width, r.Max.Y), clr)
	FillRect(target, image.Rect(r.Max.X-width, r.Min.Y, r.Max.X, r.Max.Y), clr)
	FillRect(target, image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Min.Y+width), clr)
	FillRect(target, image.Rect(r.Min.X, r.Max.Y-width, r.Max.X, r.Max.Y), clr)
}
