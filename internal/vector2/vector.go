package vector2

import "math"

type Vector struct {
	X float64
	Y float64
}

func FromPoint(x, y int) Vector {
	return Vector{X: float64(x), Y: float64(y)}
}

func FromRotation(alpha float64) Vector {
	return Vector{X: 1, Y: 0}.RotateAndNormalize(alpha)
}

func (v Vector) Point() (int, int) {
	return int(math.Round(v.X)), int(math.Round(v.Y))
}

func (v Vector) Add(other Vector) Vector {
	return Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector) Subtract(other Vector) Vector {
	return Vector{X: v.X - other.X, Y: v.Y - other.Y}
}

func (v Vector) Divide(l float64) Vector {
	if l == 0.0 {
		return v
	}
	return Vector{X: v.X / l, Y: v.Y / l}
}

func (v Vector) Multiply(l float64) Vector {
	return Vector{X: v.X * l, Y: v.Y * l}
}

func (v Vector) Dot(other Vector) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.LengthSqr())
}

func (v Vector) LengthSqr() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector) Normalize() Vector {
	length := v.Length()
	return v.Divide(length)
}

func (v Vector) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vector) Rotate(alpha float64) Vector {
	length := v.Length()
	return v.RotateAndNormalize(alpha).Multiply(length)
}

func (v Vector) RotateAndNormalize(alpha float64) Vector {
	angle := v.Angle() + alpha
	return Vector{X: math.Cos(angle), Y: math.Sin(angle)}
}
