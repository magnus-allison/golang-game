package main

import "golang-game/config"

type GameState struct {
	score int
}

func (gs *GameState) IncrementScore() {
	gs.score++
}

func (gs *GameState) GetScore() int {
	return gs.score
}

func (gs *GameState) CheckDeadEnemies(enemies []*Enemy) {
	for _, e := range enemies {
		if (e.hp <= 0) {
			e.x = randFloat32(0, float32(config.S_WIDTH))
			e.y = randFloat32(0, float32(config.S_HEIGHT))
			e.hp = 10
			gs.IncrementScore()
		}
	}
}
