package first

import (
	"fmt"
	"testing"
)

func expandAroundCenter(s string, left int, right int) string {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return s[L+1 : R] //求出后L已被减一，R已被加一，所以要L+1,右边不取，R不用减一
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	//start, end := 0, 0
	result := ""
	for i := 0; i < len(s); i++ {
		s1 := expandAroundCenter(s, i, i)
		s2 := expandAroundCenter(s, i, i+1)
		if len(s1) > len(result) {
			result = s1
		}
		if len(s2) > len(result) {
			result = s2
		}
	}
	return result
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
