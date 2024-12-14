package geometry

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
