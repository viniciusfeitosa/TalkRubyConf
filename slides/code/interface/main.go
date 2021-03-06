package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Square struct {
	width, height float64
}

func (s *Square) area() float64 {
	return 2*s.width + 2*s.height
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := &Circle{0, 0, 5}
	s := &Square{5, 5}
	fmt.Println(totalArea(c))
	fmt.Println(totalArea(s))
}
