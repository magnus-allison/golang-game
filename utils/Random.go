package utils

import (
	"encoding/hex"
	"golang-game/config"
	"image/color"
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

func RandomColor() color.RGBA {
	return color.RGBA{uint8(RandInt(0, 255)), uint8(RandInt(0, 255)), uint8(RandInt(0, 255)), 255}
}

func RandomColorFromHash(hash string) color.RGBA {
	hashBytes, err := hex.DecodeString(hash)
	if err != nil {
		panic(err)
	}
	r := uint8((int(hashBytes[0])*2 + int(hashBytes[1])*3) % 256)
	g := uint8((int(hashBytes[2])*5 + int(hashBytes[3])*7) % 256)
	b := uint8((int(hashBytes[4])*11 + int(hashBytes[5])*13) % 256)

	if r > 127 { r = 255 - r }
	if g > 127 { g = 255 - g }
	if b > 127 { b = 255 - b }

	return color.RGBA{r, g, b, 255}
}