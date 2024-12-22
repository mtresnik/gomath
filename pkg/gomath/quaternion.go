package gomath

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math"
)

type Quaternion struct {
	W, X, Y, Z float64
}

func NewQuaternionFromVector(w float64, v Vector) *Quaternion {
	return NewQuaternion(w, v.X(), v.Y(), v.Z())
}

func NewQuaternion(w, x, y, z float64) *Quaternion {
	return &Quaternion{w, x, y, z}
}

func NewPureQuaternion(x, y, z float64) *Quaternion {
	return &Quaternion{0, x, y, z}
}

func NewRotationQuaternion(angle float64, axis Vector) *Quaternion {
	return NewQuaternionFromVector(math.Cos(angle/2), axis.Normalize().Scale(math.Sin(angle/2)))
}

// Slerp
// You know, like how you should eat soup.
func Slerp(q1, q2 *Quaternion, t float64) *Quaternion {
	theta := q1.Theta(q2)
	return q1.Scale(math.Sin((1-t)*theta) / math.Sin(theta)).Add(q2.Scale(math.Sin(t*theta) / math.Sin(theta)))
}

func (q *Quaternion) Normalize() *Quaternion {
	norm := q.Norm()
	if norm == 0 {
		return q
	}
	return q.Scale(1 / norm)
}

func (q *Quaternion) Norm() float64 {
	return Vector{Values: []float64{q.W, q.X, q.Y, q.Z}}.Magnitude()
}

func (q *Quaternion) Scale(s float64) *Quaternion {
	return NewQuaternion(q.W*s, q.X*s, q.Y*s, q.Z*s)
}

func (q *Quaternion) Add(other *Quaternion) *Quaternion {
	return NewQuaternion(q.W+other.W, q.X+other.X, q.Y+other.Y, q.Z+other.Z)
}

func (q *Quaternion) Subtract(other *Quaternion) *Quaternion {
	return NewQuaternion(q.W-other.W, q.X-other.X, q.Y-other.Y, q.Z-other.Z)
}

func (q *Quaternion) Conjugate() *Quaternion {
	return NewQuaternion(q.W, -q.X, -q.Y, -q.Z)
}

func (q *Quaternion) Inverse() *Quaternion {
	return q.Conjugate().Scale(1 / math.Pow(q.Norm(), 2))
}

func (q *Quaternion) Multiply(other *Quaternion) *Quaternion {
	w1 := q.W
	x1 := q.X
	y1 := q.Y
	z1 := q.Z

	w2 := other.W
	x2 := other.X
	y2 := other.Y
	z2 := other.Z

	w3 := w1*w2 - x1*x2 - y1*y2 - z1*z2
	x3 := w1*x2 + x1*w2 + y1*z2 - z1*y2
	y3 := w1*y2 - x1*z2 + y1*w2 + z1*x2
	z3 := w1*z2 + x1*y2 - y1*x2 + z1*w2
	return NewQuaternion(w3, x3, y3, z3)
}

func (q *Quaternion) Dot(other *Quaternion) float64 {
	return q.W*other.W + q.X*other.X + q.Y*other.Y + q.Z*other.Z
}

func (q *Quaternion) Theta(other *Quaternion) float64 {
	return math.Acos(q.Dot(other) / (q.Norm() * other.Norm()))
}

func (q *Quaternion) Slice() []float64 {
	return []float64{q.W, q.X, q.Y, q.Z}
}

func (q *Quaternion) String() string {
	return goutils.Float64ArrayToString(q.Slice())
}

func (q *Quaternion) Complex() complex128 {
	return complex(q.X, q.Y)
}
