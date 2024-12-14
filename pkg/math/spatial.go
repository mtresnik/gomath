package math

import "fmt"

type Spatial interface {
	GetValues() []float64
	SetValues([]float64)
	Size() int
	X() float64
	Y() float64
	Z() float64
	W() float64
}

func ToVector(s Spatial) Vector {
	return Vector{s.GetValues()}
}

func ToPoint(sp Spatial) Point {
	return Point{sp.GetValues()}
}

func String(sp Spatial) string {
	return fmt.Sprintf("Values: %v", sp.GetValues())
}
