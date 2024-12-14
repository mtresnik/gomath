package math

import (
	"math"
)

type Vector struct {
	Values []float64
}

func (a Vector) ToPoint() Point {
	return Point{a.Values}
}

func (a Vector) GetValues() []float64 {
	return a.Values
}

func (a Vector) SetValues(values []float64) {
	a.Values = values
}

func (a Vector) X() float64 {
	if len(a.Values) < 1 {
		return 0.0
	}
	return a.Values[0]
}

func (a Vector) Y() float64 {
	if len(a.Values) < 2 {
		return 0.0
	}
	return a.Values[1]
}

func (a Vector) Z() float64 {
	if len(a.Values) < 3 {
		return 0.0
	}
	return a.Values[2]
}

func (a Vector) W() float64 {
	if len(a.Values) < 4 {
		return 0.0
	}
	return a.Values[3]
}

func (a Vector) Add(other Vector) Vector {
	var newValues = make([]float64, max(len(a.Values), len(other.Values)))
	for i := 0; i < len(newValues); i++ {
		if i < len(a.Values) && i < len(other.Values) {
			newValues[i] = a.Values[i] + other.Values[i]
		} else if i < len(a.Values) {
			newValues[i] = a.Values[i]
		} else if i < len(other.Values) {
			newValues[i] = other.Values[i]
		} else {
			newValues[i] = 0
		}
	}
	return Vector{newValues}
}

func (a Vector) AddPoint(other Point) Point {
	added := a.Add(other.ToVector())
	return added.ToPoint()
}

func (a Vector) Subtract(other Vector) Vector {
	var newValues = make([]float64, max(len(a.Values), len(other.Values)))
	for i := 0; i < len(newValues); i++ {
		if i < len(a.Values) && i < len(other.Values) {
			newValues[i] = a.Values[i] - other.Values[i]
		} else if i < len(a.Values) {
			newValues[i] = a.Values[i]
		} else if i < len(other.Values) {
			newValues[i] = other.Values[i]
		} else {
			newValues[i] = 0
		}
	}
	return Vector{newValues}
}

func (a Vector) DotProduct(other Vector) float64 {
	over := min(len(a.Values), len(other.Values))
	sum := 0.0
	for i := 0; i < over; i++ {
		if i < len(a.Values) && i < len(other.Values) {
			sum += a.Values[i] * other.Values[i]
		}
	}
	return sum
}

func (a Vector) Magnitude() float64 {
	dot := a.DotProduct(a)
	return math.Sqrt(dot)
}

func (a Vector) CrossProduct(other Vector) (Vector, *error) {
	X := a.Y()*other.Z() - other.Y()*a.Z()
	Y := a.Z()*other.X() - other.Values[2]*a.X()
	Z := a.X()*other.Y() - other.X()*a.Y()
	return Vector{[]float64{X, Y, Z}}, nil
}

func (a Vector) Scale(other float64) Vector {
	newValues := make([]float64, len(a.Values))
	for i := 0; i < len(a.Values); i++ {
		newValues[i] = a.Values[i] * other
	}
	return Vector{newValues}
}

func (a Vector) Sum() float64 {
	sum := 0.0
	for i := 0; i < len(a.Values); i++ {
		sum += a.Values[i]
	}
	return sum
}

func (a Vector) Size() int {
	return len(a.Values)
}

func (a Vector) Normalize() Vector {
	magnitude := a.Magnitude()
	if magnitude == 0 || math.IsNaN(magnitude) || math.IsInf(magnitude, 0) {
		return a
	}
	return a.Scale(1 / magnitude)
}

func (a Vector) Equals(other Vector) bool {
	if a.Size() != other.Size() {
		return false
	}
	for i := 0; i < a.Size(); i++ {
		if a.Values[i] != other.Values[i] {
			return false
		}
	}
	return true
}
