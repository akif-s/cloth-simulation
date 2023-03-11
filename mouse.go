package main

import "github.com/akif-s/vector"

type Mouse struct {
	pos     vector.Vector2
	isPress bool
	isDragg bool
}

func NewMouse(x, y float64) *Mouse {
	return &Mouse{
		pos: vector.NewVector2(x, y),
	}
}

func (m *Mouse) SetPosition(pos vector.Vector2) {
	m.pos = pos
}

func (m *Mouse) SetPress(isPressed bool) {
	m.isPress = isPressed
}

func (m *Mouse) SetDragg(isDragg bool) {
	m.isDragg = isDragg
}
