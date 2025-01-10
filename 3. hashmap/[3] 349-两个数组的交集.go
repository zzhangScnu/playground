package hashmap

// 给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
//
// 示例 1：
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2]
//
// 示例 2：
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[9,4]
// 解释：[4,9] 也是可通过的
//
// 提示：
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
func intersection(nums1 []int, nums2 []int) []int {
	cnt := make(map[int]int)
	for _, num := range nums1 {
		cnt[num]++
	}
	resMap := make(map[int]bool)
	for _, num := range nums2 {
		if cnt[num] > 0 {
			resMap[num] = true
		}
	}
	var res []int
	for num, flag := range resMap {
		if flag {
			res = append(res, num)
		}
	}
	return res
}

/**
go没有set，只能用map来实现set。
*/
