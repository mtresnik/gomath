package gomath

import "sort"

type Quad struct {
	Points []Point
}

func NewQuad(a, b, c, d Point) Quad {
	points := []Point{a, b, c, d}
	sort.Slice(points, func(i, j int) bool {
		compared := ComparePoints(points[i], points[j])
		if compared < 0 {
			return true
		}
		return false
	})
	return Quad{
		Points: points,
	}
}

func (q Quad) Area() float64 {
	/*
		0		3

		1		2
	*/
	// Construct two triangles ane see what the area is.
	firstTriangle := NewTriangle(q.Points[0], q.Points[1], q.Points[2])
	secondTriangle := NewTriangle(q.Points[0], q.Points[2], q.Points[3])
	return firstTriangle.Area() + secondTriangle.Area()
}

func (q Quad) Contains(point Point) bool {
	/*
		0		3
			p
		1		2
	*/
	// Construct 4 triangles and add the area
	firstTriangle := NewTriangle(q.Points[0], q.Points[1], point)
	secondTriangle := NewTriangle(q.Points[0], point, q.Points[2])
	thirdTriangle := NewTriangle(q.Points[0], q.Points[2], point)
	fourthTriangle := NewTriangle(q.Points[0], point, q.Points[3])
	return firstTriangle.Area()+secondTriangle.Area()+thirdTriangle.Area()+fourthTriangle.Area() == q.Area()
}
