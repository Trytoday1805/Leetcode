package array

import (
    "testing"
)

func TestMaxContainer(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{1,8,6,2,5,4,8,3,7}, expected: 49},
        {nums: []int{1,1}, expected: 1},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			result := maxArea(tc.nums)
			if result != tc.expected {
				t.Errorf("for input %v, expected %d, got %d", tc.nums, tc.expected, result)
			}
		})
	}
}
