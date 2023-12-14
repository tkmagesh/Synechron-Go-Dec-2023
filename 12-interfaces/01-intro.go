package main

import (
	"fmt"
	"math"
)

// v1.0
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Define a Rectangle struct (height & width)
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// v2.0
// utility

type AreaFinder interface {
	Area() float64
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

// v3.0 (perimeter)
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type PerimeterFinder interface {
	Perimeter() float64
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// v4.0
type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

// v5.0 (introducing Triangle)
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) Perimeter() float64 {
	return t.Base + math.Sqrt((t.Base*t.Base)+(4*t.Height*t.Height))
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShapeStats(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShapeStats(r)

	t := Triangle{Base: 10, Height: 15}
	PrintShapeStats(t)
}
