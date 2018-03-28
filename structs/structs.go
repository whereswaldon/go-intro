package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func distance(a, b Point) float64 {
	deltaX := float64(a.X - b.X)
	deltaY := float64(a.Y - b.Y)
	deltaX *= deltaX
	deltaY *= deltaY
	return math.Sqrt(deltaX + deltaY)
}

func mutate(p Point) {
	p.X = -11
}

func main() {
	f := Point{X: 4, Y: 9}
	q := Point{X: 3, Y: 4}
	fmt.Println(distance(f, q))
	fmt.Println(q)
	mutate(q)
	fmt.Println(q)
}
