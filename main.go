package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{
	Player *Player
	Enemies []*Enemy
	UI *UI
	Projectiles []*Projectile
}

var S_WIDTH = 1080
var S_HEIGHT = 720

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
				fmt.Print("can_s")
				<-time.After(200 * time.Millisecond)
				p.canShoot = true
			}()
		}
	}

	for i := len(g.Projectiles) - 1; i >= 0; i-- {
		projectile := g.Projectiles[i]
		projectile.update()
		projectile.checkCollision(g.Enemies)

		if projectile.x < 0 || projectile.x > S_WIDTH || projectile.y < 0 || projectile.y > S_HEIGHT {
			g.Projectiles[i] = g.Projectiles[len(g.Projectiles)-1]
			g.Projectiles = g.Projectiles[:len(g.Projectiles)-1]
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	p := g.Player
	vector.DrawFilledRect(screen, 0, 0, float32(S_WIDTH), float32(S_HEIGHT), color.RGBA{22, 33, 43, 255}, true)

	p.draw(screen)
	for _, enemy := range g.Enemies {
		enemy.draw(screen)
	}
	for _, projectile := range g.Projectiles {
		lenx := len(g.Projectiles)
		ebitenutil.DebugPrint(screen, "Entites: " + strconv.Itoa(lenx))
		projectile.draw(screen)
	}

	g.UI.drawPlayerHearts(screen, p.hearts)
	if (p.hearts <= 0) {
		g.UI.drawGameOverScreen(screen)
		go func() {
			<-time.After(3 * time.Second)
			p.respawn()
			g.Projectiles = []*Projectile{}
		}()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return S_WIDTH, S_HEIGHT
}

func main() {
	ebiten.SetWindowSize(S_WIDTH, S_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	player := createPlayer()
	player.respawn()
	enemies := [10]*Enemy{}
	for i := range enemies {
		enemies[i] = createEnemy()
	}
	if err := ebiten.RunGame(&Game{
		Player: player,
		Enemies: enemies[:],
		UI: createUI(),
		Projectiles: []*Projectile{},
	}); err != nil {
		log.Fatal(err)
	}
}