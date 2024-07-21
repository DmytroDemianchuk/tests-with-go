package main

import "testing"

func TestMax(t *testing.T) {
	numbers := []int{1, 2, 5, 7, -2, 16, 15}
	expected := 15

	result := Max(numbers)

	if result != expected {
		t.Errorf("Incorrect result. Expect %d, got %d",
			expected, result)
	}
}
