package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"golang-game/config"
	"golang-game/entities"
	"golang-game/levels"
	"golang-game/ui"
)

type Game struct{
	State GameState
	UI *ui.UI
	LevelManager *levels.LevelManager
	EntityManager *entities.EntityManager
}

func (g *Game) Update() error {
	p := g.EntityManager.Player

	g.EntityManager.UpdateEntities()
	p.Update()
	p.CheckCollision(g.EntityManager.Enemies)
	p.CheckCollision(g.EntityManager.PowerUps)

	g.EntityManager.CheckForDeadEnemies(&g.State)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	p := g.EntityManager.Player
	g.LevelManager.Level.DrawFloor(screen)
	g.EntityManager.DrawEntities(screen)

	g.UI.DrawPlayerHearts(screen, p.Hearts)
	g.UI.DrawPlayerScore(screen, g.State.GetScore())
	g.UI.DrawCursor(screen)

	if (p.Hearts <= 0) {
		g.UI.DrawGameOverScreen(screen)
		go func() {
			<-time.After(3 * time.Second)
			p.Respawn()
			// g.Projectiles = []*Projectile{}
		}()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.S_WIDTH, config.S_HEIGHT
}

func main() {
	ebiten.SetWindowSize(config.S_WIDTH, config.S_HEIGHT)
	ebiten.SetWindowTitle("Golang Game")
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	if err := ebiten.RunGame(&Game{
		State: GameState{},
		EntityManager: entities.CreateEntityManager(),
		LevelManager: levels.CreateLevelManager(),
		UI: ui.CreateUI(),
	}); err != nil {
		log.Fatal(err)
	}
}