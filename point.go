package main

import (
	"image/color"
)

type Point struct {
	x, y     float64
	px, py   float64
	color    color.NRGBA
	isPinned bool
	isActive bool
}

//var deltaTime float64 = 0.4

func NewPoint(x, y float64, color color.NRGBA, isPinned, isActive bool) *Point {
	p := &Point{
		x:        x,
		y:        y,
		px:       x,
		py:       y,
		color:    color,
		isPinned: isPinned,
		isActive: isActive,
	}

	return p
}

var oldDt float64 = 0.000001

func (p *Point) update(dt float64) {
	accX := 0.
	accY := 1.2

	tmpx, tmpy := p.x, p.y

	if !p.isPinned {

		p.x = p.x + (p.x-p.px)*(dt/oldDt) + float64(accX)*(dt+oldDt)/2*dt
		p.y = p.y + (p.y-p.py)*(dt/oldDt) + float64(accY)*(dt+oldDt)/2*dt
		oldDt = dt
	}

	p.px = tmpx
	p.py = tmpy
}
