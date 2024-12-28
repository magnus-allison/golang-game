package entities

import (
	"golang-game/config"
	"golang-game/entities/enemies"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


type EntityManager struct {
	Player *Player
	Enemies []enemies.EnemyInterface
	Projectiles []*Projectile
	PowerUps []*PowerUp
}

func CreateEntityManager() *EntityManager {

	powerUps := []*PowerUp{
		createPowerUp(),
	}

	player := CreatePlayer()
	player.Respawn()

	return &EntityManager{
		PowerUps: powerUps,
		Player: player,
		Enemies: enemies.CreateEnemies(),
	}
}

func (em *EntityManager) UpdateEntities() {
	for _, p := range em.PowerUps {
		p.update()
	}
	for _, e := range em.Enemies {
		e.Update(em.Player.X, em.Player.Y)
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
		e.Draw(screen)
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
		if e.IsDead() {
			e.Respawn()
			gs.IncrementScore()
		}
	}
}