package main

import (
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max - min)
}

func randFloat32(min, max float32) float32 {
	return min + rand.Float32() * (max - min)
}

func randomPosition() (float32, float32) {
	return randFloat32(0, float32(S_WIDTH)), randFloat32(0, float32(S_HEIGHT))
}