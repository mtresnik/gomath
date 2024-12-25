package gomath

type Triangle struct {
	points []Point
	area   float64
}

func (t *Triangle) Contains(point Point, distanceFunction ...DistanceFunction) bool {
	area1 := 0.5 * ((t.points[0].X() * (t.points[1].Y() - point.Y())) +
		(t.points[1].X() * (point.Y() - t.points[0].Y())) +
		(point.X() * (t.points[0].Y() - t.points[1].Y())))

	area2 := 0.5 * ((t.points[1].X() * (t.points[2].Y() - point.Y())) +
		(t.points[2].X() * (point.Y() - t.points[1].Y())) +
		(point.X() * (t.points[1].Y() - t.points[2].Y())))

	area3 := 0.5 * ((t.points[2].X() * (t.points[0].Y() - point.Y())) +
		(t.points[0].X() * (point.Y() - t.points[2].Y())) +
		(point.X() * (t.points[2].Y() - t.points[0].Y())))

	totalArea := t.Area()
	return area1+area2+area3 == totalArea
}

func NewTriangle(points ...Point) *Triangle {
	if len(points) < 3 {
		panic("Invalid number of points")
	}
	return &Triangle{points: points[:3]}
}

func (t *Triangle) GetPoints() []Point {
	return t.points
}

func (t *Triangle) Area() float64 {
	if t.area > 0 {
		return t.area
	}

	t.area = 0.5 * ((t.points[0].X() * (t.points[1].Y() - t.points[2].Y())) +
		(t.points[1].X() * (t.points[2].Y() - t.points[0].Y())) +
		(t.points[2].X() * (t.points[0].Y() - t.points[1].Y())))

	return t.area
}
