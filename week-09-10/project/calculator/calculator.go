package calculator

import "errors"

// Common errors
var (
	ErrDivideByZero = errors.New("division by zero")
	ErrNegativeRoot = errors.New("cannot calculate square root of negative number")
)

// Add returns the sum of two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference of two numbers
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// SquareRoot returns the square root of a number
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeRoot
	}
	
	// For simplicity, using an approximation method
	// For real implementation, we'd use math.Sqrt()
	
	// Newton's method for square root
	guess := a / 2
	for i := 0; i < 10; i++ {
		guess = (guess + a/guess) / 2
	}
	
	return guess, nil
}

// Power returns a raised to the power of b
func Power(a, b float64) float64 {
	if b == 0 {
		return 1
	}
	
	result := a
	for i := 1; i < int(b); i++ {
		result *= a
	}
	
	return result
}

// TODO: Implement more calculator functions
