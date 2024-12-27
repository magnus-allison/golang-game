package main

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct{
	x, y float32
	vx, vy float32
	size int
	isAttacking bool
	hearts int
	image *ebiten.Image
	frameIdx, animCounter int
	invincible bool
	tintDuration int
	canShoot bool
}

func createPlayer() *Player {
	img, _, err := ebitenutil.NewImageFromFile("assets/player2.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		size: 42,
		image: img,
		canShoot: true,
	}
}

func (p *Player) draw(screen *ebiten.Image) {
	frameWidth := 32
	frameHeight := 32
	// cropping rect for the current frame
	frameX := (p.frameIdx % (p.image.Bounds().Dx() / frameWidth)) * frameWidth
	frameY := (p.frameIdx / (p.image.Bounds().Dx() / frameWidth)) * frameHeight
	cropRect := image.Rect(frameX, frameY, frameX+frameWidth, frameY+frameHeight)

	frame := p.image.SubImage(cropRect).(*ebiten.Image)
	opts := &ebiten.DrawImageOptions{}

	if p.tintDuration > 0 {
		opts.ColorScale.Scale(0.99, 0.222, 0.114, 1)
		p.tintDuration--
	}

	scaleX := float64(p.size) / float64(frameWidth)
	scaleY := float64(p.size) / float64(frameHeight)
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(frame, opts)
}


func (p *Player) update() {
	const speed = 1.5 // max speed
	const acceleration = 0.5
	const friction = 0.9

	// S = 0 - 11
	// A = 12 - 23
	// D = 24 - 35
	// W = 36 - 47

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.animationFrame(36, 47)
		p.vy -= acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.animationFrame(0, 11)
		p.vy += acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.animationFrame(12, 23)
		p.vx -= acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.animationFrame(24, 35)
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

	// bound check
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

func (p *Player) animationFrame(start int, end int) {
	const frameRate = 8
	if (p.frameIdx < start || p.frameIdx > end) {
		p.frameIdx = start
	}
	p.animCounter++
	if p.animCounter >= frameRate {
		p.frameIdx++
		if p.frameIdx > end { p.frameIdx = start }
		p.animCounter = 0
	}
}

func (p *Player) checkCollision(enemies []*Enemy) {
	for _, enemy := range enemies {
		playerRight := p.x + float32(p.size)
		playerBottom := p.y + float32(p.size)
		enemyRight := enemy.x + float32(enemy.size)
		enemyBottom := enemy.y + float32(enemy.size)

		if p.x < enemyRight && playerRight > enemy.x && p.y < enemyBottom && playerBottom > enemy.y {
			enemy.isColliding = true
			p.takeDamage()
		} else {
			enemy.isColliding = false
		}
	}
}

func (p *Player) takeDamage() {
	if (p.invincible) {
		return
	}
	p.tintDuration = 30
	p.hearts--
	p.invincible = true

	go func() {
		time.Sleep(1 * time.Second)
		p.invincible = false
	}()
}

func (p *Player) attack() {
	p.isAttacking = true
	go func() {
		<-time.After(200 * time.Millisecond)
		p.isAttacking = false
	}()
}

func (p *Player) respawn() {
	p.hearts = 5
	fmt.Println("respawning")
	p.invincible = true
	p.x = randFloat32(0, float32(S_WIDTH))
	p.y = randFloat32(0, float32(S_HEIGHT))
	go func() {
		<-time.After(1 * time.Second)
		p.invincible = false
	}()
}