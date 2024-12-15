package gomath

type Segment interface {
	From() Spatial
	To() Spatial
}

func String(segment Segment) string {
	return "[" + SpatialString(segment.From()) + " -> " + SpatialString(segment.To()) + "]"
}

func Length(segment Segment, distanceFunction ...DistanceFunction) float64 {
	return ToPoint(segment.From()).DistanceTo(ToPoint(segment.To()), distanceFunction...)
}

type LineSegment struct {
	from Spatial
	to   Spatial
}

func (a LineSegment) From() Spatial {
	return a.from
}

func (a LineSegment) To() Spatial {
	return a.to
}
