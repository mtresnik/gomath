package gomath

type BoundingBox struct {
	MinX, MinY, MaxX, MaxY float64
}

func (r BoundingBox) GetPoints() []Point {
	return []Point{
		{Values: []float64{r.MinX, r.MinY}},
		{Values: []float64{r.MinX, r.MaxY}},
		{Values: []float64{r.MaxX, r.MaxY}},
		{Values: []float64{r.MaxX, r.MinY}},
	}
}

func (r BoundingBox) Area() float64 {
	return (r.MaxX - r.MinX) * (r.MaxY - r.MinY)
}

func (r BoundingBox) Contains(point Spatial, distanceFunction ...DistanceFunction) bool {
	return point.X() >= r.MinX && point.X() <= r.MaxX && point.Y() >= r.MinY && point.Y() <= r.MaxY
}

func NewBoundingBox(points ...Spatial) BoundingBox {
	if len(points) == 0 {
		return BoundingBox{0, 0, 0, 0}
	}
	minX := points[0].X()
	minY := points[0].Y()
	maxX := points[0].X()
	maxY := points[0].Y()
	for _, p := range points {
		if p.X() < minX {
			minX = p.X()
		}
		if p.Y() < minY {
			minY = p.Y()
		}
		if p.X() > maxX {
			maxX = p.X()
		}
		if p.Y() > maxY {
			maxY = p.Y()
		}
	}
	return BoundingBox{minX, minY, maxX, maxY}
}
