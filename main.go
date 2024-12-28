package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"golang-game/config"
	"golang-game/levels"
	"golang-game/ui"
)

type Game struct{
	State GameState
	Player *Player
	Enemies []*Enemy
	UI *ui.UI
	Projectiles []*Projectile
	LevelManager *levels.LevelManager
}

func (g *Game) Update() error {
	p := g.Player
	p.update()
	p.checkCollision(g.Enemies)
	for _, enemy := range g.Enemies {
		enemy.update(p.x, p.y)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if (p.canShoot) {
			p.canShoot = false
			g.Projectiles = append(g.Projectiles, createProjectile(int(p.x) + (p.size / 2), int(p.y) + 20, mx, my))
			go func() {
				// fmt.Print("can_s")
				<-time.After(200 * time.Millisecond)
				p.canShoot = true
			}()
		}
	}

	for i := len(g.Projectiles) - 1; i >= 0; i-- {
		projectile := g.Projectiles[i]
		projectile.update()
		projectile.checkCollision(g.Enemies)

		// cleanup
		if projectile.x < 0 || projectile.x > config.S_WIDTH || projectile.y < 0 || projectile.y > config.S_HEIGHT {
			g.Projectiles[i] = g.Projectiles[len(g.Projectiles)-1]
			g.Projectiles = g.Projectiles[:len(g.Projectiles)-1]
		}
	}

	g.State.CheckDeadEnemies(g.Enemies)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	p := g.Player
	g.LevelManager.Level.DrawFloor(screen)
	p.draw(screen)
	for _, enemy := range g.Enemies {
		enemy.draw(screen)
	}
	for _, projectile := range g.Projectiles {
		// lenx := len(g.Projectiles)
		// ebitenutil.DebugPrint(screen, "Entites: " + strconv.Itoa(lenx))
		projectile.draw(screen)
	}

	g.UI.DrawPlayerHearts(screen, p.hearts)
	g.UI.DrawPlayerScore(screen, g.State.GetScore())
	g.UI.DrawCursor(screen)

	if (p.hearts <= 0) {
		g.UI.DrawGameOverScreen(screen)
		go func() {
			<-time.After(3 * time.Second)
			p.respawn()
			g.Projectiles = []*Projectile{}
		}()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.S_WIDTH, config.S_HEIGHT
}

func main() {
	ebiten.SetWindowSize(config.S_WIDTH, config.S_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	player := createPlayer()
	player.respawn()
	enemies := [10]*Enemy{}
	for i := range enemies {
		enemies[i] = createEnemy()
	}
	if err := ebiten.RunGame(&Game{
		State: GameState{},
		Player: player,
		Enemies: enemies[:],
		UI: ui.CreateUI(),
		Projectiles: []*Projectile{},
		LevelManager: levels.CreateLevelManager(),
	}); err != nil {
		log.Fatal(err)
	}
}