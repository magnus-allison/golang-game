package entities

import (
	"golang-game/config"
	"golang-game/utils"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


type EntityManager struct {
	Player *Player
	Enemies []*Enemy
	Projectiles []*Projectile
	PowerUps []*PowerUp
}

func CreateEntityManager() *EntityManager {

	powerUps := []*PowerUp{
		createPowerUp(),
	}

	player := CreatePlayer()
	player.Respawn()

	enemies := []*Enemy{}
	enemyCount := 12
	for i := 0; i < enemyCount; i++ {
		enemies = append(enemies, CreateEnemy())
	}

	return &EntityManager{
		PowerUps: powerUps,
		Player: player,
		Enemies: enemies,
	}
}

func (em *EntityManager) UpdateEntities() {
	for _, p := range em.PowerUps {
		p.update()
	}
	for _, e := range em.Enemies {
		e.update(em.Player)
	}

	// handle player shooting
	if (em.Player.canShoot && em.Player.isShooting) {
		mx, my := ebiten.CursorPosition()
		em.Projectiles = append(em.Projectiles, createProjectile(int(em.Player.X) + (em.Player.Size / 2), int(em.Player.Y) + 20, mx, my, 1.0))
	}
	for _, p := range em.Projectiles {
		p.update()
		p.checkCollision(em.Enemies)
	}

	em.DeleteProjectiles()
}

func (em *EntityManager) DrawEntities(screen *ebiten.Image) {
	for _, p := range em.PowerUps {
		p.draw(screen)
	}
	for _, e := range em.Enemies {
		e.draw(screen)
	}
	for _, p := range em.Projectiles {
		p.draw(screen)
	}

	em.Player.Draw(screen)

	if (config.DEBUG) {
		len := strconv.Itoa(len(em.Projectiles))
		ebitenutil.DebugPrint(screen, "ProjectileCount: " + len)
	}
}

func (em *EntityManager) DeleteProjectiles() {
	for i := len(em.Projectiles) - 1; i >= 0; i-- {
		p := em.Projectiles[i]
		if p.x < 0 || p.x > config.S_WIDTH || p.y < 0 || p.y > config.S_HEIGHT {
			em.Projectiles[i] = em.Projectiles[len(em.Projectiles)-1]
			em.Projectiles = em.Projectiles[:len(em.Projectiles)-1]
		}
	}
}

func (em *EntityManager) CheckForDeadEnemies(gs interface { IncrementScore() }) {
	for _, e := range em.Enemies {
		if (e.hp <= 0) {
			e.x = utils.RandFloat32(0, float32(config.S_WIDTH))
			e.y = utils.RandFloat32(0, float32(config.S_HEIGHT))
			e.hp = 10
			gs.IncrementScore()
		}
	}
}