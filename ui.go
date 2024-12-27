package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// vector.DrawFilledRect(screen, enemy.x, enemy.y, 16, 16, enemy.color, true)
func drawPlayerHearts(screen *ebiten.Image, amount int) {
	size := 32
	for i := 0; i < amount; i++ {
		vector.DrawFilledRect(screen, float32((S_WIDTH - size * 2) - (i*size) - 5*i), float32(size),
			float32(size), float32(size), color.RGBA{255, 0, 0, 255}, true)
	}
}