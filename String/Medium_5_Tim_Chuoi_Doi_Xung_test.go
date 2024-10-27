package string

import (
    "testing"
)

func TestLongestPalindrome(t *testing.T) {
    testCases := []struct {
        input    string
        expected string
    }{
        {"babad", "bab"},
        {"cbbd", "bb"},
        {"a", "a"},
        {"ac", "a"},
        {"forgeeksskeegfor", "geeksskeeg"},
        {"", ""},
        {"abcba", "abcba"},
        {"abcd", "a"},
    }

    for _, tc := range testCases {
        result := longestPalindrome(tc.input)
        if result != tc.expected && len(result) != len(tc.expected) {
            t.Errorf("Test failed for input: %s. Expected: %s, Got: %s", tc.input, tc.expected, result)
        }
    }
}
