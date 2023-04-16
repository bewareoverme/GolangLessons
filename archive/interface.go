package main

import (
	"fmt"
	"math"
)

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

func (s Square) Perimeter() float32 {
	return s.sideLength * 4
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{5}
	// circle := Circle{8}

	printShapeArea(square)
	// printShapeArea(circle)
	// printInterface(square)
	// printInterface(circle)
	printInterface(212)
	printInterface("213214")
	// printInterface(13)
	// printInterface(true)
}

func printShapeArea(shape Shape) {
	fmt.Print(shape.Area())
	fmt.Print(shape.Perimeter())
}

func printInterface(i interface{}) {
	// switch value := i.(type) {
	// case int:
	// 	fmt.Println("int", value)
	// case bool:
	// 	fmt.Println("bool", value)
	// default:
	// 	fmt.Println("unknown type", value)
	// }

	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
		return
	}
	fmt.Println(len(str))

	// fmt.Printf("%+v\n", i)
}