package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct{
	x, y float32
	color color.RGBA
	isAttacking bool
	isColliding bool
}

func createPlayer() *Player {
	return &Player{
		x: 0,
		y: 0,
		color: color.RGBA{255, 255, 255, 255},
	}
}

func (p *Player) update() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.x += 1
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