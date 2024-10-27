package string
// algorithm 

// ___________________________________________________________
// Thuật toán này được gọi là thuật toán mở rộng từ trung tâm (expand around center). Đây là một phương pháp phổ biến để tìm chuỗi con đối xứng dài nhất (palindromic substring) trong một chuỗi cho trước.

// Cách hoạt động của thuật toán mở rộng từ trung tâm
// Ý tưởng chính là xem mỗi ký tự hoặc cặp ký tự trong chuỗi như là một trung tâm của chuỗi đối xứng. Sau đó, mở rộng từ trung tâm này ra hai bên để kiểm tra chuỗi con đối xứng.
// Thuật toán sẽ kiểm tra hai loại trung tâm:
// Trung tâm là một ký tự (dành cho các chuỗi đối xứng lẻ, ví dụ: "aba").
// Trung tâm là hai ký tự liên tiếp (dành cho các chuỗi đối xứng chẵn, ví dụ: "abba").
// Với mỗi trung tâm, thuật toán mở rộng ra hai bên và ghi nhận chiều dài của chuỗi đối xứng tìm được. Nếu chuỗi đối xứng mới dài hơn chuỗi trước đó, nó sẽ cập nhật kết quả.