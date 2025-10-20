package prefixsum

//给你一个二进制数组 nums ，你需要从中删掉一个元素。
//
// 请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。
//
// 如果不存在这样的子数组，请返回 0 。
//
//
//
// 提示 1：
//
//
//输入：nums = [1,1,0,1]
//输出：3
//解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
//
// 示例 2：
//
//
//输入：nums = [0,1,1,1,0,1,1,0,1]
//输出：5
//解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。
//
// 示例 3：
//
//
//输入：nums = [1,1,1]
//输出：2
//解释：你必须要删除一个元素。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// nums[i] 要么是 0 要么是 1 。
func longestSubarray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	pre, suf := make([]int, n+1), make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if nums[i-1] == 1 {
			pre[i] = pre[i-1] + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 1 {
			suf[i] = suf[i+1] + 1
		}
	}
	var maxLen int
	for i := 1; i < n+1; i++ {
		maxLen = max(maxLen, pre[i-1]+suf[i])
	}
	return maxLen
}
