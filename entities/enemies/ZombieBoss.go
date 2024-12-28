package enemies

import (
	"golang-game/config"
	"golang-game/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ZombieBoss struct {
	Enemy
	phase       int
	attackPower int
}

func CreateZombieBoss() *Zombie {
	img, _, err := ebitenutil.NewImageFromFile("assets/player2.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Zombie{
		&Enemy{
			x: utils.RandFloat32(0, float32(config.S_WIDTH)),
			y: utils.RandFloat32(0, float32(config.S_HEIGHT)),
			size: 64,
			image: img,
			hp: 5,
			tintDuration: 0,
		},
	}
}

