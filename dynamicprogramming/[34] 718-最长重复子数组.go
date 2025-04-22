package dynamicprogramming

import "math"

// 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
//
// 示例 1：
//
// 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
// 输出：3
// 解释：长度最长的公共子数组是 [3,2,1] 。
//
// 示例 2：
//
// 输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
// 输出：5
//
// 提示：
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 100
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	res := math.MinInt
	for j := 0; j < n; j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
			if res < dp[0][j] {
				res = dp[0][j]
			}
		}
	}
	for i := 0; i < m; i++ {
		if nums2[0] == nums1[i] {
			dp[i][0] = 1
			if res < dp[i][0] {
				res = dp[i][0]
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

/**
DP数组及下标含义：
- i：nums1的下标；
- j：nums2的下标；
- dp[i][j]：【以nums1[i]结尾的子数组】和【以nums2[j]结尾的子数组】中的公共最长子数组长度。

递推公式：
if nums1[i] == nums2[j]
dp[i][j] = dp[i-1][j-1] + 1
用一个res动态记录最大值。
因为dp[m-1][n-1]表示以nums1[m-1]结尾的子数组和nums2[n-1]结尾的子数组中的公共最长子数组长度，不一定是最大值。

但res需要特殊处理，如(1, 2, 3, 2, 8)和(5, 6, 1, 4, 7)，答案在第一行和第一列初始化时就应得出。

初始化：
第一行和第一列。

遍历方向：从左到右
*/

/**
解法二：
不在初始化过程中特殊处理res，而是从i, j == 0, 0开始推导结果，同时需要注意i-1和j-1可能导致的数组越界问题。
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
		}
	}
	for i := 0; i < m; i++ {
		if nums2[0] == nums1[i] {
			dp[i][0] = 1
		}
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] == nums2[j] && i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

解法三：
将DP数组的含义变为：
【以nums1[i-1]结尾的子数组】和【以nums2[j-1]结尾的子数组】中的公共最长子数组长度。
则第一行和第一列是默认零值，没有实际语义，只是为了推导结果、规避i-1和j-1可能带来的数组越界问题。
*/
