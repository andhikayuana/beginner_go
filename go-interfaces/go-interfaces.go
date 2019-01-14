package main

import (
	"fmt"
	"math"
)

// define interface
type Shape interface {
	area() float64
}

// define circle
type Circle struct {
	x, y, radius float64
}

//define rectagle
type Rectangle struct {
	width, height float64
}

//method for circle (implementation of Shape.area())
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

//method for rectangle (implementation of Shape.area())
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

//mehotd for shape
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rect := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area : %f\n", getArea(circle))
	fmt.Printf("Rectangle area : %f\n", getArea(rect))
}
