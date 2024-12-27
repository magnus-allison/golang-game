package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type UI struct {
	playerHeartImage *ebiten.Image
}

func createUI() *UI {
	img, _, err := ebitenutil.NewImageFromFile("assets/hearts.png")
	if err != nil {
		log.Fatal(err)
	}
	return &UI{
		playerHeartImage: img,
	}
}

func (ui *UI) drawPlayerHearts(screen *ebiten.Image, amount int) {

	size := 32

	for i := 0; i < amount; i++ {
		rect := image.Rect(0, 0, 16, 16)
		croppedImage := ui.playerHeartImage.SubImage(rect).(*ebiten.Image)

		// Create draw options
		opts := &ebiten.DrawImageOptions{}
		scaleX := float64(size) / float64(croppedImage.Bounds().Dx())
		scaleY := float64(size) / float64(croppedImage.Bounds().Dy())
		opts.GeoM.Scale(scaleX, scaleY)
		opts.GeoM.Translate(
			float64((S_WIDTH - size*2) - (i*size) - 0*i),
			float64(size / 2),
		)

		screen.DrawImage(croppedImage, opts)
	}
}
