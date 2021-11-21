package main

import "fmt"

//인터페이스 정의
type Shape interface {
	plussize()
	getsize() float64
}
type Rect struct {
	width, heigth float64
}

type Circle struct {
	radius float64
}

//인터페이스 구현

//포인터 전달하여 struct내 값 변경 반영
//point receiver

func (r *Rect) plussize() {
	r.heigth = r.heigth + 1.
	r.width = r.width + 1.
	println("Rect plussize\n ")
}
func (r *Rect) getsize() float64 {

	println("Rect getsize\n ")
	return r.heigth + r.heigth
}

//반환값 포인터로 주지않으면 value copy
//Value receiver
func (c *Circle) plussize() {
	c.radius = c.radius + 1
	println("Circle plussize\n ")
}
func (c *Circle) getsize() float64 {

	println("Circle getsize\n ")
	return c.radius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.getsize()
		s.plussize()
		println(a)

	}
}

func printk(v interface{}) {
	fmt.Println(v)
}

func main() {

	//구조체 포인터로 선언
	r := &Rect{10., 29.}
	c := &Circle{2.}

	r.plussize() //point receiver
	c.plussize() //value receiver

	//pointer receiver
	showArea(r, c)

	//빈 인터페이스 void*와 비슷
	var x interface{}

	x = 1
	x = "string"

	println("interface{}  x")

	println(x) //주소값 2개 출력됨??
	printk(x)

	println("interface{}  intx")
	var intx interface{}

	intx = 4
	intx = 33

	println(intx)
	printk(intx)

}
