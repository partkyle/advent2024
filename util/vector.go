package util

import (
	"fmt"
)

type mathable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Vector[T mathable] struct {
	X, Y T
}

func Scalar[T mathable](s T) Vector[T] {
	return Vector[T]{X: s, Y: s}
}

func (v Vector[T]) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}

func (v Vector[T]) Add(o Vector[T]) Vector[T] {
	return Vector[T]{v.X + o.X, v.Y + o.Y}
}

func (v Vector[T]) Sub(o Vector[T]) Vector[T] {
	return Vector[T]{v.X - o.X, v.Y - o.Y}
}

func (v Vector[T]) Within(lo Vector[T], hi Vector[T]) bool {
	return lo.X <= v.X && v.X <= hi.X &&
		lo.Y <= v.Y && v.Y <= hi.Y
}

func (v Vector[T]) MulScalar(t T) Vector[T] {
	return Vector[T]{v.X * t, v.Y * t}
}

type CVec complex128

func NewCvec(x, y int) CVec {
	return CVec(complex(float64(x), float64(y)))
}

func (c CVec) X() int {
	return int(real(c))
}

func (c CVec) Y() int {
	return int(imag(c))
}
