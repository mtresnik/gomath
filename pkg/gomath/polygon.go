package gomath

import (
	"math"
	"sort"
)

type Polygon struct {
	Points []Point
	bbox   BoundingBox
	Shape
}

func NewPolygon(points ...Point) *Polygon {
	sort.Slice(points, func(i, j int) bool {
		return ComparePoints(points[i], points[j]) < 0
	})
	return &Polygon{
		Points: points,
		bbox:   NewBoundingBox(PointsToSpatial(points...)...),
	}
}

func (a *Polygon) GetPoints() []Point {
	return a.Points
}

func (a *Polygon) Area() float64 {
	return 0
}

func (a *Polygon) Contains(point Point, distanceFunction ...DistanceFunction) bool {
	if !a.bbox.Contains(point) {
		return false
	}
	if len(a.Points) < 3 {
		return false
	}
	inside := false
	n := len(a.Points)

	j := n - 1
	for i := 0; i < n; {
		vertexI := a.Points[i]
		vertexJ := a.Points[j]
		if isPointOnLine(point, vertexI, vertexJ, distanceFunction...) {
			return true
		}
		if ((vertexI.Y() > point.Y()) != (vertexJ.Y() > point.Y())) &&
			(point.X() < (vertexJ.X()-vertexI.X())*(point.Y()-vertexI.Y())/
				(vertexJ.Y()-vertexI.Y())+vertexI.X()) {
			inside = !inside
		}
		j = i
		i++
	}
	return inside
}

func isPointOnLine(point Point, vertexI, vertexJ Point, pDistanceFunction ...DistanceFunction) bool {
	var distanceFunction DistanceFunction = EuclideanDistance{}
	if len(pDistanceFunction) > 0 {
		distanceFunction = pDistanceFunction[0]
	}
	d1 := distanceFunction.Eval(point, vertexI)
	d2 := distanceFunction.Eval(point, vertexJ)
	lineLength := distanceFunction.Eval(vertexI, vertexJ)

	epsilon := 0.0001

	return math.Abs(d1+d2-lineLength) < epsilon
}
