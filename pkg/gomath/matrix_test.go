package gomath

import (
	"math/rand"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	matrix := NewMatrix(3, 3, func(x, y int) float64 { return rand.Float64() })
	println(matrix.String())
}

func TestMatrix_Inverse(t *testing.T) {
	matrix := NewMatrixFromValues([][]float64{{4.0, 7.0}, {2.0, 6.0}})
	println(matrix.String())
	println(matrix.Determinant())
	println(matrix.Inverse().String())
	println(matrix.Multiply(matrix.Inverse()).String())

}

func TestMatrix_Multiply(t *testing.T) {
	// Input matrices
	matrixA := NewMatrixFromValues([][]float64{
		{1, 2},
		{3, 4},
	})

	matrixB := NewMatrixFromValues([][]float64{
		{2, 0},
		{1, 2},
	})

	// Expected result of multiplication
	expected := NewMatrixFromValues([][]float64{
		{4, 4},
		{10, 8},
	})

	// Perform matrix multiplication
	result := matrixA.Multiply(matrixB)

	// Verify the result matches the expected matrix
	if !MatrixEquals(result, expected) {
		t.Errorf("Matrix multiplication failed.\nExpected:\n%v\nGot:\n%v", expected.String(), result.String())
	}
}

func TestMatrix_RemoveRowCol(t *testing.T) {
	matrix := NewMatrixFromValues([][]float64{
		{4, 7, 2},
		{3, 6, 1},
		{8, 5, 9},
	})

	expected := NewMatrixFromValues([][]float64{
		{3, 6},
		{8, 5},
	})

	result := matrix.RemoveRowCol(0, 2)
	if !MatrixEquals(result, expected) {
		t.Errorf("Unexpected RemoveRowCol: Got %v, Expected %v", result, expected)
	}
}
