package main

import (
	"image"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
    x, y     float32
    vx, vy   float32
	size int
	image *ebiten.Image
    isColliding bool
	frameIdx int
	hp int
	tintDuration int
}


func createEnemy() *Enemy {

	img, _, err := ebitenutil.NewImageFromFile("assets/player2.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Enemy{
		x: randFloat32(0, float32(S_WIDTH)),
		y: randFloat32(0, float32(S_HEIGHT)),
		size: 42,
		image: img,
		hp: 10,
	}
}


func (e *Enemy) draw(screen *ebiten.Image) {
	frameWidth := 32
	frameHeight := 32
	// cropping rect for the current frame
	frameX := (e.frameIdx % (e.image.Bounds().Dx() / frameWidth)) * frameWidth
	frameY := (e.frameIdx / (e.image.Bounds().Dx() / frameWidth)) * frameHeight
	cropRect := image.Rect(frameX, frameY, frameX+frameWidth, frameY+frameHeight)

	frame := e.image.SubImage(cropRect).(*ebiten.Image)
	opts := &ebiten.DrawImageOptions{}
	scaleX := float64(e.size) / float64(frameWidth)
	scaleY := float64(e.size) / float64(frameHeight)

	if e.isColliding {
		opts.ColorScale.Scale(0.2, 0.822, 0.814, 1)
	} else {
		opts.ColorScale.Scale(0.1, 0.722, 0.114, 1)
	}

	if e.tintDuration > 0 {
		opts.ColorScale.Scale(0.99, 0.222, 0.114, 1)
		e.tintDuration--
	}

	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(float64(e.x), float64(e.y))

	screen.DrawImage(frame, opts)
}

func (e *Enemy) update(playerX, playerY float32) {

	const friction = 0.95

	if (e.hp <= 0) {
		e.x = randFloat32(0, float32(S_WIDTH))
		e.y = randFloat32(0, float32(S_HEIGHT))
		e.hp = 10
	}

	// chance to move randomly
	chance := randFloat32(0, 1)
	if randFloat32(0, 1) < chance {
		e.vx = randFloat32(-0.5, 0.5) // Smaller movement range
		e.vy = randFloat32(-0.5, 0.5)
	}

	// Add some goal-seeking behavior toward the player with reduced frequency and intensity
	if randFloat32(0, 1) < 0.25 { // Reduced frequency (5% chance)
		dx := playerX - e.x
		dy := playerY - e.y
		mag := float32(math.Sqrt(float64(dx*dx + dy*dy)))
		if mag > 0 {
			// Smooth goal-seeking behavior
			e.vx += 0.05 * dx / mag // Gradual adjustment of velocity
			e.vy += 0.05 * dy / mag
		}
	}

	e.vx *= friction
	e.vy *= friction

	// pos
	e.x += e.vx
	e.y += e.vy

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

// apply velocity to position from damage source
func (e *Enemy) takeDamage(vx, vy float32) {
	e.hp--
	// apply a little knockback
	e.vx += vx
	e.vy += vy
	e.tintDuration = 5
}
