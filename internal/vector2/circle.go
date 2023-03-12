package vector2

import (
	"math"
)

type Circle struct {
	Origin Vector
	Radius float64
}

func (c Circle) Intersect(r Ray) []Vector {
	b := 2 * r.Origin.Subtract(c.Origin).Dot(r.Dir)
	num1 := b * b
	num2 := 4 * r.Dir.LengthSqr() * (r.Origin.Subtract(c.Origin).LengthSqr() - c.Radius*c.Radius)

	if num1-num2 < 0 {
		return []Vector{}
	}

	if num1 == num2 {
		return []Vector{r.At(-b)}
	}
	sqrt := math.Sqrt(num1 - num2)
	sqrt /= 2 * r.Dir.LengthSqr()
	c1 := -b + sqrt
	c2 := -b - sqrt
	return []Vector{r.At(c1), r.At(c2)}
}
