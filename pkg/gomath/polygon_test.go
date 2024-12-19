package gomath

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestPolygon_Contains(t *testing.T) {

	imageSize := 100
	points := []Point{
		*NewPoint(30, 30),
		*NewPoint(50, 50),
		*NewPoint(70, 30),
		*NewPoint(90, 20),
		*NewPoint(10, 20),
	}

	polygon := NewPolygon(points...)
	imgRect := image.Rect(0, 0, int(imageSize), int(imageSize))
	img := image.NewRGBA(imgRect)
	for row := 0; row < int(imageSize); row++ {
		for col := 0; col < int(imageSize); col++ {
			img.Set(col, row, image.White)
		}
	}
	for row := 0; row < int(imageSize); row++ {
		for col := 0; col < int(imageSize); col++ {
			if polygon.Contains(*NewPoint(float64(col), float64(row))) {
				img.Set(col, row, color.RGBA{255, 0, 0, 255})
			}
		}
	}

	file, err := os.Create("TestPolygon_Contains.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)

}
