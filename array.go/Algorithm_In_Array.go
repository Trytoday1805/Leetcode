package array

// algorithm often use TWO POINTER.

// ___________________________________________________________
// TWO POINTER:
// Điều kiện: Mảng phải được sắp xếp trước khi áp dụng.
// Hiệu quả: Thay vì thử mọi cặp phần tử trong mảng, thuật toán này giúp giảm đáng kể số lượng phép tính cần thực hiện bằng cách khai thác sự sắp xếp của mảng để loại bỏ các tổ hợp không cần thiết.
// Độ phức tạp: O(n) cho việc duyệt hai con trỏ trong mỗi vòng lặp, và O(n^2) cho toàn bộ thuật toán vì có một vòng lặp i duyệt qua mảng.
