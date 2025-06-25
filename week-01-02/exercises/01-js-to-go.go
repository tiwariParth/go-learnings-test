package main

import (
	"fmt"
	"encoding/json"
)

// Convert this JavaScript object to a Go struct:
/*
const person = {
  name: "John",
  age: 30,
  isEmployed: true,
  skills: ["JavaScript", "HTML", "CSS"],
  address: {
    street: "123 Main St",
    city: "Anytown",
    zipCode: "12345"
  }
};
*/

// TODO: Define the appropriate Go structs here

// TODO: Create a constructor function that creates a new Person

func main() {
	// TODO: Create a person instance with your own data
	
	// TODO: Print out the person details
	
	// TODO: Convert the person to JSON and print it
	
	// Bonus: Read a JSON string and convert it to a Person object
	jsonStr := `{"name":"Alice","age":25,"isEmployed":false,"skills":["Go","Python","SQL"],"address":{"street":"456 Oak Ave","city":"Techville","zipCode":"67890"}}`
	
	// TODO: Parse the JSON into a Person object and print it
}
