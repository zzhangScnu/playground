package hashmap

// 给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
//
// 0 <= i, j, k, l < n
// nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
//
// 示例 1：
//
// 输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
// 输出：2
// 解释：
// 两个元组如下：
// 1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1)
// + 2 = 0
// 2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1)
// + 0 = 0
//
// 示例 2：
//
// 输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
// 输出：1
//
// 提示：
//
// n == nums1.length
// n == nums2.length
// n == nums3.length
// n == nums4.length
// 1 <= n <= 200
// -2²⁸ <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2²⁸
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	cnt12 := make(map[int]int)
	size := len(nums1)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cnt12[nums1[i]+nums2[j]]++
		}
	}
	var res int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if _, ok := cnt12[0-nums3[i]-nums4[j]]; ok {
				res += cnt12[0-nums3[i]-nums4[j]]
			}
		}
	}
	return res
}

/**
使用哈希，将O(n4)的时间复杂度降低到O(n2)。
*/
