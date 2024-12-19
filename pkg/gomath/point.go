package gomath

import (
	"math"
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
			} else if v > p2.Values[i] {
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
		return function[0](p, other)
	}
	return EuclideanDistance(p, other)
}

func (p Point) Size() int {
	return len(p.Values)
}

func (p Point) Equals(other Point) bool {
	return p.ToVector().Equals(other.ToVector())
}

func Theta(p1, p2 Point) float64 {
	retAngle := math.Atan2(p2.Y()-p1.Y(), p2.X()-p1.X())
	if retAngle < 0 {
		retAngle += 2 * math.Pi
	}
	return retAngle
}

func Average(points ...Point) *Point {
	if len(points) == 0 {
		return NewPoint(0, 0)
	}
	var retPoint Point
	for _, p := range points {
		retPoint = retPoint.Add(p)
	}
	vectorSum := retPoint.ToVector()
	retPoint = vectorSum.Scale(1.0 / float64(len(points))).ToPoint()
	return &retPoint
}

func Centroid(points ...Point) *Point {
	var centroid *Point
	n := len(points)
	if n < 3 {
		return Average(points...)
	}

	var cx, cy, areaAcc float64

	for i := 0; i < n; i++ {
		x1, y1 := points[i].X(), points[i].Y()
		x2, y2 := points[(i+1)%n].X(), points[(i+1)%n].Y()

		cross := x1*y2 - x2*y1

		cx += (x1 + x2) * cross
		cy += (y1 + y2) * cross

		areaAcc += cross
	}

	area := math.Abs(areaAcc) / 2

	cx /= (6 * area)
	cy /= (6 * area)

	centroid = NewPoint(cx, cy)
	return centroid
}
