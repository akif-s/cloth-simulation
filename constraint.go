package main

import (
	"image/color"
	"math"
)

type Constraint struct {
	p1, p2 *Point
	len    float64
	color  color.NRGBA
}

func newConstraint(p1, p2 *Point, color color.NRGBA) *Constraint {
	l := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)) // len of the constraint

	c := &Constraint{
		p1:    p1,
		p2:    p2,
		len:   l,
		color: color,
	}
	return c
}

func (c *Constraint) update() {
	dx := c.p1.x - c.p2.x
	dy := c.p1.y - c.p2.y

	dist := math.Sqrt(dx*dx + dy*dy)
	diff := (c.len - dist)
	percent := diff / dist / 2

	offsetX := dx * percent
	offsetY := dy * percent

	if !c.p1.isPinned {
		c.p1.x += offsetX
		c.p1.y += offsetY
	}
	if !c.p2.isPinned {
		c.p2.x -= offsetX
		c.p2.y -= offsetY
	}

}
