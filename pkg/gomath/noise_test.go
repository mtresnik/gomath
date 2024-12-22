package gomath

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"image"
	"image/color"
	"image/gif"
	"os"
	"testing"
)

func TestPerlinNoise(t *testing.T) {
	size := 200
	duration := 50
	scalar := 4.0
	resultScalar := 2.0
	images := make([]*image.Paletted, 0)

	for frame := 0; frame < duration; frame++ {
		z := scalar * float64(frame+1) / float64(duration)
		img := image.NewRGBA(image.Rect(0, 0, size, size))
		goutils.FillRectangle(img, 0, 0, size, size, color.White)
		for row := 0; row < size; row++ {
			for col := 0; col < size; col++ {
				img.Set(col, row, goutils.GradientGreenToRed(resultScalar*PerlinNoise(float64(col+1)/float64(size), float64(row)/float64(size), z)))
			}
		}
		paletted := goutils.ConvertImageToPaletted(img)
		images = append(images, paletted)
		println("frame:", len(images))
	}
	delays := make([]int, len(images))
	for i := 0; i < len(images); i++ {
		delays[i] = 5
	}
	retGif := &gif.GIF{
		Image:     images,
		Delay:     delays,
		LoopCount: 0,
	}
	f, err := os.Create("TestPerlinNoise.gif")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	// Encode and save the GIF
	err = gif.EncodeAll(f, retGif)
	if err != nil {
		panic(err)
	}
}