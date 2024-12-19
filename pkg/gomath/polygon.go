package gomath

import (
	"math"
	"sort"
)

type Polygon struct {
	Points   []Point
	Segments []Segment
	bbox     BoundingBox
	centroid *Point
	area     float64
	Shape
}

func NewPolygon(points ...Point) *Polygon {
	sort.Slice(points, func(i, j int) bool {
		compared := ComparePoints(points[i], points[j])
		if compared < 0 {
			return true
		}
		return false
	})

	var segments []Segment
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]
		segments = append(segments, NewSegment(p1, p2))
	}

	return &Polygon{
		Points:   points,
		Segments: segments,
		bbox:     NewBoundingBox(PointsToSpatial(points...)...),
	}
}

func (p *Polygon) Area() float64 {
	n := len(p.Points)
	if n < 3 {
		return 0
	}

	area := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += p.Points[i].X()*p.Points[j].Y() - p.Points[j].X()*p.Points[i].Y()
	}
	return math.Abs(area) / 2
}

func (p *Polygon) GetCentroid() *Point {
	if p.centroid != nil {
		return p.centroid
	}
	p.centroid = Centroid(p.Points...)
	return p.centroid
}

func (p *Polygon) Contains(point Point) bool {
	tempPoints := append(p.Points, p.Points...)
	tempPoints = append(tempPoints, point)
	return p.Area() == NewPolygon(ConvexHull(tempPoints...)...).Area()

}
