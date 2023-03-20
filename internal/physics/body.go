package physics

import "example.com/proto/internal/vector2"

type Body struct {
	Pos  vector2.Vector
	Vel  vector2.Vector
	Mass float64
}

func (b *Body) Step() {
	b.Pos = b.Pos.Add(b.Vel)
}
