# Go Fundamentals MCQ Test

Answer the following 20 multiple-choice questions to test your understanding of Go fundamentals.

## Basic Syntax & Structure

1. What is the correct way to declare a variable in Go?
   - [ ] A. `var x int = 10`
   - [ ] B. `x := 10`
   - [ ] C. `int x = 10`
   - [ ] D. Both A and B are correct

2. What is the zero value for an integer in Go?
   - [ ] A. nil
   - [ ] B. 0
   - [ ] C. undefined
   - [ ] D. ""

3. Which keyword is used to declare a new package in Go?
   - [ ] A. pkg
   - [ ] B. module
   - [ ] C. package
   - [ ] D. namespace

## Control Structures

4. Which loop construct does Go support?
   - [ ] A. while
   - [ ] B. do-while
   - [ ] C. for
   - [ ] D. All of the above

5. What will the following code print?
   ```go
   x := 5
   if x > 3 {
       fmt.Println("A")
   } else if x > 7 {
       fmt.Println("B")
   } else {
       fmt.Println("C")
   }
   ```
   - [ ] A. A
   - [ ] B. B
   - [ ] C. C
   - [ ] D. Nothing

## Data Types

6. Which of the following is NOT a basic data type in Go?
   - [ ] A. int
   - [ ] B. string
   - [ ] C. object
   - [ ] D. bool

7. What's the difference between an array and a slice in Go?
   - [ ] A. Arrays have fixed length, slices are dynamic
   - [ ] B. Arrays can only hold ints, slices can hold any type
   - [ ] C. Arrays are passed by value, slices are passed by reference
   - [ ] D. Both A and C

## Maps & Structs

8. How do you declare a map in Go?
   - [ ] A. `map[string]int{}`
   - [ ] B. `map<string, int>{}`
   - [ ] C. `new Map(string, int)`
   - [ ] D. `dict[string]int{}`

9. What does the following struct declaration define?
   ```go
   type Person struct {
       Name string
       Age  int
   }
   ```
   - [ ] A. A class with two properties
   - [ ] B. A struct with two fields
   - [ ] C. An interface with two methods
   - [ ] D. An enum with two values

## Functions & Methods

10. How do you declare a function in Go?
    - [ ] A. `function add(a int, b int) int {}`
    - [ ] B. `func add(a, b int) int {}`
    - [ ] C. `def add(a int, b int) int {}`
    - [ ] D. `procedure add(a int, b int) returns int {}`

11. What is the correct way to define a method on a struct?
    - [ ] A. `func Person.GetName() string {}`
    - [ ] B. `func GetName(p Person) string {}`
    - [ ] C. `func (p Person) GetName() string {}`
    - [ ] D. `func GetName() string for Person {}`

## Packages & Imports

12. How do you import multiple packages in Go?
    - [ ] A. `import fmt, math`
    - [ ] B. `import "fmt" import "math"`
    - [ ] C. `import ("fmt"; "math")`
    - [ ] D. `import ("fmt"; "math")`

13. What is the correct way to export a function in Go?
    - [ ] A. Use the `export` keyword
    - [ ] B. Start the function name with an uppercase letter
    - [ ] C. Use the `pub` keyword
    - [ ] D. Add `@export` annotation

## Error Handling

14. How are errors typically handled in Go?
    - [ ] A. Using try/catch blocks
    - [ ] B. Using exceptions
    - [ ] C. By returning error values
    - [ ] D. With the `@throws` annotation

15. What does the following function signature indicate?
    ```go
    func divide(a, b float64) (float64, error) {}
    ```
    - [ ] A. Function returns two float64 values
    - [ ] B. Function returns a float64 and might return an error
    - [ ] C. Function takes an error as input
    - [ ] D. Function throws an exception

## Concurrency Basics

16. Which keyword is used to start a goroutine?
    - [ ] A. `thread`
    - [ ] B. `async`
    - [ ] C. `go`
    - [ ] D. `spawn`

17. What is the purpose of channels in Go?
    - [ ] A. For communication between functions
    - [ ] B. For communication between goroutines
    - [ ] C. For handling file operations
    - [ ] D. For networking only

## Project Context

18. What is the correct way to store tasks in memory for our ToDo CLI app?
    - [ ] A. Using a database
    - [ ] B. Using a global variable
    - [ ] C. Using a slice or map
    - [ ] D. Using a text file

19. Which package would be most appropriate for parsing command-line arguments in our ToDo CLI?
    - [ ] A. `fmt`
    - [ ] B. `flag`
    - [ ] C. `args`
    - [ ] D. `cli`

20. How would you structure a Task in the ToDo CLI app?
    - [ ] A. As a map with string keys
    - [ ] B. As a custom struct
    - [ ] C. As a simple string
    - [ ] D. As JSON text

## Coding Tasks

1. Write a function that takes a slice of integers and returns the sum and average.

2. Create a simple struct to represent a book with title, author, and year fields, then write a function that prints all books published after a given year.
