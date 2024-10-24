package array

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	test := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{-1, 2, 1, -4}, expected: 4},
		{nums: []int{0, 0, 0},  expected: 1},
	}

	for _, tc := range test {
		t.Run("", func(t *testing.T) {
			result := removeDuplicates(tc.nums)
			if result != tc.expected {
				t.Errorf("for input %v, expected %d, got %d", tc.nums, tc.expected, result)
			}
		})
	}
}
