package gomath

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/mtresnik/goutils/pkg/goutils"
)

func TestConvexHull_Points(t *testing.T) {
	imageSize := 1000.0
	points := RandomPoints(10, 2, imageSize)
	unqiuePoints := make([]Point, 0)
	for _, point := range points {
		contains := false
		for _, other := range unqiuePoints {
			dist := EuclideanDistance
			if dist(point, other) < 1.0 {
				contains = true
				break
			}
		}
		if !contains {
			unqiuePoints = append(unqiuePoints, point)
			println(point.String())
		}
	}
	hull := ConvexHull(unqiuePoints...)
	println("Hull: ", len(hull))
	polygon := NewPolygon(hull...)
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
	for _, point := range unqiuePoints {
		goutils.FillCircle(img, int(point.X()), int(point.Y()), 10, color.Black)
	}
	for _, point := range hull {
		goutils.FillCircle(img, int(point.X()), int(point.Y()), 10, color.RGBA{0, 0, 255, 255})
	}
	file, err := os.Create("TestConvexHull_Points.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}
