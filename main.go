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

var origin vector2.Vector = vector2.FromPoint(0, 0)

func update(window draw.Window) {
	screen := vector2.FromPoint(window.Size())
	center := screen.Divide(2)
	mouse := vector2.FromPoint(window.MousePosition())

	dir := mouse.Subtract(center)
	mouseInCircle := dir.LengthSqr() < 20*20

	color := draw.DarkRed
	if mouseInCircle {
		color = draw.Red
	}
	window.ShowCursor(false)
	from := mouse.Add(vector2.FromRotation(math.Pi / 4).Multiply(14))
	shapes.DrawArrow(window, from, mouse, draw.White)

	ray := vector2.Ray{Origin: origin, Dir: mouse.Subtract(origin)}

	bb := vector2.Init(vector2.FromPoint(0, 0), screen)

	v, ok := bb.Intersect(ray)
	if ok {
		x, y := v.Point()
		window.FillEllipse(x-5, y-5, 10, 10, draw.LightPurple)
		ox, oy := origin.Point()
		window.DrawLine(ox, oy, x, y, draw.White)

		for _, click := range window.Clicks() {
			if click.Button == draw.RightButton {
				fmt.Printf("%+v", v)
			}
		}
	}

	for _, click := range window.Clicks() {
		if click.Button == draw.LeftButton {
			origin = vector2.FromPoint(click.X, click.Y)
		}
	}

	window.DrawScaledText(fmt.Sprintf("Angle %f", radToDegrees(mouse.Subtract(center).Angle())), 0, 0, 1.6, draw.RGB(0.2, 0.5, 0.3))

	centerX, centerY := center.Point()
	window.FillEllipse(centerX-20, centerY-20, 40, 40, color)
	window.DrawEllipse(centerX-20, centerY-20, 40, 40, draw.White)
	if mouseInCircle {
		window.DrawScaledText("Close!", centerX-40, centerY+25, 1.6, draw.RGB(0.2, 0.5, 0.3))
	}

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		dx, dy := click.X-centerX, click.Y-centerY
		squareDist := dx*dx + dy*dy
		if squareDist <= 20*20 {
			// close the window and end the application
			window.Close()
		}
	}
}
