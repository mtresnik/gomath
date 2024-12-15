package mrmath

type Shape interface {
	GetPoints() []Point
	Area() float64
	Contains(point Point, distanceFunction ...DistanceFunction) bool
}
