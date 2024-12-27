package main

import (
	"image/color"
	"math"
)

type Enemy struct {
    x, y     float32
    vx, vy   float32
    color    color.RGBA
    isAttacking bool
    isColliding bool
}


func createEnemy() *Enemy {
	return &Enemy{
		x: randFloat32(0, float32(S_WIDTH)),
		y: randFloat32(0, float32(S_HEIGHT)),
		color: color.RGBA{0, 255, 55, 255},
	}
}

func (e *Enemy) update(playerX, playerY float32) {
	// Update color based on collision
	if e.isColliding {
		e.color = color.RGBA{255, 0, 0, 255} // Red when colliding
	} else {
		e.color = color.RGBA{0, 255, 55, 255} // Green when not colliding
	}

	// Maintain direction for a short duration
	if randFloat32(0, 1) < 0.05 { // 5% chance to change direction
		e.vx = randFloat32(-1, 1)
		e.vy = randFloat32(-1, 1)
	}

	// Add some goal-seeking behavior toward the player
	if randFloat32(0, 1) < 0.1 { // Occasionally adjust direction toward the player
		dx := playerX - e.x
		dy := playerY - e.y
		mag := float32(math.Sqrt(float64(dx*dx + dy*dy)))
		if mag > 0 {
			e.vx += 0.1 * dx / mag
			e.vy += 0.1 * dy / mag
		}
	}

	// Update position
	e.x += e.vx
	e.y += e.vy

	// Boundary check to prevent moving out of screen
	if e.x < 0 {
		e.x = 0
		e.vx = -e.vx
	} else if e.x > float32(S_WIDTH-16) {
		e.x = float32(S_WIDTH - 16)
		e.vx = -e.vx
	}

	if e.y < 0 {
		e.y = 0
		e.vy = -e.vy
	} else if e.y > float32(S_HEIGHT-16) {
		e.y = float32(S_HEIGHT - 16)
		e.vy = -e.vy
	}
}
