package math

import "math"

type DistanceFunction interface {
	Eval(a Spatial, b Spatial) float64
}

type EuclideanDistance struct{}

func (f EuclideanDistance) Eval(a, b Spatial) float64 {
	return math.Sqrt(ToVector(a).Subtract(ToVector(b)).Magnitude())
}

type HaversineDistance struct {
}

func (f HaversineDistance) Eval(one, other Spatial) float64 {
	lon1, lat1 := one.X(), one.Y()
	lon2, lat2 := other.X(), other.Y()
	earthRadius := 6371000.0
	toRadians := math.Pi / 180.0
	deltaLat := (lat2 - lat1) * toRadians
	deltaLon := (lon2 - lon1) * toRadians
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) + math.Cos(lat1*toRadians)*math.Cos(lat2*toRadians)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c
}

type ManhattanDistance struct{}

func (f ManhattanDistance) Eval(one, other Spatial) float64 {
	return math.Abs(one.X()-other.X()) + math.Abs(one.Y()-other.Y())
}
