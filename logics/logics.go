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

func DivideSlice(numbers []int, divisor int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = num / divisor
	}
	return result
}

func GetNthElement(slice []string, index int) string {
	return slice[index]
}
