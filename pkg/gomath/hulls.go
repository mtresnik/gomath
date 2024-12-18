package gomath

import "sort"

type ConvexHull struct {
	points []Point
	bbox   BoundingBox
}

func (a *ConvexHull) Points() []Point {
	return a.points
}

func NewConvexHull(pPoints ...Point) *Polygon {
	if len(pPoints) <= 1 {
		return NewPolygon(pPoints...)
	}
	points := make([]Point, len(pPoints))
	copy(points, pPoints)
	sort.Slice(points, func(i, j int) bool {
		return ComparePoints(points[i], points[j]) < 0
	})
	hull := make([]Point, 0)
	for _, point := range points {
		for len(hull) >= 2 && !isCounterClockwise(hull[len(hull)-2], hull[len(hull)-1], point) {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, point)
	}
	return NewPolygon(hull...)
}

func isCounterClockwise(a, b, c Point) bool {
	return ((b.X() - a.X()) * (c.Y() - a.Y())) > ((b.Y() - a.Y()) * (c.X() - a.X()))
}
