package vector2

type ColliderRay struct {
	ray    Ray
	length float64
}

type PolygonCollider struct {
	segments []ColliderRay
}

func (p PolygonCollider) Collide(r Ray) []Vector {
	vecs := make([]Vector, 0)
	for _, segment := range p.segments {
		_, hit := r.Intersect(segment.ray)
		if hit {
			t := r.Dir.Dot(segment.ray.Dir) / segment.length
			n := r.Dir.Normal().Dot(segment.ray.Dir) / segment.length
			vecs = append(vecs, Vector{X: t, Y: n})
		}
	}
	return vecs
}
