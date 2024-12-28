package levels

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"golang-game/config"
	_ "image/jpeg"
)

type Level struct{
	name string
	image *ebiten.Image
}

func createLevelOne() *Level {
	img, _, err := ebitenutil.NewImageFromFile("assets/floor.jpg")
	if err != nil {
		log.Fatal(err)
	}

	return &Level{
		name: "Level One",
		image: img,
	}
}

func (p *Level) DrawFloor(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}

	// scale to fit the screen
	scaleX := float64(config.S_WIDTH) / float64(p.image.Bounds().Dx())
	scaleY := float64(config.S_HEIGHT) / float64(p.image.Bounds().Dy())
	opts.GeoM.Scale(scaleX, scaleY)
	// apply mask
	opts.ColorScale.Scale(0.32, 0.2, 0.2, 0.7)

	screen.DrawImage(p.image, opts)
}