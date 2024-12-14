package math

import (
	"fmt"
	"math"
	"strings"
)

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

func (p Polynomial) String() string {
	stringArray := make([]string, len(p.Coefficients))
	for i, coefficient := range p.Coefficients {
		stringArray[i] = fmt.Sprintf("%f*x^%d", coefficient, i)
	}
	return strings.Join(stringArray, " + ")
}
