// pkg/stdlib/math.go
package stdlib

// Add adds two numbers and returns the result.
func Add(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		return a + b.(int)
	case float64:
		return a + b.(float64)
	default:
		return nil
	}
}

// Subtract subtracts two numbers and returns the result.
func Subtract(a, b interface{}) interface{} {
	// Implement the Subtract function
	return nil
}

// Multiply multiplies two numbers and returns the result.
func Multiply(a, b interface{}) interface{} {
	// Implement the Multiply function
	return nil
}

// Divide divides two numbers and returns the result.
func Divide(a, b interface{}) interface{} {
	// Implement the Divide function
	return nil
}
