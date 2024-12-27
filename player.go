package main

import (
	"image"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct{
	x, y   float32
	vx, vy float32
	size int
	color color.RGBA
	isAttacking bool
	isColliding bool
	hearts int
	image *ebiten.Image
	frameIdx int
	animCounter int
	invincible bool
}

func createPlayer() *Player {

	img, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		x: 0,
		y: 0,
		size: 60,
		color: color.RGBA{255, 255, 255, 255},
		hearts: 3,
		image: img,
	}
}

func (p *Player) draw(screen *ebiten.Image) {
	frameWidth := 64
	frameHeight := 64
	// calculate the cropping rectangle for the current frame
	frameX := (p.frameIdx % (p.image.Bounds().Dx() / frameWidth)) * frameWidth
	frameY := (p.frameIdx / (p.image.Bounds().Dx() / frameWidth)) * frameHeight
	cropRect := image.Rect(frameX, frameY, frameX+frameWidth, frameY+frameHeight)

	frame := p.image.SubImage(cropRect).(*ebiten.Image)
	opts := &ebiten.DrawImageOptions{}
	scaleX := float64(p.size) / float64(frameWidth)
	scaleY := float64(p.size) / float64(frameHeight)
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(frame, opts)
}


func (p *Player) update() {
	const speed = 2.0        // Maximum speed
	const acceleration = 0.5 // How quickly the player accelerates
	const friction = 0.9     // Slow down when no key is pressed

	const frameRate = 8

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.animCounter++
		if p.animCounter >= frameRate {
			p.frameIdx++
			if p.frameIdx > 15 { // Cycle between frames 12 and 15
				p.frameIdx = 12
			}
			p.animCounter = 0
		}
		p.vy -= acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.animCounter++
		if p.animCounter >= frameRate {
			p.frameIdx++
			if p.frameIdx > 3 { // Cycle between frames 0 and 3
				p.frameIdx = 0
			}
			p.animCounter = 0
		}
		p.vy += acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.animCounter++
		if p.animCounter >= frameRate {
			p.frameIdx++
			if p.frameIdx > 7 { // Cycle between frames 4 and 7
				p.frameIdx = 4
			}
			p.animCounter = 0
		}
		p.vx -= acceleration
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.animCounter++
		if p.animCounter >= frameRate {
			p.frameIdx++
			if p.frameIdx > 11 { // Cycle between frames 8 and 11
				p.frameIdx = 8
			}
			p.animCounter = 0
		}
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
		// Calculate the bounding box for both the player and the enemy
		playerRight := p.x + float32(p.size)
		playerBottom := p.y + float32(p.size)
		enemyRight := enemy.x + float32(enemy.size)
		enemyBottom := enemy.y + float32(enemy.size)

		// Check for intersection between the player's and enemy's bounding boxes
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
	p.hearts--
	p.invincible = true

	go func() {
		// Set invincibility duration (e.g., 1 second)
		time.Sleep(1 * time.Second) // Sleep for 1 second (adjustable)
		p.invincible = false        // Reset invincibility flag after 1 second
	}()
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