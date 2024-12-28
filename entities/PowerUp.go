package entities

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)


type PowerUp struct {
	x, y float32
	size float32
	color color.RGBA
	rotation float64
}

func createPowerUp(x, y float32) *PowerUp {
	return &PowerUp{
		x: x,
		y: y,
		size: 16,
		color: color.RGBA{255, 0, 0, 255},
	}
}

func (p *PowerUp) update() {
	// Update the powerup
	p.rotation += 0.1
}

func (p *PowerUp) draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(p.x), float32(p.y), p.size, p.size, p.color, true)
}



