package Mortar

import "math"

func (m *Mortar) CalculateCalculations() {
	dx := m.hitX - m.mortarX
	dy := m.hitY - m.mortarY

	// 1. Regner ut lengden (avstanden)
	m.distance = math.Hypot(dx, dy)

	// 2. Regner ut vinkelen i radianer
	m.angle = math.Atan2(dy, dx)

	// 3. Konverterer vinkelen til grader
	m.degree = m.angle * (180 / math.Pi)
}

// NewMortar oppretter en ny Mortar og regner ut lengde og vinkler automatisk
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
