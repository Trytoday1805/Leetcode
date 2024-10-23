package array

import (
	"sort"
	"math"
)
// Input : Array INTERGER
// OUTPUT: Sum 3 interger in array same target
// MINSET: Target - sum(a+b+c) = min return sum(a+b+c)
// HOW TO HANDLE sum(a+b+C) optimize {SORT Array -> Use Two Pointer}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
    closest := nums[0] + nums[1] + nums[2]

	for i:= 0 ; i < len(nums) -2 ; i++ {
		left, right := i+1, len(nums)-1
		// For l < r 
		for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == target {
                return sum
            }
            
            if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
                closest = sum
            }
            
            if sum < target {
                left++
            } else {
                right--
            }
        }
	}
	return closest
}