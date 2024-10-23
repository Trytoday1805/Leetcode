package array

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	test := []struct {
		nums     []int
		target   int
		expected int
	}{
		{nums: []int{-1, 2, 1, -4}, target: 1, expected: 2},
		{nums: []int{0, 0, 0}, target: 1, expected: 0},
	}

	for _, tc := range test {
		t.Run("", func(t *testing.T) {
			result := threeSumClosest(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("for input %v, expected %d, got %d", tc.nums, tc.expected, result)
			}
		})
	}
}
