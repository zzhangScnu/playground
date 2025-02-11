package divide_conquer

import "math"

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
// 示例 1：
//
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
//
// 示例 2：
//
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
// 提示：
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	shorter, longer := nums1, nums2
	if len(nums1) > len(nums2) {
		shorter, longer = nums2, nums1
	}
	total := len(shorter) + len(longer)
	half := (total + 1) / 2
	low, high := 0, len(shorter)-1
	for {
		mid := low + (high-low)/2
		shorterIdx, longerIdx := mid, half-mid-2

		shorterLeft, shorterRight := math.MinInt, math.MaxInt
		if shorterIdx >= 0 {
			shorterLeft = shorter[shorterIdx]
		}
		if shorterIdx+1 < len(shorter) {
			shorterRight = shorter[shorterIdx+1]
		}

		longerLeft, longerRight := math.MinInt, math.MaxInt
		if longerIdx >= 0 {
			longerLeft = longer[longerIdx]
		}
		if longerIdx+1 < len(longer) {
			longerRight = longer[longerIdx+1]
		}

		if shorterLeft <= longerRight && longerLeft <= shorterRight {
			if total%2 == 1 {
				return float64(max(shorterLeft, longerLeft)) // 修正点2：返回左半部分的最大值
			} else {
				return float64(max(shorterLeft, longerLeft)+min(shorterRight, longerRight)) / 2.0
			}
		} else if shorterLeft > longerRight {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}
