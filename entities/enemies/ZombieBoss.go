package enemies

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ZombieBoss struct {
	Enemy
	phase       int
	attackPower int
}

func CreateZombieBoss() *Zombie {
	img, _, err := ebitenutil.NewImageFromFile("assets/goblinbozz.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Zombie{
		CreateEnemy(&EnemyParams{
			name: "Zombie",
			w:    100,
			h:    115,
			img:  img,
			hp:   33,
			frameHeight: 64,
			frameWidth: 64,
			WalkMap: &WalkMap{
				N: []int{4, 7},
				S: []int{0, 3},
				E: []int{8, 11},
				W: []int{12, 15},
			},
		}),
	}
}

