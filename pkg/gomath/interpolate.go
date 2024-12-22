package gomath

func LerpFloats(start, end float64, t float64) float64 {
	return start + t*(end-start)
}

func Lerp(p1 Point, p2 Point, t float64) Point {
	return p2.Subtract(p1).Scale(t).AddPoint(p1)
}
