package enemies

import (
	"golang-game/config"
	"golang-game/utils"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	x, y float32
    vx, vy float32
	size int
	image *ebiten.Image
	hp int
	frameIdx int
	tintDuration int
}


func (e *Enemy) Draw(screen *ebiten.Image) {
	utils.DrawDebugBorder(screen, e.x, e.y, float32(e.size), float32(e.size))
	if (config.DISABLE_DRAW) { return }
	//
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

	opts.ColorScale.Scale(0.1, 0.622, 0.414, 0.5)

	if e.tintDuration > 0 {
		opts.ColorScale.Scale(0.99, 0.222, 0.114, 1)
		e.tintDuration--
	}

	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(float64(e.x), float64(e.y))

	screen.DrawImage(frame, opts)
}

func (e *Enemy) Update(x, y float32) {
	playerX := x
	playerY := y

	const friction = 0.95

	// chance to move randomly
	chance := utils.RandFloat32(0, 1)
	if utils.RandFloat32(0, 1) < chance {
		e.vx = utils.RandFloat32(-0.5, 0.5) // Smaller movement range
		e.vy = utils.RandFloat32(-0.5, 0.5)
	}

	// Add some goal-seeking behavior toward the player with reduced frequency and intensity
	if utils.RandFloat32(0, 1) < 0.25 { // Reduced frequency (5% chance)
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
	} else if e.x > float32(config.S_WIDTH-16) {
		e.x = float32(config.S_WIDTH - 16)
		e.vx = -e.vx
	}

	if e.y < 0 {
		e.y = 0
		e.vy = -e.vy
	} else if e.y > float32(config.S_HEIGHT-16) {
		e.y = float32(config.S_HEIGHT - 16)
		e.vy = -e.vy
	}
}

func (e *Enemy) IsDead() bool {
	return e.hp <= 0
}

func (e *Enemy) Respawn() {
	e.x = utils.RandFloat32(0, float32(config.S_WIDTH))
	e.y = utils.RandFloat32(0, float32(config.S_HEIGHT))
	e.hp = 10
}

func (e *Enemy) GetPosition() (float32, float32) {
	return e.x, e.y
}

func (e *Enemy) GetSize() int {
	return e.size
}

func (e *Enemy) TakeDamage(vx, vy float32) {
	e.hp--
	// apply a little knockback
	e.vx += vx
	e.vy += vy
	e.tintDuration = 5
}

type EnemyInterface interface {
	Draw(screen *ebiten.Image)
	Update(float32, float32)
	IsDead() bool
	Respawn()
	GetPosition() (float32, float32)
	GetSize() int
	TakeDamage(float32, float32)
}


func CreateEnemies() []EnemyInterface {
	enemies := []EnemyInterface{}
	zombie := CreateZombie()
	enemies = append(enemies, zombie)
	return enemies
}



