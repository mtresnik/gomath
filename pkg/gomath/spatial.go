package gomath

import "github.com/mtresnik/goutils/pkg/goutils"

type Spatial interface {
	GetValues() []float64
	SetValues([]float64)
	Size() int
	X() float64
	Y() float64
	Z() float64
	W() float64
}

func ToVector(sp Spatial) Vector {
	return Vector{sp.GetValues()}
}

func ToPoint(sp Spatial) Point {
	return Point{sp.GetValues()}
}

func SpatialString(sp Spatial, braces ...string) string {
	return goutils.Float64ArrayToString(sp.GetValues(), braces...)
}
