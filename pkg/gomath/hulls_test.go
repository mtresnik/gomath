package gomath

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestConvexHull_Points(t *testing.T) {
	imageSize := 1000.0
	points := RandomPoints(20, 2, imageSize)
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
			if polygon.Contains(NewPoint(float64(col), float64(row))) {
				img.Set(col, row, color.RGBA{255, 0, 0, 255})
			}
		}
	}
	for _, point := range points {
		goutils.FillCircle(img, int(point.X()), int(point.Y()), 10, image.Black)
	}
	file, err := os.Create("TestConvexHull_Points.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}
