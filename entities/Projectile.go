package entities

import (
	"golang-game/entities/enemies"
	"golang-game/utils"
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

func createProjectile(ox, oy, tx, ty int, speedModifier float32) *Projectile {
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
		var speed = 1.0 * speedModifier
		p.vx = dx / dist * speed // Velocity in the x direction
		p.vy = dy / dist * speed // Velocity in the y direction
	}

	return p
}

func (p *Projectile) update() {
	const speed = 7.0
	p.x += int(p.vx * speed)
	p.y += int(p.vy * speed)

		// for i := len(g.Projectiles) - 1; i >= 0; i-- {
		// 	projectile := g.Projectiles[i]
		// 	projectile.update()
		// 	projectile.checkCollision(g.EntityManager.Enemies)
}

func (p *Projectile) checkCollision(enemies []enemies.EnemyInterface) {
	for _, e := range enemies {
		eX, eY := e.GetPosition()
		eW, eH := e.GetSize()
		if utils.Collides(float32(p.x), float32(p.y), float32(p.size), float32(p.size), eX, eY, float32(eW), float32(eH)) {
			e.TakeDamage(p.vx * 0.4, p.vy * 0.4)
			// render offscreen for cleanup
			p.x = -100
			p.y = -100
		}
	}
}

func (p *Projectile) draw(screen *ebiten.Image) {

	// draw outline
	vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(p.size+2), color.RGBA{0, 0, 0, 255}, true)
	// draw inner circle
	vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(p.size), p.color, true)
}