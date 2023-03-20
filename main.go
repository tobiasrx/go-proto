package main

import (
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
	Pos:    vector2.FromPoint(50, 100),
	Vel:    vector2.FromPoint(5, 0),
	Mass:   5.0,
	Radius: 50,
}

var body2 = physics.Body{
	Pos:    vector2.FromPoint(400, 100),
	Vel:    vector2.FromPoint(1, 0),
	Mass:   1.0,
	Radius: 20,
}

func update(window draw.Window) {
	body.Draw(window)
	body2.Draw(window)

	if window.WasKeyPressed(draw.KeySpace) {
		body.Step(window)
		body2.Step(window)
		body.Collide(&body2)
	}

	// window.DrawScaledText(fmt.Sprintf("Angle"), 0, 0, 1.6, draw.RGB(0.2, 0.5, 0.3))
}
