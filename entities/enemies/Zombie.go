package enemies

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"golang-game/config"
	"golang-game/utils"
)

type Zombie struct {
    *Enemy
}

func CreateZombie() *Zombie {
	img, _, err := ebitenutil.NewImageFromFile("assets/player2.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Zombie{
		&Enemy{
			x: utils.RandFloat32(0, float32(config.S_WIDTH)),
			y: utils.RandFloat32(0, float32(config.S_HEIGHT)),
			size: 39,
			image: img,
			hp: 5,
			tintDuration: 0,
		},
	}
}
