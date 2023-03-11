package vector2

type BoundingBox struct {
	rayTop    Ray
	rayRight  Ray
	rayBottom Ray
	rayLeft   Ray
}

func Init(topLeft, bottomRight Vector) BoundingBox {
	rayTop := Ray{Origin: topLeft, Dir: Vector{X: 1, Y: 0}}
	rayRight := Ray{Origin: bottomRight, Dir: Vector{X: 0, Y: -1}}
	rayBottom := Ray{Origin: bottomRight, Dir: Vector{X: -1, Y: 0}}
	rayLeft := Ray{Origin: topLeft, Dir: Vector{X: 0, Y: 1}}
	return BoundingBox{rayTop, rayRight, rayBottom, rayLeft}
}

func (b BoundingBox) Rays() [4]Ray {
	return [4]Ray{b.rayTop, b.rayRight, b.rayBottom, b.rayLeft}
}

func (b BoundingBox) Intersect(r Ray) (Vector, bool) {
	isOk := false
	a := 0.0
	for _, ray := range b.Rays() {
		c, ok := r.Intersect(ray)
		if ok && c >= 0 && (!isOk || c < a) {
			a = c
			isOk = true
		}
	}
	if isOk {
		return r.At(a), true
	}

	return Vector{}, false
}
