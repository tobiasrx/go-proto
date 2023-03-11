package vector2

type BoundingBox struct {
	rayTop    Ray
	rayRight  Ray
	rayBottom Ray
	rayLeft   Ray
}

func Init(topLeft, bottomRight Vector) BoundingBox {
	diff := bottomRight.Subtract(topLeft)
	rayTop := Ray{Origin: topLeft, Dir: Vector{X: diff.X, Y: 0}}
	rayRight := Ray{Origin: bottomRight, Dir: Vector{X: 0, Y: -diff.Y}}
	rayBottom := Ray{Origin: bottomRight, Dir: Vector{X: -diff.X, Y: 0}}
	rayLeft := Ray{Origin: topLeft, Dir: Vector{X: 0, Y: diff.Y}}
	return BoundingBox{rayTop, rayRight, rayBottom, rayLeft}
}

func (b BoundingBox) Rays() [4]Ray {
	return [4]Ray{b.rayTop, b.rayRight, b.rayBottom, b.rayLeft}
}

func (b BoundingBox) Intersect(r Ray) []Vector {
	vecs := make([]Vector, 0)
	for _, ray := range b.Rays() {
		v, ok := r.Intersect(ray)
		if ok {
			vecs = append(vecs, v)
		}
	}
	return vecs
}
