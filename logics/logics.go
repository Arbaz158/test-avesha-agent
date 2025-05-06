package logics

// Max returns the larger of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Contains checks if a string slice contains a specific string
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Filter returns a new slice containing only the elements that satisfy the predicate function
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map applies a function to each element of a slice and returns a new slice with the results
func Map[T any, R any](slice []T, transform func(T) R) []R {
	result := make([]R, len(slice))
	for i, item := range slice {
		result[i] = transform(item)
	}
	return result
}

// DivideSlice divides each element in a slice by a divisor
// BUG: No check for division by zero
func DivideSlice(numbers []int, divisor int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = num / divisor // Potential division by zero if divisor is 0
	}
	return result
}

// GetNthElement returns the nth element from a slice
// BUG: No bounds checking
func GetNthElement(slice []string, index int) string {
	return slice[index] // Potential out of bounds access if index >= len(slice) or index < 0
}
