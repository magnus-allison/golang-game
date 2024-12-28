package entities

import (
	"golang-game/config"
	"golang-game/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


type PowerUp struct {
	x, y float32
	w, h float32
	image *ebiten.Image
	rotation float64
}

func createPowerUp() *PowerUp {
	img, _, err := ebitenutil.NewImageFromFile("assets/powerup.png")
	if err != nil {
		log.Fatal(err)
	}

	return &PowerUp{
		x: utils.RandFloat32(0, float32(config.S_WIDTH)),
		y: utils.RandFloat32(0, float32(config.S_HEIGHT)),
		w: 20,
		h: 30,
		image: img,
	}
}

func (p *PowerUp) update() {
	// Update the powerup
	p.rotation += 0.02
}

func (p *PowerUp) draw(screen *ebiten.Image) {

opts := &ebiten.DrawImageOptions{}

	utils.DrawDebugBorder(screen, p.x, p.y, p.w, p.h)
	if (config.DISABLE_DRAW) { return }

	imgWidth, imgHeight := p.image.Bounds().Dx(), p.image.Bounds().Dy()
	scaleX := float64(p.w) / float64(imgWidth)
	scaleY := float64(p.h) / float64(imgHeight)

	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.image, opts)

}




