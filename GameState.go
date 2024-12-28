package main

type GameState struct {
	score int
}

func (gs *GameState) IncrementScore() {
	gs.score++
}

func (gs *GameState) GetScore() int {
	return gs.score
}