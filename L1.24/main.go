package main

import (
	"fmt"
	"math"
)

// Point - структура для представления точки на плоскости
type Point struct {
	x, y float64
}

// NewPoint - конструктор точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Distance вычисляет евклидово расстояние от текущей точки до другой точки other
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаём две точки
	p1 := NewPoint(1.5, 2.5)
	p2 := NewPoint(4.5, 6.5)

	// Вычисляем расстояние между ними
	dist := p1.Distance(p2)

	fmt.Printf("Расстояние между точками (%g, %g) и (%g, %g) = %g\n",
		p1.x, p1.y, p2.x, p2.y, dist)
}
