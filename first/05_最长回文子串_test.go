package first

import (
	"fmt"
	"math"
	"testing"
)

func expandAroundCenter(s string, left int, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := int(math.Max(float64(len1), float64(len2)))
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return string(s[start : end+1])
}
func TestLongestPalindrome(t *testing.T) {
	s1 := "abacdfgdcaba"
	s2 := "babad"
	s3 := "cbbd"
	fmt.Println(longestPalindrome(s1))
	fmt.Println(longestPalindrome(s2))
	fmt.Println(longestPalindrome(s3))

}

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
// 示例 1：
// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。

// 示例 2：
// 输入: "cbbd"
// 输出: "bb"
