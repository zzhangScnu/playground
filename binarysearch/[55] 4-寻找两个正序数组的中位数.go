package binarysearch

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
	half := (total + 1) / 2 // 这里是做了 total + 1 的兼容处理
	l, r := 0, len(A)
	for l <= r { // 这里的结束条件是 l > r
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
				return (float64(max(Aleft, Bleft)) + float64(min(Aright, Bright))) / 2.0
			}
		} else if Aleft > Bright {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func findMedianSortedArrays0415(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	if len(A) > len(B) { // 以较小的数组作为二分搜索的锚点，避免数组越界
		B, A = A, B
	}
	total := len(A) + len(B)
	half := total / 2 // A + B 的长度向下取整。如果在 total 为奇数的情况下，half 会偏小
	low, high, mid := 0, len(A)-1, 0
	for { // 一定能找到结果，所以不需要显式指定结束条件
		sum := low + high
		mid = sum / 2
		// 处理负数除法，统一向下取整（默认向零取整）。避免 mid 取不到负数
		// 为什么会有负数？当 low 或 high 超出了有效范围，即 A / B 中有一个全都不会纳入结果集
		if sum < 0 && sum%2 != 0 {
			mid -= 1
		}
		Aleft := math.MinInt // 兼容该场景：如果 C' 全都取自 B'，即 A 上的索引会向左收缩直到小于0。此时 Aleft 为负无穷，一定小于 Bright
		if mid >= 0 {
			Aleft = A[mid]
		}
		Aright := math.MaxInt // 兼容该场景：如果 C' 全都取自 A'，即 B 上的索引会向右收缩直到大于数组末位元素索引。此时 Aright 为负无穷，一定小于 Bright
		if mid+1 < len(A) {
			Aright = A[mid+1]
		}
		Bleft := math.MinInt
		if half-mid-2 >= 0 {
			Bleft = B[half-mid-2]
		}
		Bright := math.MaxInt
		if half-mid-1 < len(B) {
			Bright = B[half-mid-1]
		}
		if Aleft <= Bright && Bleft <= Aright {
			if total%2 == 1 {
				return float64(min(Aright, Bright))
			}
			return float64(max(Aleft, Bleft)+min(Aright, Bright)) / 2.0
		}
		if Aleft > Bright {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}

/**
思路：
有序数组 A 和数组 B的中位数 == 排序(A + B)的中间元素 == (A 的前半段 A' + B 的前半段 B') 组成的 C' 的末尾元素
所以用二分搜索来锚定 A' 的大小，从而计算出 B' 的大小。
通过比较 A' 边界两边的元素与 B' 边界两边的元素，来判断 A' 和 B' 组成的【前半段】数组 C' 是否合法。
如果合法，则在边界处取中位数；
如果不合法，则通过二分搜索增加 / 减少 A' 的大小。
*/
