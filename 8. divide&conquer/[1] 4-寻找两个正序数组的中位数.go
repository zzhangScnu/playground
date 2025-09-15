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
	A, B := nums1, nums2
	if len(A) > len(B) {
		A, B = nums2, nums1
	}
	total := len(nums1) + len(nums2)
	half := total / 2
	l, r := 0, len(A)
	for l <= r {
		m := l + (r-l)>>1
		i, j := m-1, half-m-1
		Aleft := math.MinInt
		if i >= 0 {
			Aleft = A[i]
		}
		Aright := math.MaxInt
		if i+1 < len(A) {
			Aright = A[i+1]
		}
		Bleft := math.MinInt
		if j >= 0 {
			Bleft = B[j]
		}
		Bright := math.MaxInt
		if j+1 < len(B) {
			Bright = B[j+1]
		}
		if Aleft <= Bright && Bleft <= Aright {
			if total%2 == 1 {
				return float64(max(Aleft, Bleft))
			} else {
				return float64((max(Aleft, Bleft) + min(Aright, Bright)) / 2)
			}
		} else if Aleft > Bright {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
