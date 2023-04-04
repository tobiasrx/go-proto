package main

import (
	"fmt"
	"math"

	"example.com/proto/internal/physics"
	"example.com/proto/internal/vector2"
	"github.com/gonutz/prototype/draw"
)

var bodies []*physics.Body = make([]*physics.Body, 0)

func main() {
	b := physics.Body{
		Pos: vector2.Vector{X: 416.94652450700806, Y: 333.0061796809844},
		Vel: vector2.Vector{X: 3, Y: 4},
		// Vel:    vector2.Vector{X: 7.600544050525671, Y: 6.498594473885775},
		Mass:   1.0,
		Radius: 50,
	}
	bodies = append(bodies, &b)
	b2 := physics.Body{
		Pos: vector2.Vector{X: 448.4932085556092, Y: 426.954579286334},
		// Vel:    vector2.Vector{X: -7.600544050525671, Y: 6.498594473885775},
		Vel:    vector2.Null(),
		Mass:   1.0,
		Radius: 50,
	}
	bodies = append(bodies, &b2)
	draw.RunWindow("Title", 640, 480, update)
}

func radToDegrees(r float64) float64 {
	return r / math.Pi * 180
}

func update(window draw.Window) {

	imp := vector2.Null()
	egy := 0.0

	for i, b := range bodies {
		b.Draw(window)
		// b.Step(window)

		for j := i + 1; j < len(bodies); j++ {
			b.Collide(bodies[j], window)
		}
		imp = imp.Add(b.Vel.Multiply(b.Mass))
		egy += 0.5 * b.Mass * b.Vel.LengthSqr()
	}
	window.DrawScaledText(fmt.Sprintf("Imp:  Egy: %v", egy), 0, 0, 1.6, draw.RGB(0.2, 0.5, 0.3))
}
