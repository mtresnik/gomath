package math

import (
	"math"
	"testing"
)

func TestPoint_String(t *testing.T) {
	p1 := Point{Values: []float64{math.Pi, math.Pi, math.Pi, math.Pi, math.Pi}}
	println(p1.String())
}

func TestPoint_Subtract(t *testing.T) {
	p1 := Point{Values: []float64{0, 1.0, 2.0, 3.0, 4.0}}
	p2 := Point{Values: []float64{5, 5, 5, 5}}
	v1 := p1.Subtract(p2)
	println(v1.String())
}
