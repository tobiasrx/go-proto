package shapes

import (
	"example.com/proto/internal/vector2"
	"github.com/gonutz/prototype/draw"
)

func DrawCircle(window draw.Window, circle vector2.Circle, color draw.Color) {
	x, y := circle.Origin.Subtract(vector2.Vector{X: circle.Radius, Y: circle.Radius}).Point()
	window.DrawEllipse(x, y, int(circle.Radius*2), int(circle.Radius*2), color)
}
