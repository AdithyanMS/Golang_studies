package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}
type shape interface {
	getArea() float64
}

func main() {
	ft := triangle{
		base:   2.04,
		height: 2.04,
	}
	fq := square{
		sideLength: 2.04,
	}
	printArea(ft)
	printArea(fq)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Printf("the area is %v \n", s.getArea())
}
