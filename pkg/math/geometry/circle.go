package geometry

import "math"

type Circle struct {
	Center Point
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Contains(point Point, distanceFunction ...DistanceFunction) bool {
	return c.Center.DistanceTo(point, distanceFunction...) <= c.Radius
}

func (c Circle) GetPoints() []Point {
	retPoints := make([]Point, 100)
	for i := 0; i < len(retPoints); i++ {
		theta := 2 * math.Pi * float64(i) / float64(len(retPoints))
		x := math.Cos(theta)*c.Radius + c.Center.X()
		y := math.Sin(theta)*c.Radius + c.Center.Y()
		retPoints[i] = Point{Values: []float64{x, y}}
	}
	return retPoints
}
