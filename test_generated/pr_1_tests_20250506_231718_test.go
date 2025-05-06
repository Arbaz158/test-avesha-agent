package test_generated

import (
	"test-avesha-agent/logics"
	"test-avesha-agent/math"
	"testing"
)

func TestMax(t *testing.T) {
	result := logics.Max(10, 5)
	if result != 10 {
		t.Errorf("Max(10, 5) = %d; want 10", result)
	}
}

func TestMin(t *testing.T) {
	result := logics.Min(10, 5)
	if result != 5 {
		t.Errorf("Min(10, 5) = %d; want 5", result)
	}
}

func TestContains(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	result := logics.Contains(slice, "banana")
	if !result {
		t.Errorf("Contains(slice, "banana") = %t; want true", result)
	}
}

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := logics.Filter(slice, func(n int) bool { return n%2 == 0 })
	expected := []int{2, 4}
	if len(result) != len(expected) {
		t.Errorf("Filter(slice, predicate) = %v; want %v", result, expected)
	}
}

func TestMap(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := logics.Map(slice, func(n int) int { return n * 2 })
	expected := []int{2, 4, 6, 8, 10}
	if len(result) != len(expected) {
		t.Errorf("Map(slice, transform) = %v; want %v", result, expected)
	}
}

func TestDivideSlice(t *testing.T) {
	numbers := []int{10, 20, 30, 40, 50}
	divisor := 2
	result := logics.DivideSlice(numbers, divisor)
	expected := []int{5, 10, 15, 20, 25}
	if len(result) != len(expected) {
		t.Errorf("DivideSlice(numbers, divisor) = %v; want %v", result, expected)
	}
}

func TestGetNthElement(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	index := 1
	result := logics.GetNthElement(slice, index)
	expected := "banana"
	if result != expected {
		t.Errorf("GetNthElement(slice, index) = %s; want %s", result, expected)
	}
}

func TestAdd(t *testing.T) {
	result := math.Add(10, 5)
	if result != 15 {
		t.Errorf("Add(10, 5) = %d; want 15", result)
	}
}

func TestSubtract(t *testing.T) {
	result := math.Subtract(10, 5)
	if result != 5 {
		t.Errorf("Subtract(10, 5) = %d; want 5", result)
	}
}

func TestMultiply(t *testing.T) {
	result := math.Multiply(10, 5)
	if result != 50 {
		t.Errorf("Multiply(10, 5) = %d; want 50", result)
	}
}

func TestDivide(t *testing.T) {
	result, _ := math.Divide(10, 5)
	if result != 2 {
		t.Errorf("Divide(10, 5) = %d; want 2", result)
	}
}

func TestPower(t *testing.T) {
	result := math.Power(2, 3)
	if result != 8 {
		t.Errorf("Power(2, 3) = %d; want 8", result)
	}
}

func TestSquareRoot(t *testing.T) {
	result, _ := math.SquareRoot(16)
	if result != 4 {
		t.Errorf("SquareRoot(16) = %d; want 4", result)
	}
}

func TestSqrt(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{4, 2},
		{9, 3},
		{16, 4},
	}

	for _, c := range cases {
		result, _ := Sqrt(c.input)
		if result != c.output {
			t.Errorf("Sqrt(%d) == %d, want %d", c.input, result, c.output)
		}
	}
}

func TestAbs(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{-5, 5},
		{0, 0},
		{10, 10},
	}

	for _, c := range cases {
		result := Abs(c.input)
		if result != c.output {
			t.Errorf("Abs(%d) == %d, want %d", c.input, result, c.output)
		}
	}
}

