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
	// Create the projectile and initialize basic attributes
	p := &Projectile{
		size:  5,
		color: color.RGBA{255, 255, 255, 255},
		x:     ox,
		y:     oy,
		tx:    tx,
		ty:    ty,
	}

	// CHATGPT - Calculate direction vector
	dx := float32(tx - ox)
	dy := float32(ty - oy)
	dist := float32(math.Sqrt(float64(dx*dx + dy*dy)))
	if dist > 0 {
		// Normalize and apply a constant speed
		const speed = 1.0
		p.vx = dx / dist * speed // Velocity in the x direction
		p.vy = dy / dist * speed // Velocity in the y direction
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
			e.takeDamage(p.vx * 0.4, p.vy * 0.4)
			p.x = -100
			p.y = -100
		}
	}
}

func (p *Projectile) draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(p.size), p.color, true)
}
