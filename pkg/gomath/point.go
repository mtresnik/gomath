package gomath

type Point struct {
	Values []float64
}

func (p Point) String() string {
	return SpatialString(p, "(", ")")
}

func (p Point) GetValues() []float64 {
	return p.Values
}

func (p Point) SetValues(values []float64) {
	p.Values = values
}

func (p Point) ToVector() Vector {
	return Vector{p.Values}
}

func (p Point) X() float64 {
	if len(p.Values) < 1 {
		return 0.0
	}
	return p.Values[0]
}

func (p Point) Y() float64 {
	if len(p.Values) < 2 {
		return 0.0
	}
	return p.Values[1]
}

func (p Point) Z() float64 {
	if len(p.Values) < 3 {
		return 0.0
	}
	return p.Values[2]
}

func (p Point) W() float64 {
	if len(p.Values) < 4 {
		return 0.0
	}
	return p.Values[3]
}

func (p Point) Add(other Point) Point {
	return p.ToVector().Add(other.ToVector()).ToPoint()
}

func (p Point) AddVector(other Vector) Point {
	return p.ToVector().Add(other).ToPoint()
}

func (p Point) Subtract(other Point) Vector {
	return p.ToVector().Subtract(other.ToVector())
}

func (p Point) DistanceTo(other Point, function ...DistanceFunction) float64 {
	if len(function) > 0 {
		return function[0].Eval(p, other)
	}
	return EuclideanDistance{}.Eval(p, other)
}

func (p Point) Size() int {
	return len(p.Values)
}

func (p Point) Equals(other Point) bool {
	return p.ToVector().Equals(other.ToVector())
}
