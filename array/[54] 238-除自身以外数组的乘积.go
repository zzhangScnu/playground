package array

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
// 示例 1:
//
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
//
// 示例 2:
//
// 输入: nums = [-1,1,0,-3,3]
// 输出: [0,0,9,0,0]
//
// 提示：
//
// 2 <= nums.length <= 10⁵
// -30 <= nums[i] <= 30
// 输入 保证 数组 answer[i] 在 32 位 整数范围内
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
func productExceptSelf(nums []int) []int { // prefix[i] = nums[0] × nums[1] × ... × nums[i-1]（即nums[i]左侧所有元素的乘积，不包括nums[i]）
	n := len(nums)
	prefix, suffix := make([]int, n), make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = prefix[i] * suffix[i]
	}
	return res
}

// ⬆️⬇️两个方法的区别主要在于：前后缀数组定义的不同，及其带来的索引取数的不同。

func productExceptSelf3(nums []int) []int { // prefix[i] = nums[0] × nums[1] × ... × nums[i]（即nums[i]左侧所有元素的乘积，包括nums[i]）
	n := len(nums)
	pre, suf := make([]int, n+1), make([]int, n+1)
	pre[0], suf[n] = 1, 1
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i]
	}
	answer := make([]int, n)
	for i := 0; i < n; i++ {
		answer[i] = pre[i] * suf[i+1]
	}
	return answer
}

func productExceptSelfII(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}
	return res
}

/**
思路一：
O(n)时间复杂度 + O(n)空间复杂度

如数组a, b, c, d, e，乘积为a * b * c * d * e。
除自身以外的乘积 = 左侧所有元素的乘积 * 右侧所有元素的乘积，即：
除c以外的数的乘积为：(a * b) * (d * e)
套用前缀和数组的思路，可以转换为前缀积数组&后缀积数组：
- 前缀积prefix[i]：[0, i - 1]索引范围内的所有元素的乘积(不包含nums[i])；
- 后缀积suffix[i]：[i + 1, n)索引范围内的所有元素的乘积(不包含nums[i])。
除nums[i]以外的数的乘积为：prefix[i] * suffix[i]。

如定义，前缀积prefix[i]表示[0, i - 1]索引范围内的所有元素的乘积，故初始化prefix[0] = 1(方便后续相乘计算)，
在正向遍历nums时，令prefix[i] = prefix[i - 1] * nums[i - 1]。
后缀积同理，但需令suffix[n-1] = 1，且反向遍历nums，令prefix[i] = prefix[i + 1] * nums[i + 1]。

前缀积和后缀积数组计算完成后，既可再次遍历nums，推导结果数组。


思路二：
O(n)时间复杂度 + O(1)空间复杂度

整体思路不变，复用结果集数组res。
先在res上记录前缀积数组，再使用一个额外的变量记录后缀积，同时计算到res上。
*/
