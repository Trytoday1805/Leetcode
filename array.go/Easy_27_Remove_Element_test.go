package array

import (
	"testing"
)

//[3,2,2,3]
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{3, 2, 2,3}, expected: 2},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			result := removeElement(tc.nums,3)
			if result != tc.expected {
				t.Errorf("for input %v, expected %d, got %d", tc.nums, tc.expected, result)
			}
		})
	}
}
