package gomath

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"hash/fnv"
	"math"
)

type Spatial interface {
	GetValues() []float64
	SetValues([]float64)
	Size() int
	X() float64
	Y() float64
	Z() float64
	W() float64
}

func HashSpatial(spatial Spatial) int64 {
	hasher := fnv.New64a()

	for _, value := range spatial.GetValues() {
		bits := math.Float64bits(value)
		buf := make([]byte, 8)
		for i := 0; i < 8; i++ {
			buf[i] = byte(bits >> (i * 8))
		}
		_, _ = hasher.Write(buf)
	}

	return int64(hasher.Sum64())
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
