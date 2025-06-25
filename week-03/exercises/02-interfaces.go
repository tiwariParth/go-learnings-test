package main

import (
	"fmt"
	"math"
)

// Shape is an interface for objects with area and perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO: Define Rectangle struct with width and height

// TODO: Implement Area() method for Rectangle

// TODO: Implement Perimeter() method for Rectangle

// TODO: Define Circle struct with radius

// TODO: Implement Area() method for Circle

// TODO: Implement Perimeter() method for Circle

// TODO: Define Triangle struct with base and height

// TODO: Implement Area() method for Triangle

// TODO: Implement Perimeter() method for Triangle

// PrintShapeInfo prints information about a shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// TODO: Create instances of each shape
	
	// TODO: Call PrintShapeInfo on each shape
	
	// TODO: Store shapes in a slice and iterate over them to print info
	
	// TODO: (Bonus) Create a function that returns the shape with the largest area
}
