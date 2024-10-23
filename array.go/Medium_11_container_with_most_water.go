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

// Ap dung Thuật toán Brute Force, còn được gọi là thuật toán vét cạn,
// là một phương pháp giải quyết vấn đề bằng cách thử tất cả các trường hợp có thể xảy ra cho đến khi tìm ra giải pháp chính xác
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0

	for l < r {
		var minHeight int

		if height[l] < height[r] {
			minHeight = height[l]
		} else {
			minHeight = height[r]
		}

		// Nếu có max mới thì cập nhật
		// So với cái trên phải chạy 2 for lồng => sẽ là n*(n-1) thì cái này chỉ mất n-1 = n
		if maxArea < (minHeight * (r - l)) {
			maxArea = minHeight * (r - l)
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
