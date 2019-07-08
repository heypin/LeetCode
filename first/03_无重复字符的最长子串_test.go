package first

import (
	"fmt"
	"math"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	n, ans := len(s), 0
	m := map[byte]int{}
	for i, j := 0, 0; j < n; j++ {
		if value, ok := m[s[j]]; ok {
			i = int(math.Max(float64(value), float64(i)))
		}
		ans = int(math.Max(float64(ans), float64(j-i+1)))
		m[s[j]] = j + 1
	}
	return ans
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s1, s2, s3 := "abcabcbb", "bbbbb", "pwwkew"
	out1 := lengthOfLongestSubstring(s1)
	out2 := lengthOfLongestSubstring(s2)
	out3 := lengthOfLongestSubstring(s3)
	fmt.Print(out1, " ", out2, " ", out3)
}

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 示例 1:
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

// 示例 2:
// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

// 示例 3:
// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
// 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
