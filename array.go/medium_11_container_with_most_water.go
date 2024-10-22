package array

import (
	"math"
)

// Case Time Limit expected
func maxAreaTimeout(height []int) int {
	maxArea := 0

	for i, val := range height {
		for k := i + 1; k < len(height); k++ {
			// Tính diện tích dựa trên chiều cao ngắn hơn và độ rộng
			area := int(math.Min(float64(val), float64(height[k]))) * (k - i)
			// Cập nhật maxArea nếu diện tích lớn hơn
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func maxArea(height []int) int {
	maxArea := 0

	for i, val := range height {
		for k := i + 1; k < len(height); k++ {
			// Tính diện tích dựa trên chiều cao ngắn hơn và độ rộng
			area := int(math.Min(float64(val), float64(height[k]))) * (k - i)
			// Cập nhật maxArea nếu diện tích lớn hơn
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}