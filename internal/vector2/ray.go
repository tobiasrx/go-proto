package vector2

type Ray struct {
	Origin Vector
	Dir    Vector
}

func (r Ray) At(c float64) Vector {
	return r.Origin.Add(r.Dir.Multiply(c))
}

func (r Ray) Intersect(r2 Ray) (Vector, bool) {
	dot := r.Dir.Dot(r2.Dir)
	denom := r.Dir.LengthSqr()*r2.Dir.LengthSqr() - dot*dot

	if denom == 0 {
		return Vector{}, false
	}

	n1 := r.Origin.Dot(r2.Dir) - r2.Origin.Dot(r2.Dir)
	n2 := r.Origin.Dot(r.Dir) - r2.Origin.Dot(r.Dir)

	num := n1*dot - n2*r2.Dir.LengthSqr()
	c := num / denom

	if c < 0 || c > 1 {
		return Vector{}, false
	}

	c2 := n1 + c*dot
	c2 /= r2.Dir.LengthSqr()

	if c2 < 0 || c2 > 1 {
		return Vector{}, false
	}

	return r.At(c), true
}
