package gomath

import (
	"github.com/mtresnik/goutils/pkg/goutils"
)

type Spatial interface {
	GetValues() []float64
	Size() int
	X() float64
	Y() float64
	Z() float64
	W() float64
}

func HashSpatial(spatial Spatial) int64 {
	return goutils.HashFloats(spatial.GetValues()...)
}

func ToVector(sp Spatial) Vector {
	v, ok := sp.(Vector)
	if ok {
		return v
	}
	return Vector{sp.GetValues()}
}

func ToPoint(sp Spatial) Point {
	p, ok := sp.(Point)
	if ok {
		return p
	}
	return Point{sp.GetValues()}
}

func SpatialString(sp Spatial, braces ...string) string {
	return goutils.Float64ArrayToString(sp.GetValues(), braces...)
}
