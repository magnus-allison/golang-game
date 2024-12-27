package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Projectile struct {
	x, y    int
	tx, ty  int
	vx, vy  float32
	size    int
	color   color.RGBA
}

func createProjectile(ox, oy, tx, ty int) *Projectile {
	// Create the projectile and calculate direction
	p := &Projectile{
		size: 5,
		color: color.RGBA{255, 255, 255, 255},
		x: ox,
		y: oy,
		tx: tx,
		ty: ty,
	}

	dx := float32(tx - ox)
	dy := float32(ty - oy)

	// GhatGPT - Normalize the direction vector
	dist := float32(math.Sqrt(float64(dx*dx + dy*dy)))
	if dist > 0 {
		p.vx = dx / dist
		p.vy = dy / dist
	}

	return p
}

func (p *Projectile) update() {
	const speed = 7.0
	p.x += int(p.vx * speed)
	p.y += int(p.vy * speed)
}

func (p *Projectile) checkCollision(enemies []*Enemy) {
	for _, e := range enemies {
		if float32(p.x) > e.x && float32(p.x) < e.x+float32(e.size) && float32(p.y) > e.y && float32(p.y) < e.y+float32(e.size) {
			e.hp--
			p.x = -100
			p.y = -100
		}
	}
}


func (p *Projectile) draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(p.size), p.color, true)
}
