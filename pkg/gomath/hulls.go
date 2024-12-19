package gomath

import (
	"sort"
)

// Edge represents a line segment between two points
type Edge struct {
	P1, P2 Point
}

// ConvexHull calculates the convex hull of the given points.
func ConvexHull(points []Point) []Point {
	if len(points) <= 1 {
		return points
	}

	// Sort the points by x, then by y.
	sort.Slice(points, func(i, j int) bool {
		if points[i].X() == points[j].X() {
			return points[i].Y() < points[j].Y()
		}
		return points[i].X() < points[j].X()
	})

	hull := make([]Point, 0)

	// Lower hull
	for _, p := range points {
		for len(hull) >= 2 && !isCounterClockwise(hull[len(hull)-2], hull[len(hull)-1], p) {
			hull = hull[:len(hull)-1] // Remove the last point
		}
		hull = append(hull, p)
	}

	// Upper hull
	minSize := len(hull) + 1
	for i := len(points) - 2; i >= 0; i-- {
		p := points[i]
		for len(hull) >= minSize && !isCounterClockwise(hull[len(hull)-2], hull[len(hull)-1], p) {
			hull = hull[:len(hull)-1] // Remove the last point
		}
		hull = append(hull, p)
	}

	// Remove the last point since it's the same as the first
	return hull[:len(hull)-1]
}

// isCounterClockwise checks if the sequence of three points makes a counterclockwise turn.
func isCounterClockwise(a, b, c Point) bool {
	// Calculate the cross product to determine counterclockwise order
	return (b.X()-a.X())*(c.Y()-a.Y()) > (b.Y()-a.Y())*(c.X()-a.X())
}
