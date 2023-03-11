package vector2

type Ray struct {
	Origin Vector
	Dir    Vector
}

func (r Ray) Intersect(r2 Ray) (Vector, bool) {
	dot := r.Dir.Dot(r2.Dir)
	denom := r.Dir.LengthSqr()*r2.Dir.LengthSqr() - dot*dot

	if denom == 0 {
		return Vector{}, false
	}

	n1 := r.Origin.Dot(r2.Dir) - r2.Origin.Dot(r2.Dir)
	n1 *= dot

	n2 := r.Origin.Dot(r.Dir) - r2.Origin.Dot(r.Dir)
	n2 *= r2.Dir.LengthSqr()

	num := n1 - n2
	c := num / denom
	return r.Origin.Add(r.Dir.Multiply(c)), true
}
