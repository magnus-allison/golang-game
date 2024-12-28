package entities

import (
	"fmt"
	"image"
	"log"
	"reflect"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"golang-game/config"
	"golang-game/entities/enemies"
	"golang-game/utils"
)

type Player struct{
	X, Y float32
	vx, vy float32
	Size int
	isAttacking bool
	Hearts int
	image *ebiten.Image
	frameIdx, animCounter int
	invincible bool
	tintDuration int
	canShoot bool
	isShooting bool
	walkSpeed float32
}

func CreatePlayer() *Player {
	img, _, err := ebitenutil.NewImageFromFile("assets/player2.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		Size: 42,
		image: img,
		canShoot: true,
		walkSpeed: 1.5,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {

	utils.DrawDebugBorder(screen, p.X, p.Y, float32(p.Size), float32(p.Size))
	if (config.DISABLE_DRAW) { return }

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

	scaleX := float64(p.Size) / float64(frameWidth)
	scaleY := float64(p.Size) / float64(frameHeight)
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(frame, opts)
}


func (p *Player) Update() {
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

	if !ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyS) && !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
		p.frameIdx = 0
	}

	p.vx *= friction
	p.vy *= friction

	if p.vx > p.walkSpeed {
		p.vx = p.walkSpeed
	} else if p.vx < -p.walkSpeed {
		p.vx = -p.walkSpeed
	}
	if p.vy > p.walkSpeed {
		p.vy = p.walkSpeed
	} else if p.vy < -p.walkSpeed {
		p.vy = -p.walkSpeed
	}

	p.X += p.vx
	p.Y += p.vy

	// bound check
	if p.X < 0 {
		p.X = 0
		p.vx = 0
	} else if p.X > float32(config.S_WIDTH - p.Size) {
		p.X = float32(config.S_WIDTH - p.Size)
		p.vx = 0
	}
	if p.Y < 0 {
		p.Y = 0
		p.vy = 0
	} else if p.Y > float32(config.S_HEIGHT-p.Size) {
		p.Y = float32(config.S_HEIGHT - p.Size)
		p.vy = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.attack()
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		p.isShooting = true
		p.shoot()
	} else {
		p.isShooting = false
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

func (p *Player) CheckCollision(entities interface{}) {
	switch e := entities.(type) {
	case []enemies.EnemyInterface:
		for _, enemy := range e {
			eX, eY := enemy.GetPosition()
			if utils.Collides(p.X, p.Y, float32(p.Size), float32(p.Size), eX, eY, float32(enemy.GetSize()), float32(enemy.GetSize())) {
				p.takeDamage()
			}
		}
	case []*PowerUp:
		for _, powerUp := range e {
			if utils.Collides(p.X, p.Y, float32(p.Size), float32(p.Size), powerUp.x, powerUp.y, float32(powerUp.w), float32(powerUp.h)) {
				p.powerUp()
				powerUp.x, powerUp.y = utils.RandomPosition()
			}
		}
	default:
		fmt.Println("Unknown type ", reflect.TypeOf(entities))
	}
}

func (p *Player) takeDamage() {
	if (p.invincible) {
		return
	}
	p.tintDuration = 30
	p.Hearts--
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

func (p *Player) Respawn() {
	p.Hearts = 5
	p.invincible = true
	p.X = utils.RandFloat32(0, float32(config.S_WIDTH))
	p.Y = utils.RandFloat32(0, float32(config.S_HEIGHT))
	go func() {
		<-time.After(1 * time.Second)
		p.invincible = false
	}()
}

func (p *Player) powerUp() {
	p.Hearts++
	p.walkSpeed += 0.03
}

func (p *Player) shoot() {
	if (p.canShoot) {
		p.canShoot = false
		go func() {
			<-time.After(200 * time.Millisecond)
			p.canShoot = true
		}()
	}
}