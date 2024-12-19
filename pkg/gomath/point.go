package gomath

import (
	"math/rand"
	"time"
)

type Point struct {
	Values []float64
}

func NewPoint(values ...float64) *Point {
	return &Point{Values: values}
}

func RandomPoints(num, dim int, pScalar ...float64) []Point {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	scalar := 1.0
	if len(pScalar) > 0 {
		scalar = pScalar[0]
	}
	retSlice := make([]Point, num)
	for i := range retSlice {
		retSlice[i] = Point{Values: []float64{}}
		for j := 0; j < dim; j++ {
			retSlice[i].Values = append(retSlice[i].Values, random.Float64()*scalar)
		}
	}
	return retSlice
}

func ComparePoints(p1, p2 Point) int {
	for i, v := range p1.Values {
		if v != p2.Values[i] {
			if v < p2.Values[i] {
				return -1
			} else {
				return 1
			}
		}
	}
	return 0
}

func PointsToSpatial(points ...Point) []Spatial {
	retSlice := make([]Spatial, len(points))
	for i, p := range points {
		retSlice[i] = p
	}
	return retSlice
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
