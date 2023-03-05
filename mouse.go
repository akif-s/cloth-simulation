package main

type Mouse struct {
	pos     Vector2
	isPress bool
	isDragg bool
}

func NewMouse(x, y float64) *Mouse {
	return &Mouse{
		pos: NewVector2(x, y),
	}
}

func (m *Mouse) SetPosition(pos Vector2) {
	m.pos = pos
}

func (m *Mouse) SetPress(isPressed bool) {
	m.isPress = isPressed
}

func (m *Mouse) SetDragg(isDragg bool) {
	m.isDragg = isDragg
}
