package main

import "fmt"

const (
	pi = 3.1415
)

type Figure interface {
	Area() float64
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius int
}

func (r Rectangle) Area() float64 {
	return float64(r.length * r.width)
}

func (c Circle) Area() float64 {
	return pi * float64(c.radius*c.radius)
}

func main() {
	var fig Figure
	fig = Rectangle{length: 3, width: 4}
	fmt.Println("Площадь прямоугольника", fig.Area())
	fig = Circle{radius: 23}
	fmt.Println("Площадь круга", fig.Area())
	return
}
