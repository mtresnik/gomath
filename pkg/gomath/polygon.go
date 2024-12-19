package gomath

import (
	"math"
	"sort"
)

type Polygon struct {
	Points   []Point // Vertices of the polygon, ordered counterclockwise
	Edges    []Edge  // Edges of the polygon
	bbox     BoundingBox
	centroid *Point
	Shape
}

// NewPolygon creates a new polygon from a slice of points.
// Assumes the points form a valid boundary (e.g., convex hull).
func NewPolygon(points ...Point) *Polygon {
	// Ensure points are sorted (by x, then y) for consistency
	sort.Slice(points, func(i, j int) bool {
		if points[i].X() == points[j].X() {
			return points[i].Y() < points[j].Y()
		}
		return points[i].X() < points[j].X()
	})

	// Construct edges from the points in counterclockwise order
	var edges []Edge
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)] // Wrap around to the first vertex
		edges = append(edges, Edge{P1: p1, P2: p2})
	}

	return &Polygon{
		Points: points,
		Edges:  edges,
		bbox:   NewBoundingBox(PointsToSpatial(points...)...),
	}
}

// Area calculates the area of the polygon using the shoelace formula.
func (p *Polygon) Area() float64 {
	n := len(p.Points)
	if n < 3 {
		return 0 // Not a valid polygon
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
	n := len(p.Points)
	if n < 3 {
		// Degenerate case: If fewer than 3 points, no valid polygon exists
		return NewPoint(0, 0)
	}

	// Initialize accumulators for centroid coordinates and the area
	var cx, cy, areaAcc float64

	for i := 0; i < n; i++ {
		// Current point and the next point (wrapping around for the last point)
		x1, y1 := p.Points[i].X(), p.Points[i].Y()
		x2, y2 := p.Points[(i+1)%n].X(), p.Points[(i+1)%n].Y()

		// Cross product for the current edge
		cross := x1*y2 - x2*y1

		// Accumulate the centroid coordinates
		cx += (x1 + x2) * cross
		cy += (y1 + y2) * cross

		// Accumulate the area (using the shoelace formula)
		areaAcc += cross
	}

	// Calculate the actual area of the polygon
	area := math.Abs(areaAcc) / 2

	// Calculate the centroid coordinates
	cx /= (6 * area)
	cy /= (6 * area)

	p.centroid = NewPoint(cx, cy)
	return p.centroid
}

// Contains checks if a point is inside the polygon using edge intersections.
func (p *Polygon) Contains(point Point) bool {
	tempPoints := append(p.Points, p.Points...)
	tempPoints = append(tempPoints, point)
	return p.Area() == NewPolygon(ConvexHull(tempPoints)...).Area()

}

func Theta(p1, p2 Point) float64 {
	return math.Atan2(p2.Y()-p1.Y(), p2.X()-p1.X())
}
