package math

import (
	"math"
)

type Vector struct {
	Values []float64
}

func (v Vector) String() string {
	return SpatialString(v, "<", ">")
}

func (v Vector) ToPoint() Point {
	return Point{v.Values}
}

func (v Vector) GetValues() []float64 {
	return v.Values
}

func (v Vector) SetValues(values []float64) {
	v.Values = values
}

func (v Vector) X() float64 {
	if len(v.Values) < 1 {
		return 0.0
	}
	return v.Values[0]
}

func (v Vector) Y() float64 {
	if len(v.Values) < 2 {
		return 0.0
	}
	return v.Values[1]
}

func (v Vector) Z() float64 {
	if len(v.Values) < 3 {
		return 0.0
	}
	return v.Values[2]
}

func (v Vector) W() float64 {
	if len(v.Values) < 4 {
		return 0.0
	}
	return v.Values[3]
}

func (v Vector) Add(other Vector) Vector {
	var newValues = make([]float64, max(len(v.Values), len(other.Values)))
	for i := 0; i < len(newValues); i++ {
		if i < len(v.Values) && i < len(other.Values) {
			newValues[i] = v.Values[i] + other.Values[i]
		} else if i < len(v.Values) {
			newValues[i] = v.Values[i]
		} else if i < len(other.Values) {
			newValues[i] = other.Values[i]
		} else {
			newValues[i] = 0
		}
	}
	return Vector{newValues}
}

func (v Vector) AddPoint(other Point) Point {
	added := v.Add(other.ToVector())
	return added.ToPoint()
}

func (v Vector) Subtract(other Vector) Vector {
	var newValues = make([]float64, max(len(v.Values), len(other.Values)))
	for i := 0; i < len(newValues); i++ {
		if i < len(v.Values) && i < len(other.Values) {
			newValues[i] = v.Values[i] - other.Values[i]
		} else if i < len(v.Values) {
			newValues[i] = v.Values[i]
		} else if i < len(other.Values) {
			newValues[i] = other.Values[i]
		} else {
			newValues[i] = 0
		}
	}
	return Vector{newValues}
}

func (v Vector) DotProduct(other Vector) float64 {
	over := min(len(v.Values), len(other.Values))
	sum := 0.0
	for i := 0; i < over; i++ {
		if i < len(v.Values) && i < len(other.Values) {
			sum += v.Values[i] * other.Values[i]
		}
	}
	return sum
}

func (v Vector) Magnitude() float64 {
	dot := v.DotProduct(v)
	return math.Sqrt(dot)
}

func (v Vector) CrossProduct(other Vector) (Vector, *error) {
	X := v.Y()*other.Z() - other.Y()*v.Z()
	Y := v.Z()*other.X() - other.Values[2]*v.X()
	Z := v.X()*other.Y() - other.X()*v.Y()
	return Vector{[]float64{X, Y, Z}}, nil
}

func (v Vector) Multiply(other Vector) Vector {
	var newValues = make([]float64, max(len(v.Values), len(other.Values)))
	for i := 0; i < len(newValues); i++ {
		if i < len(v.Values) && i < len(other.Values) {
			newValues[i] = v.Values[i] * other.Values[i]
		} else {
			newValues[i] = 0
		}
	}
	return Vector{newValues}
}

func (v Vector) Scale(other float64) Vector {
	newValues := make([]float64, len(v.Values))
	for i := 0; i < len(v.Values); i++ {
		newValues[i] = v.Values[i] * other
	}
	return Vector{newValues}
}

func (v Vector) Sum() float64 {
	sum := 0.0
	for i := 0; i < len(v.Values); i++ {
		sum += v.Values[i]
	}
	return sum
}

func (v Vector) Size() int {
	return len(v.Values)
}

func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()
	if magnitude == 0 || math.IsNaN(magnitude) || math.IsInf(magnitude, 0) {
		return v
	}
	return v.Scale(1 / magnitude)
}

func (v Vector) Equals(other Vector) bool {
	if v.Size() != other.Size() {
		return false
	}
	for i := 0; i < v.Size(); i++ {
		if v.Values[i] != other.Values[i] {
			return false
		}
	}
	return true
}
