package array

// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3,_,_,_]
func removeElement(nums []int, val int) int {
    lenNums := len(nums)
    for i := 0; i < lenNums; i++ {
        if nums[i] == val {
            for j := i; j < lenNums-1; j++ {
                nums[j] = nums[j+1]
            }
            
            lenNums--
            i--
        }
    }
    return lenNums
}
