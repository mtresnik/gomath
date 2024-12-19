package gomath

import "math"

type DistanceFunction func(one, other Spatial) float64

var EuclideanDistance = func(one, other Spatial) float64 {
	return math.Sqrt(ToVector(one).Subtract(ToVector(other)).Magnitude())
}

var HaversineDistance = func(one, other Spatial) float64 {
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

var ManhattanDistance = func(one, other Spatial) float64 {
	retValue := 0.0
	for i := 0; i < max(len(one.GetValues()), len(other.GetValues())); i++ {
		left, right := 0.0, 0.0
		if i < len(one.GetValues()) {
			left = one.GetValues()[i]
		}
		if i < len(other.GetValues()) {
			right = other.GetValues()[i]
		}
		delta := math.Abs(left - right)
		retValue += delta
	}
	return retValue
}
