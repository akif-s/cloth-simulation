package main

import (
	"image/color"
)

type Point struct {
	pos      Vector2
	vel      Vector2
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

func (p *Point) update(dt float64) {
	acc := NewVector2(0, .8)

	tmpPos := p.pos

	if !p.isPinned {

		p.vel = p.vel.Sum(acc)

		if p.pos.y >= WINDOW_HEIGHT*2 {
			p.pos.y = WINDOW_HEIGHT * 2
		} else {
			// Verlet Integration
			p.pos = p.pos.Sum(p.pos.Substract(p.pPos).Product(0.99)).Sum(acc.Product(dt * dt))
		}

		p.vel.x, p.vel.y = 0.0, 0.0

	}

	p.pPos = tmpPos
}
