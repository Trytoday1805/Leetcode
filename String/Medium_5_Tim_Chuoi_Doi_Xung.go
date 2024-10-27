package string

// Input: s = "cbbd"
// Output: "bb"
// Hàm extendRange mở rộng từ vị trí l và r, trả về chỉ số giới hạn của chuỗi con đối xứng
func extendRange(s string, l, r int) (int, int) {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }
    return l + 1, r - 1
}

func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }

    start, end := 0, 0

    for i := 0; i < len(s); i++ {
        // Mở rộng từ vị trí với một ký tự ở giữa (đối xứng lẻ)
        l1, r1 := extendRange(s, i, i)
        // Mở rộng từ vị trí với hai ký tự ở giữa (đối xứng chẵn)
        l2, r2 := extendRange(s, i, i+1)

        // Cập nhật nếu tìm được chuỗi con đối xứng dài hơn
        if r1-l1 > end-start {
            start, end = l1, r1
        }
        if r2-l2 > end-start {
            start, end = l2, r2
        }
    }

    return s[start : end+1]
}
