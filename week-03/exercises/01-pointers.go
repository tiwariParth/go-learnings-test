package main

import "fmt"

// This exercise focuses on pointers and how to use them to modify values

// swapInts swaps the values of two integers
func swapInts(a, b *int) {
	// TODO: Implement the swap using pointers
}

// doubleSlice doubles each element in a slice
func doubleSlice(numbers []int) {
	// TODO: Implement the doubling of each element
	// Note: Slices are reference types, so you don't need to return anything
}

// updatePerson updates the age and address of a person
func updatePerson(person *Person, newAge int, newAddress string) {
	// TODO: Implement updating the person's age and address
}

// Person represents a person with name, age and address
type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// Test swapInts
	x, y := 10, 20
	fmt.Printf("Before swap: x=%d, y=%d\n", x, y)
	// TODO: Call swapInts
	fmt.Printf("After swap: x=%d, y=%d\n", x, y)
	
	// Test doubleSlice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before doubling: %v\n", numbers)
	// TODO: Call doubleSlice
	fmt.Printf("After doubling: %v\n", numbers)
	
	// Test updatePerson
	person := Person{
		Name:    "Alice",
		Age:     30,
		Address: "123 Main St",
	}
	
	fmt.Printf("Before update: %+v\n", person)
	// TODO: Call updatePerson
	fmt.Printf("After update: %+v\n", person)
}
