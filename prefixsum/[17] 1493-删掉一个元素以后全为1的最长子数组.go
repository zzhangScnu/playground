package prefixsum

// 给你一个二进制数组 nums ，你需要从中删掉一个元素。
//
// 请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。
//
// 如果不存在这样的子数组，请返回 0 。
//
// 提示 1：
//
// 输入：nums = [1,1,0,1]
// 输出：3
// 解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
//
// 示例 2：
//
// 输入：nums = [0,1,1,1,0,1,1,0,1]
// 输出：5
// 解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。
//
// 示例 3：
//
// 输入：nums = [1,1,1]
// 输出：2
// 解释：你必须要删除一个元素。
//
// 提示：
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

/**
思路：
遍历每个位置i，若删除本元素即第 i - 1 个元素，
此时最长连续 1 的长度 = 左侧连续 1 长度 + 右侧连续 1 长度
				   = pre[i-1] + suf[i]

pre 和 suf 的长度均为 n + 1，避免处理时做特殊边界兼容。
但是前缀和 / 后缀和的计算逻辑存在差异。
pre[i] = nums[0, i - 1]子数组中的元素 1 长度，pre 的 i -> nums 的 i - 1
suf[i] = nums[i, n - 1]子数组中的元素 1 长度，suf 的 i -> nums 的 i

最终执行删除遍历时，i 的范围为[1, n]，表示的是 nums[0...n - 1]，
所以对于 i，删除的是nums[i - 1]，
此时左边的连续最长 1 的长度 = nums[0, i - 2]子数组中的元素 1 长度 = pre[i - 1]
此时右边的连续最长 1 的长度 = nums[i, n - 1]子数组中的元素 1 长度 = suf[i]
*/
