package utils

import (
	"crypto/sha256"
	"fmt"
	"golang-game/config"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func Collides(x1, y1, w1, h1, x2, y2, w2, h2 float32) bool {
	return x1 < x2+w2 && x1+w1 > x2 && y1 < y2+h2 && y1+h1 > y2
}

func DrawDebugBorder(screen *ebiten.Image, x, y, w, h float32) {
	if (!config.DEBUG) {
		return
	}

	opts := &ebiten.DrawImageOptions{}
	borderImg := ebiten.NewImage(int(w), int(h))
	borderImg.Fill(color.RGBA{0, 0, 0, 0}) // Transparent fill for the inside

	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%f%f", w, h)))

	borderCol := RandomColorFromHash(fmt.Sprintf("%x", hash.Sum(nil)))
	var lineSize float32 = 1

	vector.DrawFilledRect(borderImg, 0, 0, float32(w), lineSize, borderCol, true) // Blue top border
	vector.DrawFilledRect(borderImg, 0, float32(h)-lineSize, float32(w), lineSize, borderCol, true) // Blue bottom border
	vector.DrawFilledRect(borderImg, 0, 0, lineSize, float32(h), borderCol, true) // Blue left border
	vector.DrawFilledRect(borderImg, float32(w)-lineSize, 0, lineSize, float32(h), borderCol, true) // Blue right border

	// Apply the transformation (scale and translate)
	opts.GeoM.Translate(float64(x), float64(y))

	// Draw the border image
	screen.DrawImage(borderImg, opts)
}