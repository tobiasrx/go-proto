package main

import (
	"fmt"
	"math"

	"example.com/proto/internal/physics"
	"example.com/proto/internal/vector2"
	"github.com/gonutz/prototype/draw"
)

func main() {
	draw.RunWindow("Title", 640, 480, update)
}

func radToDegrees(r float64) float64 {
	return r / math.Pi * 180
}

var body = physics.Body{
	Pos:    vector2.FromPoint(50, 101),
	Vel:    vector2.FromPoint(12, 5),
	Mass:   1.0,
	Radius: 50,
}

var bodies []*physics.Body = make([]*physics.Body, 0)

func update(window draw.Window) {
	for _, cl := range window.Clicks() {
		if cl.Button == draw.LeftButton {
			b := physics.Body{
				Pos:    vector2.FromPoint(cl.X, cl.Y),
				Vel:    vector2.FromPoint(10, 0),
				Mass:   1.0,
				Radius: 50,
			}
			bodies = append(bodies, &b)
		}
	}
	vel := 0.0
	mass := 0.0
	imp := 0.0
	for i, b := range bodies {
		b.Draw(window)
		b.Step(window)
		for j := i + 1; j < len(bodies); j++ {
			b.Collide(bodies[j])
		}
		vel += b.Vel.Length()
		mass += b.Mass
		imp += vel * mass
	}

	window.DrawScaledText(fmt.Sprintf("Vel: %v", vel), 0, 0, 1.6, draw.RGB(0.2, 0.5, 0.3))
}
