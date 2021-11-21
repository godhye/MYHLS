package main

import (
	"fmt"
	"math"
)

//인터페이스 정의
type Shape interface {
	area() float64
	perimeter() float64
	plussize()
}

type Rect struct {
	width, heigth float64
}

type Circle struct {
	radius float64
}

//인터페이스 구현
func (r Rect) area() float64 {
	return r.width * r.heigth
}

func (r Rect) perimeter() float64 {
	return 2*r.width + 2*r.heigth
}

func (r *Rect) plussize() {
	r.heigth = r.heigth + 1.
	r.width = r.width + 1.
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}
func (c *Circle) plussize() {
	c.radius = c.radius + 1
}

func main() {

	r := Rect{10., 29.}
	c := Circle{2.}

	fmt.Printf("%.2f\n", r.area())
	fmt.Printf("%.2f\n", r.perimeter())

	fmt.Printf("%.2f\n", c.area())
	fmt.Printf("%.2f\n", c.perimeter())

	r.plussize()
	c.plussize()

	fmt.Printf("%.2f\n", r.area())
	fmt.Printf("%.2f\n", r.perimeter())
	fmt.Printf("%.2f\n", c.area())
	fmt.Printf("%.2f\n", c.perimeter())

}
