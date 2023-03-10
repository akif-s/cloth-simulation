package main

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type Cloth struct {
	X           int // upleft x-coordinate of cloth
	Y           int // upleft y-coordinate of cloth
	width       int // number of points in x-axis
	height      int // number of points in y-axis
	gap         int // distance between two points
	drawPoints  bool
	points      []*Point
	constraints []*Constraint
}

func newCloth(X, Y, w, h, gap int, drawPoints bool) *Cloth {
	var points []*Point
	var constraints []*Constraint

	// Create the points
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			var pin bool
			if y == 0 && (x == 0 || (x+1)%7 == 0) {
				pin = true
			}
			points = append(points, NewPoint(float64(x*gap+X), float64(y*gap+Y), color.NRGBA{R: 0, G: 0, B: 0, A: 0xff}, pin, true))
		}
	}

	// Creat the constraintse
	for i := 0; i < h; i++ { // y-axis
		for j := 0; j < w; j++ { // x-axis
			count := j + i*w
			if j < w-1 {
				c := newConstraint(points[count], points[count+1], color.NRGBA{R: 0, G: 0, B: 0, A: 0xff}) // connect to right adjacent point
				constraints = append(constraints, c)
			}
			if count < w*(h-1) {
				c := newConstraint(points[count], points[count+w], color.NRGBA{R: 0, G: 0, B: 0, A: 0xff}) // connect to bottom adjacent point
				constraints = append(constraints, c)
			}
		}
	}

	return &Cloth{
		X:           X,
		Y:           X,
		width:       w,
		height:      h,
		gap:         gap,
		drawPoints:  drawPoints,
		points:      points,
		constraints: constraints,
	}
}

func (cloth *Cloth) draw(gtx layout.Context, dt float64) {

	//Drawing and updating the points
	for _, p := range cloth.points {
		if !p.isActive {
			continue
		}
		if cloth.drawPoints { // Don't draw the points if not wanted.
			a := clip.Rect{Min: image.Pt(int(p.pos.X), int(p.pos.Y)), Max: image.Pt(int(p.pos.X)+5, int(p.pos.Y)+5)}.Push(gtx.Ops)
			paint.ColorOp{Color: p.color}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			a.Pop()
		}
		if !p.isPinned {
			p.update(dt)
		}

	}

	//Drawing and updating the constraints
	for _, c := range cloth.constraints {
		if c.p1.isActive && c.p2.isActive {
			var path clip.Path
			path.Begin(gtx.Ops)
			path.MoveTo(f32.Pt(float32(c.p1.pos.X+2.5), float32(c.p1.pos.Y+2.5)))
			path.LineTo(f32.Pt(float32(c.p2.pos.X+2.5), float32(c.p2.pos.Y+2.5)))
			path.Close()

			a := clip.Stroke{Path: path.End(), Width: 1}.Op().Push(gtx.Ops)
			paint.ColorOp{Color: c.color}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			a.Pop()
			c.update()
		}
	}
}
