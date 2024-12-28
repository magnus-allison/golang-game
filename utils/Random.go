package utils

import (
	"golang-game/config"
	"math/rand"
)

func RandInt(min, max int) int {
	return min + rand.Intn(max - min)
}

func RandFloat32(min, max float32) float32 {
	return min + rand.Float32() * (max - min)
}

func RandomPosition() (float32, float32) {
	return RandFloat32(0, float32(config.S_WIDTH)), RandFloat32(0, float32(config.S_HEIGHT))
}