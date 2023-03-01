package main

import "image/color"

type Point struct {
	x, y     float64
	px, py   float64
	color    color.NRGBA
	isPinned bool
}

var deltaTime float64 = 0.1

func NewPoint(x, y float64, color color.NRGBA, isPinned bool) *Point {
	p := &Point{
		x:        x,
		y:        y,
		px:       x,
		py:       y,
		color:    color,
		isPinned: isPinned,
	}

	return p
}

func (p *Point) update() {
	accX := 0
	accY := 1

	tmpx, tmpy := p.x, p.y

	if !p.isPinned {
		p.x = 2*p.x - p.px + float64(accX)*deltaTime*deltaTime
		p.y = 2*p.y - p.py + float64(accY)*deltaTime*deltaTime
	}

	p.px = tmpx
	p.py = tmpy
}