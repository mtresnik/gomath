package gomath

import (
	"sort"
)

func ConvexHull(points []Point) []Point {
	if len(points) <= 1 {
		return points
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].X() == points[j].X() {
			return points[i].Y() < points[j].Y()
		}
		return points[i].X() < points[j].X()
	})

	hull := make([]Point, 0)

	for _, p := range points {
		for len(hull) >= 2 && !isCounterClockwise(hull[len(hull)-2], hull[len(hull)-1], p) {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}

	minSize := len(hull) + 1
	for i := len(points) - 2; i >= 0; i-- {
		p := points[i]
		for len(hull) >= minSize && !isCounterClockwise(hull[len(hull)-2], hull[len(hull)-1], p) {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}

	return hull[:len(hull)-1]
}

func isCounterClockwise(a, b, c Point) bool {
	return (b.X()-a.X())*(c.Y()-a.Y()) > (b.Y()-a.Y())*(c.X()-a.X())
}
