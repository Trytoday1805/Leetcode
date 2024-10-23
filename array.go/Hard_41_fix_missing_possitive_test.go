package array

import (
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{1, 2, 0}, expected: 3},
		{nums: []int{3, 4, -1, 1}, expected: 2},
		{nums: []int{7, 8, 9, 11, 12}, expected: 1},
		{nums: []int{1}, expected: 2},
		{nums: []int{2, 3, 4}, expected: 1},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			result := firstMissingPositive(tc.nums)
			if result != tc.expected {
				t.Errorf("for input %v, expected %d, got %d", tc.nums, tc.expected, result)
			}
		})
	}
}
