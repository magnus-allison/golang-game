package enemies

import (
	"golang-game/config"
	"golang-game/utils"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type WalkMap struct {
	N []int
	S []int
	E []int
	W []int
}

type Enemy struct {
	name string
	w, h int
	x, y float32
    vx, vy float32
	img *ebiten.Image
	hp int
	frameIdx int
	tintDuration int
	frameWidth, frameHeight int
	walkMap *WalkMap
	animCounter int
	targetX, targetY float32
	pauseTimer int
}

type EnemyParams struct {
	name string
	w, h int
	img *ebiten.Image
	hp int
	WalkMap *WalkMap
	frameWidth, frameHeight int
}

func CreateEnemy(enemyParams *EnemyParams) *Enemy {
	enemy := &Enemy{
		name: enemyParams.name,
		x: utils.RandFloat32(0, float32(config.S_WIDTH)),
		y: utils.RandFloat32(0, float32(config.S_HEIGHT)),
		w: enemyParams.w,
		h: enemyParams.h,
		img: enemyParams.img,
		hp: enemyParams.hp,
		frameWidth: enemyParams.w,
		frameHeight: enemyParams.h,
	}

	enemy.walkMap = &WalkMap{
		N: []int{0, 1},
		S: []int{2, 3},
		E: []int{4, 5},
		W: []int{6, 7},
	}

	if (enemyParams.frameWidth != 0) {
		enemy.frameWidth = enemyParams.frameWidth
	}
	if (enemyParams.frameHeight != 0) {
		enemy.frameHeight = enemyParams.frameHeight
	}

	if (enemyParams.WalkMap != nil) {
		enemy.walkMap = enemyParams.WalkMap
	}

	return enemy
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	utils.DrawDebugBorder(screen, e.x, e.y, float32(e.w), float32(e.h))
	if (config.DISABLE_DRAW) { return }
	frameWidth := e.frameWidth
	frameHeight := e.frameHeight
	frameX := (e.frameIdx % (e.img.Bounds().Dx() / frameWidth)) * frameWidth
	frameY := (e.frameIdx / (e.img.Bounds().Dx() / frameWidth)) * frameHeight
	cropRect := image.Rect(frameX, frameY, frameX+frameWidth, frameY+frameHeight)

	frame := e.img.SubImage(cropRect).(*ebiten.Image)
	opts := &ebiten.DrawImageOptions{}
	scaleX := float64(e.w) / float64(frameWidth)
	scaleY := float64(e.h) / float64(frameHeight)

	if e.tintDuration > 0 {
		opts.ColorScale.Scale(0.99, 0.222, 0.114, 1)
		e.tintDuration--
	}

	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(float64(e.x), float64(e.y))

	screen.DrawImage(frame, opts)
}

const (
	playerDetectionRange = 0.0
	pauseDuration = 60
)

func (e *Enemy) Update(x, y float32) {
	playerX := x
	playerY := y

	const friction = 0.95

	dx := playerX - e.x
	dy := playerY - e.y
	distanceToPlayer := float32(math.Sqrt(float64(dx*dx + dy*dy)))

	// Check if the player is within range
	if distanceToPlayer < playerDetectionRange {
		// Move towards the player
		e.moveTowardsPlayer(dx, dy)

		// Start pause behavior
		if e.pauseTimer <= 0 {
			e.vx += 0.1 * dx / distanceToPlayer
			e.vy += 0.1 * dy / distanceToPlayer
		}
	} else {
		// Enemy is not in range of the player, pick a random target point
		if e.pauseTimer <= 0 {
			// Pick a new random target point
			e.targetX = utils.RandFloat32(0, float32(config.S_WIDTH))
			e.targetY = utils.RandFloat32(0, float32(config.S_HEIGHT))

			// Set pause timer (pause before moving to the new target)
			e.pauseTimer = pauseDuration
		}

		// Move towards the random target point
		e.moveToTarget(e.targetX, e.targetY)
	}

	e.vx *= friction
	e.vy *= friction

	// pos
	e.x += e.vx
	e.y += e.vy

	if (e.vx > 0) {
		e.animationFrame(e.walkMap.E[0], e.walkMap.E[1])
	} else if (e.vx < 0) {
		e.animationFrame(e.walkMap.W[0], e.walkMap.W[1])
	} else if (e.vy > 0) {
		e.animationFrame(e.walkMap.S[0], e.walkMap.S[1])
	} else if (e.vy < 0) {
		e.animationFrame(e.walkMap.N[0], e.walkMap.N[1])
	}

	// Decrement pause timer
	if e.pauseTimer > 0 {
		e.pauseTimer--
	}

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

func (e *Enemy) moveTowardsPlayer(dx, dy float32) {
	mag := float32(math.Sqrt(float64(dx*dx + dy*dy)))
	if mag > 0 {
		// Gradually adjust velocity to move towards the player
		e.vx += 0.1 * dx / mag
		e.vy += 0.1 * dy / mag
	}
}

// Move towards the random target position
func (e *Enemy) moveToTarget(targetX, targetY float32) {
	dx := targetX - e.x
	dy := targetY - e.y
	mag := float32(math.Sqrt(float64(dx*dx + dy*dy)))

	if mag > 0 {
		// Gradual movement towards the target
		e.vx += 0.05 * dx / mag
		e.vy += 0.05 * dy / mag
	}
}

func (e *Enemy) animationFrame(start int, end int) {
	const frameRate = 8
	if (e.frameIdx < start || e.frameIdx > end) {
		e.frameIdx = start
	}
	e.animCounter++
	if e.animCounter >= frameRate {
		e.frameIdx++
		if e.frameIdx > end { e.frameIdx = start }
		e.animCounter = 0
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

func (e *Enemy) GetSize() (int, int) {
	return e.w, e.h
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
	GetSize() (int, int)
	TakeDamage(float32, float32)
}


func CreateEnemies() []EnemyInterface {
	enemies := []EnemyInterface{}
	// add 10 zombies
	for i := 0; i < 10; i++ {
		enemies = append(enemies, CreateZombie())
	}
	// add 5 skeletons
	for i := 0; i < 5; i++ {
		enemies = append(enemies, CreateSkeleton())
	}

	// add zombie boss
	enemies = append(enemies, CreateZombieBoss())
	return enemies
}



