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
	half := (total + 1) / 2
	l, r := 0, len(A)
	for l <= r {
		m := l + (r-l)>>1 // A数组贡献给整体中位数左边数组的个数
		i, j := m-1, half-m
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

// 这是由 中位数的定义 和 half = total / 2 的左半部分长度设计 共同决定的，核心是 “找到所有元素按序排列后的中间位置元素”，具体可通过「长度计算」和「直观例子」理解：
//1. 先明确核心：中位数的位置
//对于长度为 total 的有序数组，中位数的位置是固定的：
//若 total 为奇数（如 total=5）：中位数是 第 (total+1)/2 个元素（按 1 开始计数，即中间那个元素）。
//若 total 为偶数（如 total=4）：中位数是 第 total/2 和 total/2 +1 个元素的平均值。
//2. 再看 half = total / 2 对应的左右部分长度
//当 half = total / 2 时，左半部分的元素个数固定为 half，右半部分为 total - half：
//奇数情况（如 total=5）：
//half = 5/2 = 2（左半部分有 2 个元素）；
//右半部分 = 5 - 2 = 3（比左半多 1 个元素）。
//按中位数定义，中间元素是 第 3 个元素（1 开始计数）—— 而左半是前 2 个元素，右半是第 3~5 个元素，因此右半的 第一个元素（最小值） 就是第 3 个元素，即中位数。
//偶数情况（如 total=4）：
//half = 4/2 = 2（左半 2 个元素，右半 2 个元素）；
//中位数是第 2 和第 3 个元素的平均值 —— 左半的最后一个元素（最大值）是第 2 个，右半的第一个元素（最小值）是第 3 个，因此计算逻辑不变。
//3. 直观例子验证
//假设两个数组合并后为 [1,3,5,7,9]（total=5，奇数）：
//half = 5/2 = 2 → 左半是前 2 个元素 [1,3]，右半是后 3 个元素 [5,7,9]；
//右半的最小值是 5，恰好是整个数组的中位数（第 3 个元素）。
//若合并后为 [1,3,5,7]（total=4，偶数）：
//half=2 → 左半 [1,3]，右半 [5,7]；
//左半最大值 3 + 右半最小值 5 的平均值 4，就是中位数。
//总结
//当 half = total / 2 时：
//奇数长度下，右半部分比左半多 1 个元素，且右半的第一个元素（最小值）正好是 “中间位置的元素”，因此中位数是 min(Aright, Bright)；
//本质是通过调整 “左右部分的长度分配”，让中位数始终对应到 “右半的起点”（奇数）或 “左右的交界”（偶数），核心逻辑与 half=(total+1)/2 一致，只是参考点从 “左半终点” 变成了 “右半起点”。
