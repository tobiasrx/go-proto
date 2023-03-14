package vector2

import (
	"math"
)

type Circle struct {
	Origin Vector
	Radius float64
}

func (c Circle) Intersect(r Ray) []Vector {
	b := r.Origin.Dot(r.Dir) - r.Dir.Dot(c.Origin)
	num1 := b * b
	num2 := r.Dir.LengthSqr() * (r.Origin.LengthSqr() + c.Origin.LengthSqr() - 2*r.Origin.Dot(c.Origin) - c.Radius*c.Radius)

	if num1-num2 < 0 {
		return []Vector{}
	}

	if num1 == num2 {
		c := -b / r.Dir.LengthSqr()

		if c < 0 || c > 1 {
			return []Vector{}
		}
		return []Vector{r.At(c)}
	}

	sqrt := math.Sqrt(num1 - num2)
	c1 := -b + sqrt
	c2 := -b - sqrt

	c1 /= r.Dir.LengthSqr()
	c2 /= r.Dir.LengthSqr()

	if (c1 < 0 || c1 > 1) && (c2 < 0 || c2 > 1) {
		return []Vector{}
	}

	if c1 < 0 || c1 > 1 {
		return []Vector{r.At(c2)}
	}

	if c2 < 0 || c2 > 1 {
		return []Vector{r.At(c1)}
	}

	return []Vector{r.At(c1), r.At(c2)}
}
