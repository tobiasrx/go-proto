package shapes

import (
	"math"

	"example.com/proto/internal/vector2"
	"github.com/gonutz/prototype/draw"
)

func DrawArrow(window draw.Window, from, to vector2.Vector, color draw.Color) {
	x, y := from.Point()
	u, v := to.Point()

	if x == u && y == v {
		return
	}

	window.DrawLine(x, y, u, v, color)

	vec := from.Subtract(to)

	vec1 := vec.RotateAndNormalize(math.Pi / 6)
	vec2 := vec.RotateAndNormalize(-math.Pi / 6)
	vec1 = vec1.Multiply(8.0)
	vec2 = vec2.Multiply(8.0)
	vec1 = vec1.Add(to)
	vec2 = vec2.Add(to)

	x, y = to.Point()
	u, v = vec1.Point()
	window.DrawLine(x, y, u, v, color)
	u, v = vec2.Point()
	window.DrawLine(x, y, u, v, color)
}
