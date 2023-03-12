package main

import (
	"fmt"
	"math"

	"example.com/proto/internal/shapes"
	"example.com/proto/internal/vector2"
	"github.com/gonutz/prototype/draw"
)

func main() {
	draw.RunWindow("Title", 640, 480, update)
}

func radToDegrees(r float64) float64 {
	return r / math.Pi * 180
}

var rotation = math.Pi / 4
var origin vector2.Vector = vector2.FromPoint(0, 0)

func update(window draw.Window) {
	screen := vector2.FromPoint(window.Size())
	center := screen.Divide(2)
	mouse := vector2.FromPoint(window.MousePosition())

	window.ShowCursor(false)
	from := mouse.Add(vector2.FromRotation(math.Pi / 4).Multiply(14))
	shapes.DrawArrow(window, from, mouse, draw.White)

	di := vector2.FromRotation(rotation)
	ray := vector2.Ray{Origin: origin, Dir: di}

	bb := vector2.InitBoundingBox(vector2.FromPoint(0, 0), screen)

	vecs := bb.Intersect(ray)
	for _, vec := range vecs {
		x, y := vec.Point()
		window.FillEllipse(x-5, y-5, 10, 10, draw.LightPurple)
	}

	if len(vecs) == 2 {
		x, y := vecs[0].Point()
		x2, y2 := vecs[1].Point()
		window.DrawLine(x, y, x2, y2, draw.Red)
	}

	for _, click := range window.Clicks() {
		if click.Button == draw.LeftButton {
			origin = vector2.FromPoint(click.X, click.Y)
		}
	}

	if window.IsMouseDown(draw.RightButton) {
		rotation += math.Pi / 32
	}

	window.DrawScaledText(fmt.Sprintf("Angle %f", radToDegrees(mouse.Subtract(center).Angle())), 0, 0, 1.6, draw.RGB(0.2, 0.5, 0.3))

	circle := vector2.Circle{Origin: vector2.Vector{X: 10, Y: -80}, Radius: 200}
	circle2 := vector2.Circle{Origin: vector2.Vector{X: 10, Y: 80}, Radius: 200}
	shapes.DrawCircle(window, circle2, draw.Red)

	vecs = circle.Intersect(ray)

	for _, vec := range vecs {
		fmt.Printf("%v", vec)
		x, y := vec.Point()
		window.FillEllipse(x-5, y-5, 10, 10, draw.LightPurple)
	}

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		squareDist := vector2.FromPoint(click.X, click.Y).Subtract(center).LengthSqr()
		if squareDist <= 20*20 {
			// close the window and end the application
			window.Close()
		}
	}
}
