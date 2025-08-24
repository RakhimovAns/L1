package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Point) GetX() float64 {
	return p.x
}

func (p Point) GetY() float64 {
	return p.y
}

func main() {
	point1 := NewPoint(3.0, 4.0)
	point2 := NewPoint(6.0, 8.0)

	distance := point1.Distance(point2)

	fmt.Println(point1.GetX(), point1.GetY())
	fmt.Println(point2.GetX(), point2.GetY())
	fmt.Println(distance)

}
