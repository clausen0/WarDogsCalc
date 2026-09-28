package Mortar

import "math"

func (m *Mortar) CalculateCalculations() {
	dx := m.hitX - m.mortarX
	dy := m.hitY - m.mortarY

	m.distance = math.Hypot(dx, dy)

	m.angle = math.Atan2(dy, dx)

	m.degree = m.angle * (180 / math.Pi)
}

func NewMortar(mx, my, hx, hy float64) Mortar {
	m := Mortar{
		mortarX: mx,
		mortarY: my,
		hitX:    hx,
		hitY:    hy,
	}
	m.CalculateCalculations()
	return m
}
