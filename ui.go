package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func (ui *UI) drawGameOverScreen(screen *ebiten.Image) {
    // Draw a semi-transparent black background
    vector.DrawFilledRect(screen, 0, 0, float32(S_WIDTH), float32(S_HEIGHT), color.RGBA{0, 0, 0, 200}, true)

    // Draw the Game Over text in the middle of the screen
    message := "GAME OVER"

    // textWidth, textHeight := ebitenutil.DebugPrintAt(screen, message, S_WIDTH/2-70, S_HEIGHT/2-20) // Position the text
    ebitenutil.DebugPrintAt(screen, message, (S_WIDTH/2-len(message)*8/2), (S_HEIGHT/2-20))
}
