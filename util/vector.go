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
