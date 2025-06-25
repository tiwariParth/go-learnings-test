package calculator

import (
	"testing"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	// Table-driven test
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zeros", 0, 0, 0},
		{"large numbers", 1e9, 1e9, 2e9},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v, expected %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestSubtract tests the Subtract function
func TestSubtract(t *testing.T) {
	// Table-driven test
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -5, -3, -2},
		{"mixed signs", -5, 3, -8},
		{"zeros", 0, 0, 0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v, expected %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestMultiply tests the Multiply function
func TestMultiply(t *testing.T) {
	// TODO: Implement table-driven test for Multiply
}

// TestDivide tests the Divide function
func TestDivide(t *testing.T) {
	// Table-driven test
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectedErr error
	}{
		{"positive numbers", 6, 3, 2, nil},
		{"negative numbers", -6, -3, 2, nil},
		{"mixed signs", -6, 3, -2, nil},
		{"divide by zero", 6, 0, 0, ErrDivideByZero},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			
			// Check error
			if err != tt.expectedErr {
				t.Errorf("Divide(%v, %v) error = %v, expected error %v", tt.a, tt.b, err, tt.expectedErr)
				return
			}
			
			// If we expected an error, don't check the result
			if tt.expectedErr != nil {
				return
			}
			
			if result != tt.expected {
				t.Errorf("Divide(%v, %v) = %v, expected %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestSquareRoot tests the SquareRoot function
func TestSquareRoot(t *testing.T) {
	// TODO: Implement table-driven test for SquareRoot
	// Remember to check for error when input is negative
	// Use math.Abs and an epsilon for floating point comparison
}

// TestPower tests the Power function
func TestPower(t *testing.T) {
	// TODO: Implement table-driven test for Power
}
