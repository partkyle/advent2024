package util

type Vector struct {
	X, Y int
}

func (v Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y}
}

func (v Vector) Sub(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y}
}

func (v Vector) Within(lo Vector, hi Vector) bool {
	return lo.X <= v.X && v.X < hi.X &&
		lo.Y <= v.Y && v.Y < hi.Y
}

func (v Vector) MulScalar(t int) Vector {
	return Vector{v.X * t, v.Y * t}
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
