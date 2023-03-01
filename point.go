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

func (p *Point) update(dt float64) {
	accX := 0
	accY := 1

	tmpx, tmpy := p.x, p.y

	if !p.isPinned {
		p.x = 2*p.x - p.px + float64(accX)*dt*dt
		p.y = 2*p.y - p.py + float64(accY)*dt*dt
	}

	p.px = tmpx
	p.py = tmpy
}
