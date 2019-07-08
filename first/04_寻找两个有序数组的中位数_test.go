package first

import (
	"sort"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
	var result float64
	for _, v := range nums1 {
		nums = append(nums, v)
	}
	for _, v := range nums2 {
		nums = append(nums, v)
	}
	sort.Sort(sort.IntSlice(nums))
	length := len(nums)
	if length%2 == 0 {
		result = float64(nums[length/2]+nums[length/2-1]) / float64(2.0)
	} else {
		result = float64(nums[length/2])
	}
	return result
}
func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	t.Log(findMedianSortedArrays(nums1, nums2))

	nums3 := []int{1, 2}
	nums4 := []int{3, 4}
	t.Log(findMedianSortedArrays(nums3, nums4))
}

// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
// 你可以假设 nums1 和 nums2 不会同时为空。

// 示例 1:
// nums1 = [1, 3]
// nums2 = [2]
// 则中位数是 2.0

// 示例 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// 则中位数是 (2 + 3)/2 = 2.5
