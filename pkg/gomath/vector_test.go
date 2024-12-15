package gomath

import "testing"

func TestVector_String(t *testing.T) {
	v1 := Vector{Values: []float64{1, 2, 3}}
	println(v1.String())
}

func TestVector_Add(t *testing.T) {
	v1 := Vector{Values: []float64{1, 2, 3}}
	v2 := Vector{Values: []float64{2, 3, 4}}
	v3 := v1.Add(v2)
	println(v3.String())
}

func TestVector_Subtract(t *testing.T) {
	v1 := Vector{Values: []float64{600, 2, 3}}
	v2 := Vector{Values: []float64{2, 10, 400}}
	v3 := v1.Subtract(v2)
	println(v3.String())
}

func TestVector_Normalize(t *testing.T) {
	v1 := Vector{Values: []float64{1, 2, 3}}
	norm := v1.Normalize()
	println(norm.String())
	mag := norm.Magnitude()
	println(mag)
}
