package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct{
	x, y   float32
	vx, vy float32
	size int
	color color.RGBA
	isAttacking bool
	isColliding bool
	hearts int
}

func createPlayer() *Player {
	return &Player{
		x: 0,
		y: 0,
		size: 26,
		color: color.RGBA{255, 255, 255, 255},
		hearts: 3,
	}
}

func (p *Player) update() {
	const speed = 2.0        // Maximum speed
	const acceleration = 0.5 // How quickly the player accelerates
	const friction = 0.9     // Slow down when no key is pressed

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.vy -= acceleration
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.vy += acceleration
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.vx -= acceleration
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.vx += acceleration
	}

	p.vx *= friction
	p.vy *= friction

	if p.vx > speed {
		p.vx = speed
	} else if p.vx < -speed {
		p.vx = -speed
	}
	if p.vy > speed {
		p.vy = speed
	} else if p.vy < -speed {
		p.vy = -speed
	}

	p.x += p.vx
	p.y += p.vy

	// Prevent player from moving out of bounds
	if p.x < 0 {
		p.x = 0
		p.vx = 0
	} else if p.x > float32(S_WIDTH - p.size) {
		p.x = float32(S_WIDTH - p.size)
		p.vx = 0
	}
	if p.y < 0 {
		p.y = 0
		p.vy = 0
	} else if p.y > float32(S_HEIGHT-p.size) {
		p.y = float32(S_HEIGHT - p.size)
		p.vy = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.attack()
	}
}

func (p *Player) checkCollision(enemies []*Enemy) {
	for _, enemy := range enemies {
		if p.x < enemy.x+16 && p.x+16 > enemy.x && p.y < enemy.y+16 && p.y+16 > enemy.y {
			enemy.isColliding = true
		} else {
			enemy.isColliding = false
		}
	}
}


func (p *Player) attack() {
	p.isAttacking = true
	p.color = color.RGBA{255, 0, 0, 255}
	go func() {
		<-time.After(200 * time.Millisecond)
		p.isAttacking = false
		p.color = color.RGBA{255, 255, 255, 255}
	}()
}