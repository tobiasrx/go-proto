package shapes

import (
	"example.com/proto/internal/vector2"
	"github.com/gonutz/prototype/draw"
)

func DrawRay(window draw.Window, ray vector2.Ray, color draw.Color) {
	from := ray.Origin
	to := ray.Origin.Add(ray.Dir)
	DrawArrow(window, from, to, color)
}
