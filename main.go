package main

import (
	"fmt"
	"golang-ninja/basic/shape"
)

func main() {
	square := shape.NewSquare(5)
	// circle := Circle{8}

	printShapeArea(square)

	// printShapeArea(circle)
	// printInterface(square)
	// printInterface(circle)
	// printInterface(212)
	printInterface("213214")
	// printInterface(13)
	// printInterface(true)
}

func printShapeArea(s shape.Shape) {
	fmt.Print(s.Area())
	fmt.Print(s.Perimeter())
}

func printInterface(i interface{}) {
	// switch value := i.(type) {
	// case int:Circle
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