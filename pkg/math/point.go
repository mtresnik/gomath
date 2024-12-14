package math

type Point struct {
	Values []float64
}

func (a Point) GetValues() []float64 {
	return a.Values
}

func (a Point) SetValues(values []float64) {
	a.Values = values
}

func (a Point) ToVector() Vector {
	return Vector{a.Values}
}

func (a Point) X() float64 {
	if len(a.Values) < 1 {
		return 0.0
	}
	return a.Values[0]
}

func (a Point) Y() float64 {
	if len(a.Values) < 2 {
		return 0.0
	}
	return a.Values[1]
}

func (a Point) Z() float64 {
	if len(a.Values) < 3 {
		return 0.0
	}
	return a.Values[2]
}

func (a Point) W() float64 {
	if len(a.Values) < 4 {
		return 0.0
	}
	return a.Values[3]
}

func (a Point) Add(other Point) Point {
	return a.ToVector().Add(other.ToVector()).ToPoint()
}

func (a Point) AddVector(other Vector) Point {
	return a.ToVector().Add(other).ToPoint()
}

func (a Point) Subtract(other Point) Vector {
	return a.ToVector().Subtract(other.ToVector())
}

func (a Point) DistanceTo(other Point, function ...DistanceFunction) float64 {
	if len(function) > 0 {
		return function[0].Eval(a, other)
	}
	return EuclideanDistance{}.Eval(a, other)
}

func (a Point) Size() int {
	return len(a.Values)
}

func (a Point) Equals(other Point) bool {
	return a.ToVector().Equals(other.ToVector())
}
