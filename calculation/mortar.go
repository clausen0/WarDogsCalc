package main

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

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

type CalculationResult struct {
	Distance float64 `json:"distance"`
	Angle    float64 `json:"angle"`
	Degree   float64 `json:"degree"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	mx, _ := strconv.ParseFloat(r.URL.Query().Get("mx"), 64)
	my, _ := strconv.ParseFloat(r.URL.Query().Get("my"), 64)
	hx, _ := strconv.ParseFloat(r.URL.Query().Get("hx"), 64)
	hy, _ := strconv.ParseFloat(r.URL.Query().Get("hy"), 64)

	m := NewMortar(mx, my, hx, hy)

	res := CalculationResult{
		Distance: m.distance,
		Angle:    m.angle,
		Degree:   m.degree,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
