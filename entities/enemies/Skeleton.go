package enemies

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Skeleton struct {
    *Enemy
}

func CreateSkeleton() *Skeleton {
	img, _, err := ebitenutil.NewImageFromFile("assets/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Skeleton{
		CreateEnemy(&EnemyParams{
			name: "ZombieBoss",
			w:    35,
			h:    55,
			img:  img,
			hp:   5,
			WalkMap: &WalkMap{
				N: []int{0, 8},
				W: []int{9, 17},
				S: []int{18, 26},
				E: []int{27, 35},
			},
		}),
	}
}
