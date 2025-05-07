package math

import "errors"

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	if b < 0 {
		a = -a
		b = -b
	}

	return a / b, nil
}

// Power returns a raised to the power of b
func Power(a, b int) int {
	if b == 0 {
		return 1
	}

	if b%2 == 0 {
		half := Power(a, b/2)
		return half * half
	} else {
		return a * Power(a, b-1)
	}
}

// SquareRoot returns the integer square root of a number
func SquareRoot(a int) (int, error) {
	if a < 0 {
		return 0, errors.New("cannot calculate square root of negative number")
	}
	if a == 0 || a == 1 {
		return a, nil
	}

	// Using binary search to find square root
	start := 1
	end := a
	result := 0

	for start <= end {
		mid := (start + end) / 2
		if mid*mid == a {
			return mid, nil
		}
		if mid*mid < a {
			start = mid + 1
			result = mid
		} else {
			end = mid - 1
		}
	}
	return result, nil
}

// Abs returns the absolute value of an integer
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
