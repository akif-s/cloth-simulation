package main

import (
	"image/color"
)

type Point struct {
	pos      Vector2
	pPos     Vector2
	color    color.NRGBA
	isPinned bool
	isActive bool
}

//var deltaTime float64 = 0.4

func NewPoint(x, y float64, color color.NRGBA, isPinned, isActive bool) *Point {
	p := &Point{
		pos:      NewVector2(x, y),
		pPos:     NewVector2(x, y),
		color:    color,
		isPinned: isPinned,
		isActive: isActive,
	}

	return p
}

var oldDt float64 = 0.000001

func (p *Point) update(dt float64) {
	acc := NewVector2(0, 0.8)

	tmpPos := p.pos

	if !p.isPinned {

		// p.x = p.x + (p.x-p.px)*(dt/oldDt) + float64(accX)*(dt+oldDt)/2*dt
		p.pos = p.pos.Sum(p.pos.Substract(p.pPos).Product(dt / oldDt).Sum(acc.Product((dt + oldDt) / 2 * dt)))
		oldDt = dt
	}

	p.pPos = tmpPos
}
