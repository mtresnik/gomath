package gomath

import (
	"fmt"
	"github.com/mtresnik/goutils/pkg/goutils"
	"math"
)

type Matrix struct {
	Values [][]float64
	Height int
	Width  int
}

func EulerMatrix(angle float64) *Matrix {
	return NewMatrix(2, 2, func(row, col int) float64 {
		if row == col {
			return math.Cos(angle)
		}
		if row == 1-col {
			return -math.Sin(angle)
		}
		return 0.0
	})
}

func NewIdentityMatrix(size int) *Matrix {
	return NewMatrix(size, size, func(row, col int) float64 {
		if row == col {
			return 1.0
		}
		return 0.0
	})
}

func NewMatrixFromValues(values [][]float64) *Matrix {
	return &Matrix{
		Values: values,
		Height: len(values),
		Width:  len(values[0]),
	}
}

func NewMatrix(height, width int, optionalInitializer ...func(row, col int) float64) *Matrix {
	initializer := func(row, col int) float64 { return 0.0 }
	if len(optionalInitializer) > 0 {
		initializer = optionalInitializer[0]
	}
	values := make([][]float64, height)
	for row := 0; row < height; row++ {
		values[row] = make([]float64, width)
		for col := 0; col < width; col++ {
			values[row][col] = initializer(row, col)
		}
	}
	return &Matrix{
		Values: values,
		Height: height,
		Width:  width,
	}
}

func MatrixEquals(a, b *Matrix) bool {
	if a.Height != b.Height || a.Width != b.Width {
		return false
	}
	for i := 0; i < a.Height; i++ {
		for j := 0; j < a.Width; j++ {
			if math.Abs(a.Values[i][j]-b.Values[i][j]) > 1e-9 {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) Map(f func(row, col int) float64) *Matrix {
	return NewMatrix(m.Height, m.Width, func(row, col int) float64 {
		return f(row, col)
	})
}

func (m *Matrix) Add(other *Matrix) *Matrix {
	if m.Width != other.Width || m.Height != other.Height {
		panic("Matrices must be the same size to add")
	}
	return m.Map(func(row, col int) float64 {
		return m.Values[row][col] + other.Values[row][col]
	})
}

func (m *Matrix) Subtract(other *Matrix) *Matrix {
	if m.Width != other.Width || m.Height != other.Height {
		panic("Matrices must be the same size to subtract")
	}
	return m.Map(func(row, col int) float64 {
		return m.Values[row][col] - other.Values[row][col]
	})
}

func (m *Matrix) Scale(scalar float64) *Matrix {
	return m.Map(func(row, col int) float64 {
		return m.Values[row][col] * scalar
	})
}

func (m *Matrix) Transpose() *Matrix {
	ret := NewMatrix(m.Width, m.Height)
	for row := 0; row < m.Height; row++ {
		for col := 0; col < m.Width; col++ {
			ret.Values[col][row] = m.Values[row][col]
		}
	}
	return ret
}

func (m *Matrix) String() string {
	retString := "["
	for row := 0; row < m.Height; row++ {
		for col := 0; col < m.Width; col++ {
			retString += fmt.Sprintf("%.4f", m.Values[row][col])
			if col < m.Width-1 {
				retString += ","
			}
		}
		if row < m.Height-1 {
			retString += "\n"
		}
	}
	retString += "]"
	return retString
}

func (m *Matrix) Copy() *Matrix {
	return NewMatrix(m.Height, m.Width, func(row, col int) float64 {
		return m.Values[row][col]
	})
}

func (m *Matrix) Multiply(other *Matrix) *Matrix {
	if m.Width != other.Height {
		panic("Matrices must be the same size to multiply")
	}
	ret := NewMatrix(m.Height, other.Width)
	for row := 0; row < m.Height; row++ {
		for col := 0; col < other.Width; col++ {
			for i := 0; i < m.Width; i++ {
				ret.Values[row][col] += m.Values[row][i] * other.Values[i][col]
			}
		}
	}
	return ret
}

func (m *Matrix) Hadamard(other *Matrix) *Matrix {
	if m.Width != other.Width || m.Height != other.Height {
		panic("Matrices must be the same size to multiply")
	}
	return m.Map(func(row, col int) float64 {
		return m.Values[row][col] * other.Values[row][col]
	})
}

func (m *Matrix) Size() int {
	return m.Height * m.Width
}

func (m *Matrix) IsSquare() bool {
	return m.Height == m.Width
}

func (m *Matrix) RemoveRowCol(row, col int) *Matrix {
	rows := goutils.RangeOfInts(0, m.Height)
	cols := goutils.RangeOfInts(0, m.Width)
	rows = append(rows[:row], rows[row+1:]...)
	cols = append(cols[:col], cols[col+1:]...)

	return NewMatrix(m.Height-1, m.Width-1, func(r, c int) float64 {
		return m.Values[rows[r]][cols[c]]
	})
}

func (m *Matrix) Determinant() float64 {
	if !m.IsSquare() {
		panic("Matrix must be square to find Determinant")
	}
	if m.Height == 1 {
		return m.Values[0][0]
	}
	if m.Height == 2 {
		return m.Values[0][0]*m.Values[1][1] - m.Values[0][1]*m.Values[1][0]
	}

	det := 0.0
	sign := 1.0
	for col := 0; col < m.Width; col++ {
		tempMatrix := m.RemoveRowCol(0, col)
		det += sign * m.Values[0][col] * tempMatrix.Determinant()
		sign = -sign
	}
	return det
}

func (m *Matrix) Cofactor() *Matrix {
	if !m.IsSquare() {
		panic("Matrix must be square to Cofactor")
	}
	return m.Map(func(row, col int) float64 {
		sign := math.Pow(-1, float64(row+col+2))
		tempMatrix := m.RemoveRowCol(row, col)
		value := sign * tempMatrix.Determinant()
		return value
	})
}

func (m *Matrix) Adjoint() *Matrix {
	return m.Cofactor().Transpose()
}

func (m *Matrix) Inverse() *Matrix {
	if !m.IsSquare() {
		panic("Matrix must be square to invert")
	}
	det := m.Determinant()
	if det == 0 {
		panic("Matrix is not invertible")
	}
	return m.Adjoint().Scale(1.0 / det)
}
