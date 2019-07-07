package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}
type area interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(a area) {
	fmt.Println(a.getArea())
}

func main() {
	sq := square{
		sideLength: 10,
	}
	tri := triangle{
		base:   10,
		height: 20,
	}
	printArea(sq)
	printArea(tri)

}
