package ui

import (
	"image"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"golang-game/config"
)

type UI struct {
	playerHeartImage *ebiten.Image
}

func CreateUI() *UI {
	img, _, err := ebitenutil.NewImageFromFile("assets/hearts.png")
	if err != nil {
		log.Fatal(err)
	}
	return &UI{
		playerHeartImage: img,
	}
}

func (ui *UI) DrawPlayerHearts(screen *ebiten.Image, amount int) {

	size := 32

	for i := 0; i < amount; i++ {
		rect := image.Rect(0, 0, 16, 16)
		croppedImage := ui.playerHeartImage.SubImage(rect).(*ebiten.Image)

		// Create draw options
		opts := &ebiten.DrawImageOptions{}
		scaleX := float64(size) / float64(croppedImage.Bounds().Dx())
		scaleY := float64(size) / float64(croppedImage.Bounds().Dy())
		opts.GeoM.Scale(scaleX, scaleY)
		opts.GeoM.Translate(
			float64((config.S_WIDTH - size*2) - (i*size) - 0*i),
			float64(size / 2),
		)

		screen.DrawImage(croppedImage, opts)
	}
}

func (ui *UI) DrawGameOverScreen(screen *ebiten.Image) {
    // Draw a semi-transparent black background
    vector.DrawFilledRect(screen, 0, 0, float32(config.S_WIDTH), float32(config.S_HEIGHT), color.RGBA{0, 0, 0, 200}, true)

    // Draw the Game Over text in the middle of the screen
    message := "GAME OVER"

    // textWidth, textHeight := ebitenutil.DebugPrintAt(screen, message, S_WIDTH/2-70, S_HEIGHT/2-20) // Position the text
    ebitenutil.DebugPrintAt(screen, message, (config.S_WIDTH/2-len(message)*8/2), (config.S_HEIGHT/2-20))
}

func (ui *UI) DrawPlayerScore(screen *ebiten.Image, score int) {

	scoreStr := strconv.Itoa(score)
	x, y := 10, 30
	text.Draw(screen, "Score: "+scoreStr, fontFace, x, y, color.White)

}

func (ui *UI) DrawCursor(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	// if (config.DEBUG) {
		// fmt.Println("Mouse X: ", mx, "Mouse Y: ", my)
	// }
	cursorSize := 5
	vector.DrawFilledCircle(screen, float32(mx), float32(my), float32(cursorSize), color.RGBA{205, 205, 255, 0}, true)
	// vector.DrawFilledRect(screen, float32(mx-cursorSize/2), float32(my-cursorSize/2), float32(cursorSize), float32(cursorSize), color.RGBA{255, 255, 255, 255}, true)
}
