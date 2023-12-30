package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Point struct { //структура Point
	x int
	y int
}

func NewP(x, y int) *Point {
	return &Point{x: x, y: y}
}

func (p Point) distancePoint(o Point) float64 { //функция для нахождения расстояния между двумя точками
	f := math.Pow(float64(o.x-p.x), 2) //sqrt(x2 - x1)^2)
	s := math.Pow(float64(o.y-p.y), 2) //sqrt(y2 - y1)^2)
	return math.Sqrt(f + s)
}

func main() {
	fPoint := NewP(rand.Intn(10), rand.Intn(10))
	sPoint := NewP(rand.Intn(10), rand.Intn(10))
	fmt.Println(fPoint.distancePoint(*sPoint))
}
