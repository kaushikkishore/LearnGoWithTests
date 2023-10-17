package main

import "fmt"

type Shape interface {
	Area() float64
}

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

func main() {
	printArea(Triangle{10.0, 10.0})
	printArea(Square{10.0})
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) Area() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}
