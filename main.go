package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{
	Player *Player
	Enemies []*Enemy
	UI *UI
}

var S_WIDTH = 640
var S_HEIGHT = 480

func (g *Game) Update() error {
	g.Player.update()
	g.Player.checkCollision(g.Enemies)
	for _, enemy := range g.Enemies {
		enemy.update(g.Player.x, g.Player.y)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.draw(screen)
	for _, enemy := range g.Enemies {
		vector.DrawFilledRect(screen, enemy.x, enemy.y, float32(enemy.size), float32(enemy.size), enemy.color, true)
	}
	g.UI.drawPlayerHearts(screen, g.Player.hearts)
	if (g.Player.hearts <= 0) {
		ebitenutil.DebugPrint(screen, "Game Over!")
		go func() {
			<-time.After(3 * time.Second)
			g.Player.hearts = 3
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
	enemies := [10]*Enemy{}
	for i := range enemies {
		enemies[i] = createEnemy()
	}
	if err := ebiten.RunGame(&Game{
		Player: player,
		Enemies: enemies[:],
		UI: createUI(),
	}); err != nil {
		log.Fatal(err)
	}
}