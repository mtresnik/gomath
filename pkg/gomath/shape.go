package gomath

type Shape interface {
	GetPoints() []Point
	Area() float64
	Contains(point Point, distanceFunction ...DistanceFunction) bool
}

func HashShape(shape Shape) int64 {
	points := shape.GetPoints()
	var hash int64 = 0
	for _, point := range points {
		hash = hash*31 + int64(point.X())*17 + int64(point.Y())*19
	}
	return hash
}
