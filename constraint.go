package main

import (
	"image/color"
)

type Constraint struct {
	p1, p2 *Point
	len    float64
	color  color.NRGBA
}

func newConstraint(p1, p2 *Point, color color.NRGBA) *Constraint {

	//l := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)) // len of the constraint

	l := p2.pos.Substract(p1.pos).Magnitude()

	c := &Constraint{
		p1:    p1,
		p2:    p2,
		len:   l,
		color: color,
	}
	return c
}

func (c *Constraint) update() {

	d1 := c.p2.pos.Substract(c.p1.pos)
	d2 := d1.Magnitude()
	d3 := (d2 - c.len) / d2

	offset := d1.Product(d3 / 2.5)

	/* d := c.p2.pos.Substract(c.p1.pos)
	dist := d.Magnitude()

	diff := (c.len - dist) / dist

	mul := diff * 0.35 * (1 - c.len/dist)

	offset := d.Product(mul) */

	if !c.p1.isPinned {
		c.p1.pos = c.p1.pos.Sum(offset)
	}
	if !c.p2.isPinned {
		c.p2.pos = c.p2.pos.Substract(offset)
	}

}
