package enemies

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Zombie struct {
    *Enemy
}

func CreateZombie() *Zombie {
	img, _, err := ebitenutil.NewImageFromFile("assets/goblin.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Zombie{
		CreateEnemy(&EnemyParams{
			name: "Zombie",
			w:    26,
			h:    33,
			img:  img,
			hp:   5,
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
