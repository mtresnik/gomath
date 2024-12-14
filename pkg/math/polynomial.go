package math

import "math"

type Polynomial struct {
	Coefficients []float64
}

func (p Polynomial) Eval(x float64) float64 {
	result := 0.0
	if x == 0.0 {
		return p.Coefficients[0]
	}
	for i, coefficient := range p.Coefficients {
		result += coefficient * math.Pow(x, float64(i))
	}
	return result
}
