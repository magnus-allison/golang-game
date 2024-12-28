package ui

import (
	"io"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	fontFace font.Face
)

func init() {
	fontFile, err := os.Open("assets/fonts/Lato/Lato-Regular.ttf")
	if err != nil {
		log.Fatalf("failed to open font file: %v", err)
	}
	defer fontFile.Close()
	fontData, err := io.ReadAll(fontFile)
	if err != nil {
		log.Fatalf("failed to read font data: %v", err)
	}

	tt, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	const dpi = 72
	fontFace, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}
}
