package physics

import (
	"example.com/proto/internal/shapes"
	"example.com/proto/internal/vector2"
	"github.com/gonutz/prototype/draw"
)

type Body struct {
	Pos    vector2.Vector
	Vel    vector2.Vector
	Mass   float64
	Radius float64
}

func (b *Body) Step(window draw.Window) {
	b.Pos = b.Pos.Add(b.Vel)

	screen := vector2.FromPoint(window.Size())

	mirrorLeft := b.Pos.X-b.Radius < 0 && b.Vel.X < 0
	mirrorRight := b.Pos.X+b.Radius > screen.X && b.Vel.X > 0

	if mirrorLeft || mirrorRight {
		b.Vel.X = -b.Vel.X
	}

	mirrorTop := b.Pos.Y-b.Radius < 0 && b.Vel.Y < 0
	mirrorBottom := b.Pos.Y+b.Radius > screen.Y && b.Vel.Y > 0

	if mirrorTop || mirrorBottom {
		b.Vel.Y = -b.Vel.Y
	}
}

func (b *Body) Collide(b2 *Body) {
	center := b2.Pos.Subtract(b.Pos).Normalize()
	ortho := center.Normal()

	if b.Pos.Subtract(b2.Pos).LengthSqr() < (b.Radius+b2.Radius)*(b.Radius+b2.Radius) {
		v1, v2 := center.Dot(b.Vel), center.Dot(b2.Vel)
		st := 2 * (b.Mass*v1 + b2.Mass*v2) / (b.Mass + b2.Mass)
		v1p := st - v1
		v2p := st - v2
		v1o := ortho.Dot(b.Vel)
		v2o := ortho.Dot(b2.Vel)
		b.Vel = center.Multiply(v1p).Add(ortho.Multiply(v1o))
		b2.Vel = center.Multiply(v2p).Add(ortho.Multiply(v2o))
	}
}

func (b Body) Draw(window draw.Window) {
	circle := vector2.Circle{
		Origin: b.Pos,
		Radius: b.Radius,
	}
	shapes.DrawCircle(window, circle, draw.LightPurple)
	shapes.DrawArrow(window, b.Pos, b.Pos.Add(b.Vel.Multiply(b.Radius).Divide(2)), draw.White)
}
