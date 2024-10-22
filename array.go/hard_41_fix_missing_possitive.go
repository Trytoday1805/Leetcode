package array

//  MULTIPLE ASSIGNT GO
// GET n = len() b/c no need value these out of range n {need min} 
// Đưa nó về đúng vị trí VIDU: 
// Áp dụng :Cyclic Sort: Thuật toán sắp xếp theo chu kỳ, đưa mỗi phần tử về đúng vị trí của nó trong một vòng lặp.

func firstMissingPositive(nums []int) int {
    n := len(nums)

    // Đưa mỗi số về đúng vị trí của nó nếu nó thuộc khoảng [1, n]
    for i := 0; i < n; i++ {
        for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// Áp dụng multiple assigst mà ko cần biến trung gian
			// Cú pháp gán nhiều giá trị trong Go cho phép bạn gán nhiều biến cùng lúc trong một dòng lệnh.
			// Phép gán này diễn ra đồng thời, nghĩa là tất cả các biểu thức bên phải được đánh giá trước, sau đó các giá trị mới sẽ được gán cho các biến tương ứng ở bên trái cùng lúc.
            nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
        }
    }

    // Duyệt lại mảng để tìm số đầu tiên không đúng vị trí
    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i + 1
        }
    }

    // Nếu tất cả số đều đúng vị trí, trả về n + 1
    return n + 1
}

// Note : về cách hoạt động nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
// Chia ra trái / phải từ dấu bằng
// Trái : nums[nums[i]-1], nums[i] =
// Phải : = nums[i], nums[nums[i]-1]
// Nó sẽ chạy bên phải trc, num[0]=3 và nums[2] là giá trị đang ở vị trí 2
// Gán num[0] cho nums[nums[i]-1](bên trái) để đúng vị trí
// Gán giá trị nums[nums[i]-1] cho nums[i].